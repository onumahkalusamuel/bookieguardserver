package account

import (
	"net/http"
	"strings"

	"bookieguardserver/config"
	"bookieguardserver/internal/helpers"
	"bookieguardserver/internal/models"
	"bookieguardserver/services/paystack"

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

		// fetch the plans
		var selectedPlan helpers.Plan
		plans := helpers.GetPlans()

		for _, plan := range plans {
			if plan.Key == json.Plan {
				selectedPlan = plan
				break
			}
		}

		// check if plan was captured
		if selectedPlan.Key == "" {
			c.String(200, "Unable to find selected plan.")
			return
		}

		amount := selectedPlan.Price * json.TotalComputers * 100

		user := models.User{}
		user.ID = userId.(string)
		user.Read()

		pR, _ := uuid.NewRandom()
		paymentReference := pR.String()

		paymentLink := paystack.CreatePaymentLink(map[string]any{
			"amount":       amount,
			"currency":     config.PaystackCurrency,
			"email":        user.Email,
			"reference":    paymentReference,
			"callback_url": config.PaystackCallBackURL,
			"channel":      config.PaystackChannels,
			"metadata": config.PaystackMetaData{
				BlockGroupID:     bg.ID,
				UserID:           user.ID,
				PaymentReference: paymentReference,
			},
		})

		if paymentLink["success"] != "true" {
			c.String(200, paymentLink["message"])
			return
		}

		// save the payment details
		payment := models.Payment{}
		payment.UserID = user.ID
		payment.BlockGroupID = bg.ID
		payment.PaymentReference = paymentReference
		payment.Amount = amount / 100
		payment.Currency = "NGN"
		payment.PlanID = selectedPlan.Key
		payment.Details = paymentLink["link"]
		payment.Gateway = "paystack"
		payment.Create()

		// send response
		c.HTML(200, "account.redirect.html", gin.H{"url": paymentLink["link"]})
	}
}
