package models

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"gorm.io/gorm"
)

// users (id, email, password, phone, address)
type User struct {
	gorm.Model
	Name     string `gorm:"default:null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"default:null"`
	Phone    string `gorm:"default:null"`
	Address  string `gorm:"default:null"`
}

// create Create function
func (m *User) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *User) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *User) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *User) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *User) ReadAll() (bool, []User) {
	var users []User
	if result := config.DB.Find(&users, &m); result.Error != nil {
		return false, users
	}
	return true, users
}
