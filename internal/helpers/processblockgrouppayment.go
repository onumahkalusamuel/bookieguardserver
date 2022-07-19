package helpers

import (
	"bookieguardserver/config"
	"bookieguardserver/internal/models"
	"bookieguardserver/services/paystack"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessBlockGroupPayment(c *gin.Context, planKey string, userID string, blockGroup models.BlockGroup) (bool, config.BodyStructure) {
	// fetch the plans
	var selectedPlan Plan
	plans := GetPlans()

	for _, plan := range plans {
		if plan.Key == planKey {
			selectedPlan = plan
			break
		}
	}

	// check if plan was captured
	if selectedPlan.Key == "" {
		return false, config.BodyStructure{"message": "Unable to find selected plan."}
	}

	var amount uint

	if selectedPlan.Key == "plan5" {
		amount = selectedPlan.Price * blockGroup.TotalComputers * 100
	} else {
		amount = selectedPlan.Price * selectedPlan.Duration * blockGroup.TotalComputers * 100
	}

	user := models.User{}
	user.ID = userID
	user.Read()

	pR, _ := uuid.NewRandom()
	paymentReference := pR.String()

	paymentLink := paystack.CreatePaymentLink(map[string]any{
		"amount":       amount,
		"currency":     config.PaystackCurrency,
		"email":        user.Email,
		"reference":    paymentReference,
		"callback_url": paystack.GetCallbackURL(c.Request.Host),
		"channel":      config.PaystackChannels,
		"metadata": config.PaystackMetaData{
			BlockGroupID:     blockGroup.ID,
			UserID:           user.ID,
			PaymentReference: paymentReference,
		},
	})

	if paymentLink["success"] != "true" {
		return false, config.BodyStructure{"message": paymentLink["message"]}
	}

	// save the payment details
	payment := models.Payment{}
	payment.UserID = user.ID
	payment.BlockGroupID = blockGroup.ID
	payment.PaymentReference = paymentReference
	payment.Amount = amount / 100
	payment.Currency = "NGN"
	payment.PlanID = selectedPlan.Key
	payment.Details = paymentLink["link"]
	payment.Gateway = "paystack"
	payment.Create()

	return true, config.BodyStructure{"link": paymentLink["link"]}
}
