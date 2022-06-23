package models

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"gorm.io/gorm"
)

// Host struct
type Host struct {
	BaseModel
	Website   string         `gorm:"not null;uniqueIndex"`
	HashedID  string         `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // added to enable soft delete
}

// CreateHost creates a new host
func (h *Host) Create() error {
	return config.DB.Create(&h).Error
}

// ReadHost reads a host
func (h *Host) Read() error {
	return config.DB.First(&h, &h).Error
}

// UpdateHost updates a host
func (h *Host) Update() error {
	return config.DB.Save(h).Error
}

// DeleteHost deletes a host
func (h *Host) Delete() bool {
	if result := config.DB.First(&h, &h); result.Error != nil {
		return false
	}
	config.DB.Delete(&h)
	return true
}

// ReadAllHosts reads all hosts
func (h *Host) ReadAll() (bool, []Host) {
	var hosts []Host
	if result := config.DB.Find(&hosts, &h); result.Error != nil {
		return false, hosts
	}
	return true, hosts
}
