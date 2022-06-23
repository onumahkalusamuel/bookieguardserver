package models

import (
	"bookieguardserver/config"
)

// payments (id, user_id, block_group_id, amount, gateway, date, details, status)
type Payment struct {
	BaseModel
	UserID           string `gorm:"not null;references:users(id)"`        // user_id
	BlockGroupID     string `gorm:"not null;references:block_groups(id)"` // block_group_id
	PaymentReference string `gorm:"default:null"`                         // the generated payment reference
	Amount           uint   `gorm:"default:null"`                         // amount to be paid
	Currency         string `gorm:"default:'NGN'"`                        // USD, GHC, NGN, etc
	PlanID           string `gorm:"default:null"`                         // in months
	Gateway          string `gorm:"default:null"`                         // paypal, stripe, flutterwave, paystack, etc
	Details          string `gorm:"default:null"`                         // details about the payment
	Status           string `gorm:"not null;default:'pending'"`           // pending, success, failed
}

// create Create function
func (m *Payment) Create() error {
	return config.DB.Create(&m).Error
}

// create Update function
func (m *Payment) Update() error {
	return config.DB.First(&m, &m).Save(&m).Error
}

// create Update function
func (m *Payment) UpdateSingle(key string, value string) error {
	return config.DB.First(&m).Update(key, value).Error
}

// Delete function
func (m *Payment) Delete() bool {
	if result := config.DB.First(&m, &m); result.Error != nil {
		return false
	}
	config.DB.Delete(&m)
	return true
}

// Read function
func (m *Payment) Read() error {
	return config.DB.First(&m, &m).Error
}

// ReadAll function
func (m *Payment) ReadAll() (bool, []Payment) {
	var payments []Payment
	if result := config.DB.Find(&payments, &m); result.Error != nil {
		return false, payments
	}
	return true, payments
}
