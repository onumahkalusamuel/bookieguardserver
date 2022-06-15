package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

// Request: hashedID, hosts
// Response: success, message

func UploadHosts(c *gin.Context) {
	r, _ := c.Get("requestBody")
	requestBody := r.(config.BodyStructure)

	if requestBody["hashedID"] == "" || requestBody["hosts"] == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Invalid request.",
		})
		return
	}

	hostsExploded := strings.Split(requestBody["hosts"], ",")

	for _, website := range hostsExploded {
		website = strings.TrimSpace(website)
		if website != "" {
			h := models.Host{}
			h.HashedID = requestBody["hashedID"]
			h.Website = website
			h.Create()
		}
	}

	c.Set("responseBody", config.BodyStructure{
		"success": "true",
		"message": "Hosts uploaded successfully.",
	})

}
