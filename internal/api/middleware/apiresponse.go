package middleware

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/pkg"
)

func ApiResponse() gin.HandlerFunc {

	return func(c *gin.Context) {

		var (
			failed   bool
			bytes    []byte
			err      error
			sendback config.BodyStructure
		)

		responseBody, exists := c.Get("responseBody")

		if !exists || responseBody == nil {
			sendback = config.BodyStructure{
				"success": "false",
				"message": "No response from server",
			}
			failed = true
		}

		if !failed {
			sendback = responseBody.(config.BodyStructure)
		}

		bytes, err = json.Marshal(sendback)
		if err != nil {
			sendback = config.BodyStructure{
				"success": "false",
				"message": "Error while preparing response. Please try again later.",
			}
		}

		c.JSON(200, config.BodyStructure{"data": pkg.Encrypt(string(bytes), config.Key)})

		c.Next()
	}
}
