package models

import (
	"time"

	"github.com/onumahkalusamuel/bookieguardserver/config"
	"gorm.io/gorm"
)

// computers (id, user_id, block_group_id, computer_name, hashed_id, unlock_code)
type Computer struct {
	gorm.Model
	UserID       uint   `gorm:"not null;references:users(id)"`
	BlockGroupID uint   `gorm:"not null;references:block_groups(id)"`
	ComputerName string `gorm:"not null"`
	HashedID     string `gorm:"not null;unique;index"`
	LastPing     string `gorm:"default:null;type:timestamp"`
}

// create Create function
func (m *Computer) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *Computer) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *Computer) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Computer) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Computer) ReadAll() (bool, []Computer) {
	var computers []Computer
	if result := config.DB.Find(&computers, &m); result.Error != nil {
		return false, computers
	}
	return true, computers
}

// UpdateLastPing with current timestamp
func (m *Computer) UpdateLastPing() error {
	return config.DB.Model(&m).Update("last_ping", time.Now()).Error
}
