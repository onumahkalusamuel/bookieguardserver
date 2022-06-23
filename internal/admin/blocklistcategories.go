package admin

import (
	"net/http"
	"strings"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

type NewBlocklistCategoryData struct {
	Title        string `form:"title" json:"title" xml:"title" binding:"required"`
	DisplayTitle string `form:"display_title" json:"display_title" xml:"display_title" binding:"required"`
}

func BlocklistCategories(c *gin.Context) {
	if c.Request.Method == "GET" {

		// fetch all categories
		bb := models.BlocklistCategory{}
		_, categories := bb.ReadAll()

		// return
		c.HTML(http.StatusOK, "admin.blocklistcategories.html", gin.H{
			"title":               "Blocklist Categories",
			"canonical":           "/admin/blocklist-categories",
			"blocklistcategories": categories,
		})

		return
	}

	if c.Request.Method == "POST" {

		var newData NewBlocklistCategoryData

		if err := c.ShouldBind(&newData); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		blocklistCategory := models.BlocklistCategory{
			Title:        strings.ReplaceAll(strings.ToLower(newData.Title), " ", "-"),
			DisplayTitle: newData.DisplayTitle,
		}

		blocklistCategory.Create()

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/blocklist-categories")
	}
}
