package admin

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func BlocklistAction(c *gin.Context) {
	blocklist_id, _ := c.Params.Get("blocklist_id")

	// fetch this blocklist
	bl := models.Blocklist{}
	bl.ID = blocklist_id
	bl.Read()

	// make sure theres no blocklist attached to it
	if bl.ID == "" {
		c.String(http.StatusNotFound, "Not found")
	}

	bl.Delete()

	c.Redirect(http.StatusMovedPermanently, "/admin/blocklists")
	return
}
