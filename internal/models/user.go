package models

import (
	"bookieguardserver/config"
)

// users (id, email, password, phone, address)
type User struct {
	BaseModel
	Name     string `gorm:"default:null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"default:null"`
	Phone    string `gorm:"default:null"`
	UserType string `gorm:"default:'account'"`
	Token    string `gorm:"default:null"`
	Address  string `gorm:"default:null"`
}

// create Create function
func (m *User) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *User) Update() error {
	finder := User{}
	finder.ID = m.ID
	return config.DB.First(&m, &finder).Save(&m).Error
}

// create Update function
func (m *User) UpdateSingle(key string, value string) error {
	return config.DB.First(&m).Update(key, value).Error
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
