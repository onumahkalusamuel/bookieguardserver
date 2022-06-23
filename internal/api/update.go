package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/helpers"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

// Request: hashedID, appVersion, blocklistHash, dashboardHash
// Response: success, message, appVersion, blocklist, blocklistHash, expired, expirationDate, dashboardFiles, dashboardHash

func Update(c *gin.Context) {

	r, _ := c.Get("requestBody")
	requestBody := r.(config.BodyStructure)

	responseBody := config.BodyStructure{}
	responseBody["success"] = "false"
	responseBody["message"] = "Invalid request."
	responseBody["appVersion"] = "1.0.0"
	responseBody["blocklist"] = ""
	responseBody["expired"] = "false"
	responseBody["dashboardFiles"] = ""

	if requestBody["hashedID"] == "" {
		c.Set("responseBody", responseBody)
		return
	}

	computer := models.Computer{}
	computer.HashedID = requestBody["hashedID"]
	computer.Read()

	if computer.ID == "" {
		responseBody["message"] = "Computer not found."
		c.Set("responseBody", responseBody)
		return
	}

	// fetch blockgroup and check expiration
	blockgroup := models.BlockGroup{}
	blockgroup.ID = computer.BlockGroupID
	blockgroup.Read()

	if blockgroup.ID == "" {
		responseBody["message"] = "Computer not found."
		c.Set("responseBody", responseBody)
		return
	}

	// set response as success
	responseBody["success"] = "true"

	// continue with the rest of the response
	if blockgroup.ExpirationDate < time.Now().Format("2006-01-02") {
		responseBody["expired"] = "true"
		responseBody["expirationDate"] = blockgroup.ExpirationDate
	}

	// check for appVersion update
	// s := models.Settings{}
	// s.Setting = "appVersion"
	// s.Read()

	// if s.Value != requestBody["appVersion"] {
	// 	responseBody["appVersion"] = s.Value
	// }

	// check for dashboardHash update
	// s.Setting = "dashboardHash"
	// s.Read()

	// if s.Value != requestBody["dashboardHash"] {
	// 	responseBody["dashboardHash"] = s.Value
	// 	// pull in the dashboard files
	// 	s.Setting = "dashboardFiles"
	// 	s.Read()
	// 	responseBody["dashboardFiles"] = s.Value
	// }

	// check for blocklist update
	blocklist := helpers.GetBlockList(blockgroup.ID)
	blocklistHash := helpers.GetHash(blocklist)

	if blocklistHash != requestBody["blocklistHash"] {
		responseBody["blocklist"] = blocklist
		responseBody["blocklistHash"] = blocklistHash
	}

	// add expirationDate
	responseBody["expirationDate"] = blockgroup.ExpirationDate

	// return response
	responseBody["message"] = "Update successful."
	c.Set("responseBody", responseBody)
}
