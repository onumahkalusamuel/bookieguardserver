package config

import (
	"net"

	"gorm.io/gorm"
)

var Key = "ab9312a52781f4b7c7edf4341ef940daff94c567ffa503c3db8125fec68c4225"

var SERVER_PROTOCOL = "http"
var SERVER_HOST = "localhost"
var SERVER_PORT = "8889"

const UpdatePath = "./updates/"

var DB *gorm.DB

var PaystackSecretKey = "sk_test_b9eb40c855f809c5f5a0633e99cf73198497fb7b"
var PaystackCallBackURL = SERVER_PROTOCOL + "://" + net.JoinHostPort(SERVER_HOST, SERVER_PORT) + "/account/paystack-callback"
var PaystackCurrency = "NGN"
var PaystackChannels = []string{"card", "bank", "ussd", "qr", "bank_transfer"}

type PaystackMetaData struct {
	BlockGroupID     string
	UserID           string
	PaymentReference string
}

type BodyStructure map[string]string
