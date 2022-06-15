package internal

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/pkg"
)

func ApiResponse(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {

		data, exists := c.Get("data")
		if !exists {
			// tell them nothing here
		}

		m, err := json.Marshal(data)
		if err != nil {
			// return false, err
		}

		sendback := config.BodyStructure{"data": pkg.Encrypt(string(m), config.Key)}

		c.JSON(200, sendback)

		c.Next()
	}
}
