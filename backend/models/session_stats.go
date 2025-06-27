package models

import (
	"time"
)

type SessionStat struct {
	ID              uint      `gorm:"primarykey;autoIncrement"`
	SessionID       string    `gorm:"column:session_id;index:idx_session_id;not null"`
	UserID          uint      `gorm:"column:user_id;index:idx_session_stats_user_id;not null"`
	TotalDistance   float64   `gorm:"column:total_distance"`
	MaxSpeed        float64   `gorm:"column:max_speed"`
	AvgSpeed        float64   `gorm:"column:avg_speed"`
	MaxRPM          int       `gorm:"column:max_rpm"`
	AvgRPM          int       `gorm:"column:avg_rpm"`
	FuelConsumed    float64   `gorm:"column:fuel_consumed"`
	AvgConsumption  float64   `gorm:"column:avg_consumption"`
	MaxTemperature  float64   `gorm:"column:max_temperature"`
	TripDuration    int       `gorm:"column:trip_duration"`
	DataPointsCount int       `gorm:"column:data_points_count"`
	CalculatedAt    time.Time `gorm:"column:calculated_at;default:CURRENT_TIMESTAMP"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID"`
	User    User    `gorm:"foreignKey:UserID;references:ID"`
}

func (SessionStat) TableName() string {
	return "session_stats"
}
