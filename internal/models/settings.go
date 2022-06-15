package models

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"gorm.io/gorm"
)

// Settings struct
type Settings struct {
	gorm.Model
	Setting string `gorm:"not null;unique_index"`
	Value   string `gorm:"not null"`
}

// Create creates a new setting
func (s *Settings) Create() error {
	return config.DB.Create(&s).Error
}

// Read reads a setting
func (s *Settings) Read() error {
	return config.DB.First(&s, &s).Error
}

// Update updates a setting
func (s *Settings) Update() error {
	return config.DB.Save(s).Error
}

// Delete deletes a setting
func (s *Settings) Delete() bool {
	if result := config.DB.First(&s, &s); result.Error != nil {
		return false
	}
	config.DB.Delete(&s)
	return true
}

// ReadAll reads all settings
func (s *Settings) ReadAll() (bool, []Settings) {
	var settings []Settings
	if result := config.DB.Find(&settings, &s); result.Error != nil {
		return false, settings
	}
	return true, settings
}
