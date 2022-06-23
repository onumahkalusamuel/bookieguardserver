package models

import (
	"bookieguardserver/config"
)

// blocklist_categories (id, title)
type BlocklistCategory struct {
	BaseModel
	Title        string `gorm:"not null;unique"`
	DisplayTitle string `gorm:"default:null;unique"`
}

// create Create function
func (m *BlocklistCategory) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *BlocklistCategory) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *BlocklistCategory) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *BlocklistCategory) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *BlocklistCategory) ReadAll() (bool, []BlocklistCategory) {
	var blocklistCategories []BlocklistCategory
	if result := config.DB.Find(&blocklistCategories, &m); result.Error != nil {
		return false, blocklistCategories
	}
	return true, blocklistCategories
}
