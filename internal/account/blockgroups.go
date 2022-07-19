package account

import (
	"net/http"
	"strings"
	"time"

	"bookieguardserver/internal/helpers"
	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NewBlockGroup struct {
	Title          string `form:"title" json:"title" xml:"title"`
	TotalComputers uint   `form:"total" json:"total" xml:"total" binding:"required"`
	Plan           string `form:"plan" json:"plan" xml:"plan"`
}

func BlockGroups(c *gin.Context) {

	userId, _ := c.Get("userId")

	bg := models.BlockGroup{}
	bg.UserID = userId.(string)

	if c.Request.Method == "GET" {

		_, blockgroups := bg.ReadAll()

		for index, blockgroup := range blockgroups {
			t, _ := time.Parse("2006-01-02T00:00:00Z", blockgroup.ExpirationDate)
			blockgroups[index].ExpirationDate = t.Format("02-Jan-2006")
		}

		c.HTML(http.StatusOK, "account.blockgroups.html", gin.H{
			"title":       "Block Groups",
			"canonical":   "/account/block-groups",
			"blockgroups": blockgroups,
			"plans":       helpers.GetPlans(),
		})
	}

	if c.Request.Method == "POST" {
		var json NewBlockGroup
		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		bg.Title = json.Title
		bg.TotalComputers = json.TotalComputers
		bg.CurrentPlan = json.Plan

		// generate activation code and unlock code
		forActivation, _ := uuid.NewRandom()
		forUnlock, _ := uuid.NewRandom()

		activation := strings.Split(strings.ToUpper(forActivation.String()), "-")
		unlock := strings.Split(strings.ToUpper(forUnlock.String()), "-")

		bg.ActivationCode = activation[1] + "-" + activation[2] + "-" + activation[3]
		bg.UnlockCode = unlock[1] + "-" + unlock[2] + "-" + unlock[3]

		bg.Create()

		success, response := helpers.ProcessBlockGroupPayment(c, json.Plan, userId.(string), bg)
		if !success {
			c.String(200, response["message"])
			return
		}
		// send response
		c.HTML(200, "account.redirect.html", gin.H{"url": response["link"]})
	}
}
