package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Device struct {
	ID        uint      `gorm:"primarykey;autoIncrement"`
	DeviceID  string    `gorm:"column:device_id;uniqueIndex:idx_device_unique_device_id;not null"`
	UserID    uint      `gorm:"column:user_id;index:idx_device_user_id;not null"`
	Version   int       `gorm:"column:v"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`

	ProfileBoostAdjust  float64 `gorm:"column:profileBoostAdjust"`
	ProfileDisplacement float64 `gorm:"column:profileDisplacement"`
	ProfileDragCoeff    float64 `gorm:"column:profileDragCoeff"`
	ProfileFuelCost     float64 `gorm:"column:profileFuelCost"`
	ProfileFuelType     int64   `gorm:"column:profileFuelType"`
	ProfileMPGAdjust    float64 `gorm:"column:profileMPGAdjust"`
	ProfileName         string  `gorm:"column:profileName"`
	ProfileOBDAdjust    float64 `gorm:"column:profileOBDAdjust"`
	ProfileOdometer     int64   `gorm:"column:profileOdometer"`
	ProfileTankCapacity float64 `gorm:"column:profileTankCapacity"`
	ProfileTankUsed     float64 `gorm:"column:profileTankUsed"`
	ProfileVe           float64 `gorm:"column:profileVe"`
	ProfileVehicleType  int64   `gorm:"column:profileVehicleType"`

	LastSeen time.Time `gorm:"column:last_seen"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (*Device) TableName() string {
	return "devices"
}

// DeviceCreate creates a new device record in the database.
// It takes a device object and inserts it into the database.
// Returns any error that occurred during the creation.
func DeviceCreate(device *Device) error {
	result := DBSQLite.Create(device)
	return result.Error
}

// DeviceCreateOrUpdate creates a new device record or updates an existing one with profile data.
// It takes a device object and either inserts it or updates the existing record.
// Returns any error that occurred during the operation.
func DeviceCreateOrUpdate(device *Device) error {
	var existingDevice Device

	// Try to find existing device
	result := DBSQLite.Where("device_id = ?", device.DeviceID).First(&existingDevice)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Device doesn't exist, create new one
			return DBSQLite.Create(device).Error
		}
		// Other error occurred
		return result.Error
	}

	// Device exists, update profile fields
	updates := map[string]interface{}{
		"profileBoostAdjust":  device.ProfileBoostAdjust,
		"profileDisplacement": device.ProfileDisplacement,
		"profileDragCoeff":    device.ProfileDragCoeff,
		"profileFuelCost":     device.ProfileFuelCost,
		"profileFuelType":     device.ProfileFuelType,
		"profileMPGAdjust":    device.ProfileMPGAdjust,
		"profileName":         device.ProfileName,
		"profileOBDAdjust":    device.ProfileOBDAdjust,
		"profileOdometer":     device.ProfileOdometer,
		"profileTankCapacity": device.ProfileTankCapacity,
		"profileTankUsed":     device.ProfileTankUsed,
		"profileVe":           device.ProfileVe,
		"profileVehicleType":  device.ProfileVehicleType,
		"v":                   device.Version,
	}

	return DBSQLite.Model(&existingDevice).Updates(updates).Error
}

// DeviceFindOrCreate finds an existing device by deviceID or creates a new one with the provided details.
// Returns the device, a boolean indicating whether a new record was created (true) or an existing one was found (false),
// and any error that occurred.
func DeviceFindOrCreate(deviceID string, userID uint) (Device, bool, error) {
	var device Device

	// Set basic device properties
	device = Device{
		DeviceID: deviceID,
		UserID:   userID,
		// LastSeen field is updated by a trigger after inserting sessions
	}

	// Try to find existing device or create a new one
	result := DBSQLite.Where("device_id = ?", deviceID).FirstOrCreate(&device)

	if result.Error != nil {
		return device, false, result.Error
	}

	// Return whether this was a create operation (rows affected > 0)
	return device, result.RowsAffected > 0, nil
}

// DeviceGetByDeviceID retrieves a device by its device ID and user ID.
func DeviceGetByDeviceID(userID uint, deviceID string) (Device, error) {
	var device Device
	err := DBSQLite.Where("user_id = ? AND device_id = ?", userID, deviceID).First(&device).Error
	return device, err
}

// DeviceListGetByUserID retrieves all devices associated with a specific user ID.
func DeviceListGetByUserID(userID uint) ([]Device, error) {
	var devices []Device
	err := DBSQLite.Where("user_id = ?", userID).Find(&devices).Error
	return devices, err
}

// GetSessions retrieves all sessions associated with this device and its user,
// ordered by end_time in descending order.
func (device *Device) GetSessions() ([]Session, error) {
	return SessionListGetByUserAndDeviceID(device.UserID, device.DeviceID)
}
