package config

import (
	"gorm.io/gorm"
)

const UpdatePath = "./updates/"

var DB *gorm.DB

var PaystackCurrency = "NGN"
var PaystackChannels = []string{"card", "bank", "ussd", "qr", "bank_transfer"}

type PaystackMetaData struct {
	BlockGroupID     string
	UserID           string
	PaymentReference string
}

type BodyStructure map[string]string
