package admin

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func ContactAction(c *gin.Context) {
	action, _ := c.Params.Get("action")
	actionId, _ := c.Params.Get("action_id")

	// fetch this blockgroup
	cont := models.Contact{}
	cont.ID = actionId
	cont.Read()

	// make sure its there
	if cont.ID == "" {
		return
	}

	if action == "delete" {
		cont.Delete()
		c.Redirect(http.StatusMovedPermanently, "/admin/contact/")
		return
	}

	if action == "mark-as-read" {
		// add to allowlist
		cont.UpdateSingle("read_status", 1)
		c.Redirect(http.StatusMovedPermanently, "/admin/contact/")
		return
	}
}
