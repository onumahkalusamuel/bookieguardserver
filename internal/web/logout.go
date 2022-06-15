package web

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout",
	})
}
