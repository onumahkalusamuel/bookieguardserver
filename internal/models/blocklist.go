package models

import (
	"bookieguardserver/config"
)

// blocklist (id, category_id, website)
type Blocklist struct {
	BaseModel
	CategoryID string `gorm:"not null;references:blocklist_categories(id)"`
	Website    string `gorm:"not null;unique"`
}

func (m *Blocklist) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *Blocklist) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// Delete function
func (m *Blocklist) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

func (m *Blocklist) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *Blocklist) ReadAll() (bool, []Blocklist) {
	var blocklists []Blocklist
	if result := config.DB.Find(&blocklists, &m); result.Error != nil {
		return false, blocklists
	}
	return true, blocklists
}

func (m *Blocklist) ReadAllFull() (bool, []Blocklist) {
	var blocklists []Blocklist
	if result := config.DB.
		Joins("LEFT JOIN blocklist_categories ON blocklist_categories.id=blocklists.category_id").
		Find(&blocklists, &m); result.Error != nil {
		return false, blocklists
	}
	return true, blocklists
}
