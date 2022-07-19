package admin

import (
	"net/http"
	"strings"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

type SettingsData struct {
	Setting string `form:"setting" json:"setting" xml:"setting" binding:"required"`
	Value   string `form:"value" json:"value" xml:"value" binding:"required"`
}

func Settings(c *gin.Context) {
	if c.Request.Method == "GET" {

		// fetch all
		bb := models.Settings{}
		_, settings := bb.ReadAll()

		// return
		c.HTML(http.StatusOK, "admin.settings.html", gin.H{
			"title":     "Settings",
			"canonical": "/admin/settings",
			"settings":  settings,
		})

		return
	}

	if c.Request.Method == "POST" {

		var newData SettingsData

		if err := c.ShouldBind(&newData); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		settings := models.Settings{Setting: strings.Trim(newData.Setting, " ")}
		settings.Read()

		if settings.ID != "" {
			settings.UpdateSingle("value", newData.Value)
		}

		if settings.ID == "" {
			settings.Value = newData.Value
			settings.Create()
		}

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/settings")
	}
}
