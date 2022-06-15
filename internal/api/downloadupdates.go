package api

import (
	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/internal/helpers"
)

// Request: hashedID, fileName
// Response: success, message, content

func DownloadUpdates(c *gin.Context) {

	// check if computer is activated
	r, _ := c.Get("requestBody")
	requestBody := r.(config.BodyStructure)

	// check if file with fileName exists on the server
	fileName := requestBody["fileName"]
	if fileName == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "Invalid request.",
		})
		return
	}

	fileContent := helpers.GetFileContent(config.UpdatePath + fileName)

	if fileContent == "" {
		c.Set("responseBody", config.BodyStructure{
			"success": "false",
			"message": "File not found.",
		})
		return
	}

	c.Set("responseBody", config.BodyStructure{
		"success": "true",
		"message": "File found.",
		"content": fileContent,
	})
}
