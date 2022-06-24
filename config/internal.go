package config

import (
	"net"
	"os"

	"gorm.io/gorm"
)

var Key = ""

const UpdatePath = "./updates/"

var DB *gorm.DB

var PaystackSecretKey = ""
var PaystackCallBackURL = "https://"
var PaystackCurrency = "NGN"
var PaystackChannels = []string{"card", "bank", "ussd", "qr", "bank_transfer"}

type PaystackMetaData struct {
	BlockGroupID     string
	UserID           string
	PaymentReference string
}

type BodyStructure map[string]string

func SetUpEnv() {

	if os.Getenv("ENV") == "dev" {
		os.Setenv("PORT", "8889")
		os.Setenv("HOST", "localhost")

		PaystackCallBackURL = "http://"
	}

	Key = os.Getenv("APP_DECRYPT_KEY")
	PaystackSecretKey = os.Getenv("PAYSTACK_SECRET_KEY")
	PaystackCallBackURL = PaystackCallBackURL + net.JoinHostPort(os.Getenv("HOST"), os.Getenv("PORT")) + "/account/paystack-callback"
}
