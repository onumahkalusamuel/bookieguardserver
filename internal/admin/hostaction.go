package admin

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func HostAction(c *gin.Context) {
	host_id, _ := c.Params.Get("host_id")
	action, _ := c.Params.Get("action")

	// fetch this host
	host := models.Host{}
	host.ID = host_id
	host.Read()

	// make sure it exits
	if host.Website == "" || action == "" {
		c.String(200, "host not found")
		return
	}

	if action == "addtoblocklist" {
		// add to blocklist
		category := models.BlocklistCategory{}
		category.Title = "other"
		category.Read()

		blocklist := models.Blocklist{}
		blocklist.CategoryID = category.ID
		blocklist.Website = host.Website
		if blocklist.Create() == nil {
			host.Delete()
		}
	}

	host.Delete()

	c.Redirect(http.StatusMovedPermanently, "/admin/hosts")
	return
}
