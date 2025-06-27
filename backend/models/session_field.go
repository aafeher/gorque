package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type SessionField struct {
	ID          uint      `gorm:"primarykey;autoIncrement"`
	SessionID   string    `gorm:"column:session_id;uniqueIndex:idx_session_field;not null"`
	UserID      uint      `gorm:"column:user_id;index:idx_session_field_user_id;not null"`
	FieldKey    string    `gorm:"column:field_key;uniqueIndex:idx_session_field;not null"`
	ShortName   string    `gorm:"column:short_name"`
	FullName    string    `gorm:"column:full_name"`
	Unit        string    `gorm:"column:unit"`
	DefaultUnit string    `gorm:"column:default_unit"`
	MinValue    *float64  `gorm:"column:min_value"`
	MaxValue    *float64  `gorm:"column:max_value"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`

	Session Session `gorm:"foreignKey:SessionID;references:SessionID"`
	User    User    `gorm:"foreignKey:UserID;references:ID"`
}

func (SessionField) TableName() string {
	return "session_fields"
}

// SessionFieldCreate creates a new session field record in the database.
// Returns any error that occurred during the creation.
func SessionFieldCreate(sessionField *SessionField) error {
	result := DBSQLite.Create(sessionField)
	return result.Error
}

// SessionFieldFindOrCreate finds a session field by session ID and field key or creates a new one.
// It takes session ID, user ID, field key, and default unit value.
// Returns the session field, a boolean indicating whether a new record was created (true) or an existing one was found (false),
// and any error that occurred.
func SessionFieldFindOrCreate(sessionID string, userID uint, fieldKey string, defaultUnit string) (SessionField, bool, error) {
	sessionField := SessionField{
		SessionID:   sessionID,
		UserID:      userID,
		FieldKey:    fieldKey,
		DefaultUnit: defaultUnit,
	}

	result := DBSQLite.Where("session_id = ? AND field_key = ?",
		sessionField.SessionID, sessionField.FieldKey).
		FirstOrCreate(&sessionField)

	if result.Error != nil {
		return sessionField, false, result.Error
	}

	// Return whether this was a create operation (rows affected > 0)
	return sessionField, result.RowsAffected > 0, nil
}

// SessionFieldFindBySessionAndKey finds a session field by session ID and field key.
// Returns the session field and an error if not found.
func SessionFieldFindBySessionAndKey(sessionID string, fieldKey string) (SessionField, error) {
	var sessionField SessionField
	err := DBSQLite.Where("session_id = ? AND field_key = ?", sessionID, fieldKey).First(&sessionField).Error
	return sessionField, err
}

// SessionFieldSaveOrUpdate either creates a new session field or updates an existing one.
// It checks if a field with the given session ID and field key exists, and creates or updates accordingly.
// Returns any error that occurred during the operation.
func SessionFieldSaveOrUpdate(field SessionField) error {
	var existingField SessionField
	result := DBSQLite.Where("session_id = ? AND field_key = ?", field.SessionID, field.FieldKey).First(&existingField)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return SessionFieldCreate(&field)
		}
		return result.Error
	}

	return SessionFieldUpdate(&existingField, field)
}

// SessionFieldUpdate updates an existing session field record.
// Takes an existing session field and the updated data to apply.
// Returns any error that occurred during the update.
func SessionFieldUpdate(existingField *SessionField, updates SessionField) error {
	result := DBSQLite.Model(existingField).Updates(updates)
	return result.Error
}

// SessionFieldUpdateDefaultUnit updates the default unit value for a session field.
// It takes session ID, field key, and the new default unit value.
// Returns an error if the update operation fails.
func SessionFieldUpdateDefaultUnit(sessionID string, fieldKey string, defaultUnit string) error {
	result := DBSQLite.Model(&SessionField{}).
		Where("session_id = ? AND field_key = ?", sessionID, fieldKey).
		Update("default_unit", defaultUnit)

	return result.Error
}
