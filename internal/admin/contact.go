package admin

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	if c.Request.Method == "GET" {

		// fetch all blocklists
		cont := models.Contact{}
		_, contacts := cont.ReadAll()

		// return
		c.HTML(http.StatusOK, "admin.contact.html", gin.H{
			"title":     "Contact",
			"canonical": "/admin/contact",
			"contacts":  contacts,
		})

		return
	}
}
