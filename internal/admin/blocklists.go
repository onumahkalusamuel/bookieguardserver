package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

type NewBlocklistData struct {
	Website    string `form:"website" json:"website" xml:"website"  binding:"required"`
	CategoryID string `form:"category_id" json:"category_id" xml:"category_id"  binding:"required"`
}

func Blocklists(c *gin.Context) {
	if c.Request.Method == "GET" {

		// fetch all blocklists
		b := models.Blocklist{}
		_, blocklists := b.ReadAllFull()

		// fetch all categories
		bb := models.BlocklistCategory{}
		_, blocklistcategories := bb.ReadAll()

		// return
		c.HTML(http.StatusOK, "admin.blocklists.html", gin.H{
			"title":               "Blocklists",
			"canonical":           "/admin/blocklists",
			"blocklists":          blocklists,
			"blocklistcategories": blocklistcategories,
		})

		return
	}

	if c.Request.Method == "POST" {

		var newData NewBlocklistData

		if err := c.ShouldBind(&newData); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		blocklist := models.Blocklist{
			Website:    newData.Website,
			CategoryID: newData.CategoryID,
		}

		blocklist.Create()

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/blocklists")
	}
}
