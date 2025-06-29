package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"math"
	"os"
	"strings"
	"time"
)

type Session struct {
	ID        uint      `gorm:"primarykey;autoIncrement"`
	SessionID string    `gorm:"column:session_id;uniqueIndex:idx_sessions_unique_session_id;not null"`
	DeviceID  string    `gorm:"column:device_id;index:idx_sessions_device_id;not null"`
	UserID    uint      `gorm:"column:user_id;index:idx_session_user_id;not null"`
	Version   int       `gorm:"column:version"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`

	StartTime       time.Time  `gorm:"column:start_time;index:idx_sessions_start_time;default:CURRENT_TIMESTAMP"`
	EndTime         *time.Time `gorm:"column:end_time"`
	UploadFrequency int        `gorm:"column:upload_frequency"`
	TotalRecords    int        `gorm:"column:total_records;default:0"`
	IsActive        bool       `gorm:"column:is_active;default:1"`

	Device Device `gorm:"foreignKey:DeviceID;references:DeviceID"`
	User   User   `gorm:"foreignKey:UserID;references:ID"`
}

func (*Session) TableName() string {
	return "sessions"
}

// SessionFindOrCreate finds an existing session by sessionID or creates a new one with the provided details.
// Returns the session and a boolean indicating whether a new record was created (true) or an existing one was found (false).
func SessionFindOrCreate(sessionID string, deviceID string, userID uint, version int, startTime time.Time) (Session, bool, error) {
	var session Session

	// Try to find existing session first
	err := DBSQLite.Where("session_id = ?", sessionID).First(&session).Error
	if err == nil {
		// Session already exists
		return session, false, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Unexpected error
		return session, false, err
	}

	// Create new session
	session = Session{
		SessionID: sessionID,
		DeviceID:  deviceID,
		UserID:    userID,
		Version:   version,
		StartTime: startTime,
		EndTime:   &startTime,
		IsActive:  true,
	}

	err = DBSQLite.Create(&session).Error
	if err != nil {
		// Handle unique constraint violation
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			// Another process created it in the meantime, try to fetch it
			if findErr := DBSQLite.Where("session_id = ?", sessionID).First(&session).Error; findErr == nil {
				return session, false, nil
			}
		}
		return session, false, err
	}

	return session, true, nil
}

// SessionGetByIdentifiers retrieves a session based on session ID, device ID, and user ID.
// Returns the session and an error if the session was not found.
func SessionGetByIdentifiers(sessionID string, deviceID string, userID uint) (Session, error) {
	var session Session
	err := DBSQLite.Where("session_id = ? AND device_id = ? AND user_id = ?", sessionID, deviceID, userID).First(&session).Error
	return session, err
}

// SessionListGetByUserAndDeviceID retrieves all sessions associated with a specific user ID and device ID,
// ordered by end_time in descending order.
func SessionListGetByUserAndDeviceID(userID uint, deviceID string) ([]Session, error) {
	var sessions []Session
	err := DBSQLite.Where("user_id = ? AND device_id = ?", userID, deviceID).Order("end_time DESC").Find(&sessions).Error
	return sessions, err
}

// SessionUpdateVehicleProfile updates a session's vehicle profile ID.
// It takes a session ID and a vehicle profile ID, and updates the session record.
// Returns an error if the update operation fails.
//func SessionUpdateVehicleProfile(sessionID string, vehicleProfileID uint) error {
//	result := DBSQLite.Model(&Session{}).
//		Where("session_id = ?", sessionID).
//		Update("vehicle_profile_id", vehicleProfileID)
//
//	return result.Error
//}

// SessionUpdateActivityAndRecords updates session fields such as end_time, is_active, and increments total_records by 1.
// It returns an error if the database operation fails.
func SessionUpdateActivityAndRecords(sessionID string, endTime time.Time) error {
	result := DBSQLite.Model(&Session{}).
		Where("session_id = ?", sessionID).
		Updates(map[string]interface{}{
			"end_time":      &endTime,
			"is_active":     true,
			"total_records": gorm.Expr("total_records + ?", 1),
		})

	return result.Error
}

// GetSessionData retrieves time-series data for this session from InfluxDB.
// It includes data from 10 minutes before the session start time to 10 minutes after the session end time.
// Returns data, GPS coordinates, and the center point of the captured coordinates.
func (session *Session) GetSessionData() (map[string]map[string]interface{}, [][]float64, []float64, error) {
	sessionStartTimeStr := session.StartTime.Add(-10 * time.Minute).Format(time.RFC3339)
	sessionEndTimeStr := session.EndTime.Add(10 * time.Minute).Format(time.RFC3339)

	query := `from(bucket:"` + os.Getenv("INFLUX_BUCKET") + `")
    |> range(start: ` + sessionStartTimeStr + `, stop: ` + sessionEndTimeStr + `)
    |> filter(fn: (r) => r._measurement == "gorque_data")
    |> filter(fn: (r) => r.id == "` + session.DeviceID + `")
    |> filter(fn: (r) => r.session == "` + session.SessionID + `")
    |> sort(columns: ["_time"], desc: false)`

	queryAPI := DBInflux.QueryAPI(os.Getenv("INFLUX_ORG"))

	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		return nil, nil, nil, err
	}

	var orderedTimes []time.Time
	timeBasedData := make(map[string]map[string]interface{})
	coords := make([][]float64, 0)

	for result.Next() {
		record := result.Record()
		recordTime := record.Time()
		timeStr := record.Time().Format(time.RFC3339)
		fieldName := record.Field()
		fieldValue := record.Value()

		if _, exists := timeBasedData[timeStr]; !exists {
			timeBasedData[timeStr] = make(map[string]interface{})
			orderedTimes = append(orderedTimes, recordTime)
		}

		timeBasedData[timeStr][fieldName] = fieldValue

		for key, value := range record.Values() {
			if key != "_time" && key != "_value" && key != "_field" && key != "_measurement" {
				timeBasedData[timeStr][key] = value
			}
		}
	}

	if result.Err() != nil {
		return nil, nil, nil, result.Err()
	}

	type TimeEntry struct {
		Time time.Time
		Data map[string]interface{}
	}

	var orderedData []TimeEntry
	for _, t := range orderedTimes {
		timeStr := t.Format(time.RFC3339)
		orderedData = append(orderedData, TimeEntry{
			Time: t,
			Data: timeBasedData[timeStr],
		})
	}

	// Convert timeBasedData coordinates to coords in chronological order
	minLat := math.MaxFloat64
	maxLat := -math.MaxFloat64
	minLon := math.MaxFloat64
	maxLon := -math.MaxFloat64
	for _, entry := range orderedData {
		data := entry.Data
		if lat, ok := data["kff1006"].(float64); ok {
			if lon, ok := data["kff1005"].(float64); ok {
				coords = append(coords, []float64{lat, lon})
				if lat < minLat {
					minLat = lat
				}
				if lat > maxLat {
					maxLat = lat
				}
				if lon < minLon {
					minLon = lon
				}
				if lon > maxLon {
					maxLon = lon
				}
			}
		}
	}

	centerLat := (minLat + maxLat) / 2
	centerLon := (minLon + maxLon) / 2
	centerPoint := []float64{centerLat, centerLon}

	return timeBasedData, coords, centerPoint, nil
}
