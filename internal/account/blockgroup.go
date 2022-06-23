package account

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func BlockGroup(c *gin.Context) {

	userId, _ := c.Get("userId")

	if c.Request.Method == "GET" {

		blockgroup_id, _ := c.Params.Get("blockgroup_id")

		// fetch this blockgroup
		bg := models.BlockGroup{}
		bg.ID = blockgroup_id
		bg.UserID = userId.(string)
		bg.Read()

		if bg.ID == "" {
			c.String(http.StatusNotFound, "Not found.")
			return
		}

		// get activated computers
		comp := models.Computer{}
		comp.BlockGroupID = bg.ID
		_, computers := comp.ReadAll()

		// get the payments
		paym := models.Payment{}
		paym.BlockGroupID = bg.ID
		paym.UserID = userId.(string)
		_, payments := paym.ReadAll()

		c.HTML(http.StatusOK, "account.blockgroup.html", gin.H{
			"title":      "Block Group",
			"canonical":  "/account/block-group",
			"blockgroup": bg,
			"computers":  computers,
			"payments":   payments,
		})
	}
}
