package models

import (
	"bookieguardserver/config"
)

// contacts (id, name, email, subject, message, read_status)
type Contact struct {
	BaseModel
	Name       string `gorm:"default:null"`
	Email      string `gorm:"default:null"`
	Subject    string `gorm:"default:null"`
	Message    string `gorm:"default:null"`
	ReadStatus uint   `gorm:"default:0"`
}

// create Create function
func (m *Contact) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *Contact) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// create Update function
func (m *Contact) UpdateSingle(key string, value any) error {
	return config.DB.First(&m).Update(key, value).Error
}

// Delete function
func (m *Contact) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Contact) Read() error {

	result := config.DB.First(&m, &m)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ReadAll function
func (m *Contact) ReadAll() (bool, []Contact) {
	var contacts []Contact
	if result := config.DB.Order("read_status asc").Find(&contacts, &m); result.Error != nil {
		return false, contacts
	}
	return true, contacts
}
