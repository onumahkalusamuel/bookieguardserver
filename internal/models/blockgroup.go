package models

import (
	"github.com/onumahkalusamuel/bookieguardserver/config"
)

// block_groups (id, user_id, title)
type BlockGroup struct {
	BaseModel
	UserID             string `gorm:"not null"`
	Title              string `gorm:"not null"`
	TotalComputers     uint   `gorm:"default:0"`
	ActivatedComputers uint   `gorm:"default:0"`
	ExpirationDate     string `gorm:"default:null;type:date"`
	UnlockCode         string `gorm:"not null;unique"`
	ActivationCode     string `gorm:"not null;unique"`
	CurrentPlan        string `gorm:"default:null;"`
}

// create Create function
func (m *BlockGroup) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *BlockGroup) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// create Update function
func (m *BlockGroup) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

// Delete function
func (m *BlockGroup) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *BlockGroup) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *BlockGroup) ReadAll() (bool, []BlockGroup) {
	var blockGroups []BlockGroup
	if result := config.DB.Find(&blockGroups, &m); result.Error != nil {
		return false, blockGroups
	}
	return true, blockGroups
}
