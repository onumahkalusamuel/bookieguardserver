package api

import (
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

// Request: email, activationCode, hashedID, computerName
// Response: email, hashedID, success, activated, blocklist, expirationDate, unlockCode

func Activation(c *gin.Context) {

	blocklistHolder := []string{}
	allowlistHolder := []string{}
	var availableBlockgroup models.BlockGroup

	r, _ := c.Get("requestBody")

	requestBody := r.(config.BodyStructure)

	valid, message := validateActivation(requestBody)

	if !valid {
		fmt.Println(message)
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

	if user.ID == 0 {
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

	if availableBlockgroup.ID == 0 {
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
	if computer.ID == 0 {

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
		availableBlockgroup.ActivatedComputers++
		if availableBlockgroup.Update() != nil {
			computer.Delete()
			c.Set("responseBody", config.BodyStructure{
				"success": "false",
				"message": "An error occured. Please try again later.",
			})
			return
		}

	}

	// activation successful. continue with the rest of the process
	// fetch all blocklists
	b := models.Blocklist{}
	_, blocklists := b.ReadAll()

	// fetch all allowlists
	a := models.Allowlist{
		BlockGroupID: availableBlockgroup.ID,
	}
	_, allowlists := a.ReadAll()

	// put allowlist.website into a string concatenation
	for _, allowlist := range allowlists {
		allowlistHolder = append(allowlistHolder, allowlist.Website)
	}
	allowed := strings.Join(allowlistHolder, ",")

	// loop through blocklists
	for _, blocklist := range blocklists {
		// check if website exists in allowlist and skip if it does
		if strings.Contains(allowed, blocklist.Website) {
			continue
		}

		// add blocklist.Website to blocklistHolder
		blocklistHolder = append(blocklistHolder, blocklist.Website)
	}

	var responseBody config.BodyStructure
	// build response
	responseBody["email"] = requestBody["email"]
	responseBody["hashedID"] = requestBody["hashedID"]
	responseBody["success"] = "true"
	responseBody["activated"] = "true"
	responseBody["blocklist"] = strings.Join(blocklistHolder, ",")
	responseBody["expirationDate"] = availableBlockgroup.ExpirationDate
	responseBody["unlockCode"] = availableBlockgroup.UnlockCode

	c.Set("responseBody", responseBody)

}

func validateActivation(requestBody config.BodyStructure) (bool, string) {

	if requestBody == nil {
		return false, "Invalid request."
	}

	if _, err := mail.ParseAddress(requestBody["email"]); err != nil {
		return false, "Invalid email provided."
	}

	if len(requestBody["activationCode"]) != 15 {
		return false, "Invalid activation code provided."
	}

	if requestBody["computerName"] == "" {
		return false, "Invalid request."
	}

	return true, ""
}
