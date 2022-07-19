package account

import (
	"bookieguardserver/internal/helpers"
	"bookieguardserver/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopUp struct {
	Plan string `form:"plan" json:"plan" xml:"plan"`
}

func BlockGroupTopUp(c *gin.Context) {
	userId, _ := c.Get("userId")
	blockgroup_id, _ := c.Params.Get("blockgroup_id")

	// fetch this blockgroup
	bg := models.BlockGroup{}
	bg.ID = blockgroup_id
	bg.UserID = userId.(string)
	bg.Read()

	// make sure its the owner
	if bg.Title == "" {
		return
	}

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "account.topup.html", gin.H{
			"title":      "Block Groups",
			"canonical":  "/account/block-groups",
			"blockgroup": bg,
			"plans":      helpers.GetPlans(),
		})
	}
	if c.Request.Method == "POST" {
		var json TopUp
		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		success, response := helpers.ProcessBlockGroupPayment(c, json.Plan, userId.(string), bg)

		if !success {
			c.String(200, response["message"])
			return
		}
		// send response
		c.HTML(200, "account.redirect.html", gin.H{"url": response["link"]})
	}

}
