package models

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
)

// allowlist (id, block_group_id, website)
type Allowlist struct {
	BaseModel
	BlockGroupID string `gorm:"not null;references:block_groups(id)"`
	Website      string `gorm:"not null;"`
}

// create Create function
func (m *Allowlist) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *Allowlist) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *Allowlist) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Allowlist) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Allowlist) ReadAll() (bool, []Allowlist) {
	var allowlists []Allowlist
	if result := config.DB.Find(&allowlists, &m); result.Error != nil {
		return false, allowlists
	}
	return true, allowlists
}
