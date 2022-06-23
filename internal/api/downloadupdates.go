package api

import (
	"bookieguardserver/config"
	"bookieguardserver/internal/helpers"

	"github.com/gin-gonic/gin"
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
