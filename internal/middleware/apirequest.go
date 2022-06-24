package middleware

import (
	"encoding/json"
	"os"

	"bookieguardserver/config"
	"bookieguardserver/pkg"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type RequestData struct {
	Data string `json:"data" binding:"required"`
}

func ApiRequest() gin.HandlerFunc {

	return func(c *gin.Context) {

		var body RequestData
		var requestBody config.BodyStructure

		if err := c.ShouldBindJSON(&body); err != nil {
			c.Set("requestBody", requestBody)
			c.Next()
			return
		}

		data := pkg.Decrypt(body.Data, os.Getenv("APP_DECRYPT_KEY"))

		json.Unmarshal([]byte(data), &requestBody)

		if len(requestBody["hashedID"]) < 20 {
			c.Set("requestBody", config.BodyStructure{})
			c.Next()
			return
		}

		c.Set("requestBody", requestBody)
		c.Next()
	}
}
