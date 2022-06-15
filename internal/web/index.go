package web

import (
	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

func Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"blocklist_categories": models.BlocklistCategory{},
		"blocklist":            models.Blocklist{},
		"blockgroup":           models.BlockGroup{},
		"user":                 models.User{},
		"allowlist":            models.Allowlist{},
		"computer":             models.Computer{},
		"payment":              models.Payment{},
	})
}
