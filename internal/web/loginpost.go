package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Binding from JSON
type LoginData struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginPost(c *gin.Context) {

	var json LoginData

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// process form
	c.JSON(200, gin.H{
		"message": "Login",
	})
}
