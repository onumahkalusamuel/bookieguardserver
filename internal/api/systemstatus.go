package api

import (
	"bookieguardserver/config"
	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

// Request: hashedID
// Response: success, message
func SystemStatus(c *gin.Context) {

	r, _ := c.Get("requestBody")
	requestBody := r.(config.BodyStructure)

	if requestBody["hashedID"] == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Invalid request.",
		})
		return
	}

	// check if computer exists and update LastPing
	computer := models.Computer{}
	computer.HashedID = requestBody["hashedID"]
	if err := computer.UpdateLastPing(); err != nil {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Computer not found.",
		})
		return
	}

	c.Set("responseBody", config.BodyStructure{
		"success": "true",
		"message": "System status retrieved successfully.",
	})

}
