package helpers

import "os"

func GetPasswordResetLink(serverHost string, token string) string {
	var protocol = "https://"
	if os.Getenv("ENV") == "dev" {
		protocol = "http://"
	}
	var passwordResetLink = protocol + serverHost + "/account/password-reset/" + token
	return passwordResetLink
}
