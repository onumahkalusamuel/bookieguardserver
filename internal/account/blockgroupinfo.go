package account

import (
	"net/http"
	"time"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func BlockGroupInfo(c *gin.Context) {

	userId, _ := c.Get("userId")

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

	// blockgroup date
	t, _ := time.Parse("2006-01-02T00:00:00Z", bg.ExpirationDate)
	bg.ExpirationDate = t.Format("02-Jan-2006")

	c.HTML(http.StatusOK, "account.blockgroupinfo.html", gin.H{
		"title":      "Block Groups",
		"canonical":  "/account/block-group",
		"blockgroup": bg,
		"computers":  computers,
		"payments":   payments,
	})

}
