package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

func BlocklistCategoriesAction(c *gin.Context) {
	category_id, _ := c.Params.Get("category_id")

	// fetch this blocklist
	bl := models.Blocklist{}
	bl.CategoryID = category_id
	bl.Read()

	// make sure theres no blocklist attached to it
	if bl.ID != "" {
		c.String(200, "Cannot delete blocklist category with blocklists")
		return
	}

	blc := models.BlocklistCategory{}
	blc.ID = category_id

	blc.Delete()

	c.Redirect(http.StatusMovedPermanently, "/admin/blocklist-categories")
	return
}
