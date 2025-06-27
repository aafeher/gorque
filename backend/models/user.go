package models

import "time"

// User represents a user entity with attributes such as ID, email, hashed password, name, and creation timestamp.
type User struct {
	ID        uint   `gorm:"primarykey"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
}

// TableName specifies the custom table name for the User struct when used with an ORM.
func (*User) TableName() string {
	return "users"
}

// UserCreate inserts a new User record into the database and returns an error if the operation fails.
func UserCreate(user *User) error {
	return DBSQLite.Create(user).Error
}

// UserGetByID retrieves a User record from the database by the given user ID. Returns the User or an error if not found.
func UserGetByID(id uint) (*User, error) {
	var user User
	if err := DBSQLite.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UserGetByEmail retrieves a User record from the database using the provided email address.
// It returns the User and an error, if any occurred during the query.
func UserGetByEmail(email string) (*User, error) {
	var user User
	if err := DBSQLite.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetDevices retrieves all devices associated with the user.
func (user *User) GetDevices() ([]Device, error) {
	return DeviceListGetByUserID(user.ID)
}

// UpdateName updates the name of a user in the database and returns an error if the operation fails.
func (user *User) UpdateName(name string) error {
	return DBSQLite.Model(user).Update("name", name).Error
}

// UpdatePassword updates the password of a user in the database and returns an error if the operation fails.
func (user *User) UpdatePassword(hashedPassword string) error {
	return DBSQLite.Model(user).Update("password", hashedPassword).Error
}
