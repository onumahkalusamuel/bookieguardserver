package api

import (
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/helpers"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

// Request: email, activationCode, hashedID, computerName
// Response: email, hashedID, success, activated, blocklist, expirationDate, unlockCode

func Activation(c *gin.Context) {

	var availableBlockgroup models.BlockGroup

	r, _ := c.Get("requestBody")

	requestBody := r.(config.BodyStructure)

	valid, message := validateActivation(requestBody)

	if !valid {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": message,
		})
		return
	}

	// check if email belongs to a user
	user := models.User{}
	user.Email = requestBody["email"]
	user.Read()

	if user.ID == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Account not found. Please check the details and try again.",
		})
		return
	}

	// check if user has blockgroup
	availableBlockgroup = models.BlockGroup{}
	availableBlockgroup.UserID = user.ID
	availableBlockgroup.ActivationCode = requestBody["activationCode"]
	availableBlockgroup.Read()

	if availableBlockgroup.ID == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Invalid activation details provided.",
		})
		return
	}

	// check the valid activation date
	if availableBlockgroup.ExpirationDate < time.Now().Format("2006-01-02") {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Blockgroup has expired. Please upgrade your account to continue using the service.",
		})
		return
	}

	// check if computer has been activated
	computer := models.Computer{}
	computer.HashedID = requestBody["hashedID"]
	computer.Read()

	// start activation for new installation
	if computer.ID == "" {

		if availableBlockgroup.ActivatedComputers >= availableBlockgroup.TotalComputers {
			c.Set("responseBody", config.BodyStructure{
				"success": "false",
				"message": "Blockgroup is full. Please try another blockgroup.",
			})
			return
		}

		// create the computer record
		computer.UserID = user.ID
		computer.ComputerName = requestBody["computerName"]
		computer.BlockGroupID = availableBlockgroup.ID

		if computer.Create() != nil {
			c.Set("responseBody", config.BodyStructure{
				"success": "false",
				"message": "An error occured. Please try again later.",
			})
			return
		}

		// increment blockgroup.ActivatedComputers and save back
		newvalue := availableBlockgroup.ActivatedComputers + 1
		err := availableBlockgroup.UpdateSingle("activated_computers", newvalue)

		if err != nil {
			computer.Delete()
			c.Set("responseBody", config.BodyStructure{
				"success": "false",
				"message": "An error occured. Please try again later.",
			})
			return
		}

	}

	computer.UpdateLastPing()

	// activation successful. continue with the rest of the process
	// fetch all blocklists

	blocklists := helpers.GetBlockList(availableBlockgroup.ID)

	var responseBody = config.BodyStructure{
		"email":          requestBody["email"],
		"hashedID":       requestBody["hashedID"],
		"success":        "true",
		"activated":      "true",
		"blocklist":      blocklists,
		"expirationDate": availableBlockgroup.ExpirationDate,
		"unlockCode":     availableBlockgroup.UnlockCode,
	}
	// build response

	c.Set("responseBody", responseBody)

}

func validateActivation(requestBody config.BodyStructure) (bool, string) {

	if requestBody == nil {
		return false, "Invalid request."
	}

	if _, err := mail.ParseAddress(requestBody["email"]); err != nil {
		return false, "Invalid email provided."
	}

	if len(requestBody["activationCode"]) != 14 {
		return false, "Invalid activation code provided."
	}

	if requestBody["computerName"] == "" {
		return false, "Invalid request."
	}

	return true, ""
}
