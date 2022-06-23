package account

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/helpers"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"github.com/onumahkalusamuel/bookieguardserver/services/paystack"
)

// Binding from JSON
type QueryData struct {
	TrxRef    string `form:"trxref" json:"trxref" xml:"trxref"  binding:"required"`
	Reference string `form:"reference" json:"reference" xml:"reference" binding:"required"`
}

func PaystackCallBack(c *gin.Context) {

	var query QueryData

	if err := c.ShouldBindQuery(&query); err != nil {
		c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
		return
	}

	// make a call to paystack to get details
	verify := paystack.VerifyPayment(query.Reference)

	if verify["success"] == "false" {
		c.String(200, verify["message"])
		return
	}

	// it went through, process now
	blockgroup := models.BlockGroup{}
	blockgroup.ID = verify["blockGroupID"]
	blockgroup.UserID = verify["userID"]
	blockgroup.Read()

	payment := models.Payment{}
	payment.PaymentReference = verify["paymentReference"]
	payment.BlockGroupID = verify["blockGroupID"]
	payment.UserID = verify["userID"]
	payment.Read()

	if payment.Status != "pending" {
		c.String(200, "Payment processing already completed. Status: "+payment.Status)
		return
	}

	// check if correct amount was paid for
	converted, _ := strconv.Atoi(verify["amount"])
	amountToUnit := converted * 100
	if uint(amountToUnit) < payment.Amount {
		c.String(200, "Invalid amount detected. Please contact admin with this reference code: "+payment.PaymentReference)
		return
	}

	// update payment
	if err := payment.UpdateSingle("status", "success"); err != nil {
		c.String(200, "Unable to save payment details at the moment. Please try again later.")
		return
	}

	// update blockgroup
	blockgroup.UpdateSingle("current_plan", payment.PlanID)

	// calculate expiry date
	var selectedPlan helpers.Plan
	plans := helpers.GetPlans()

	for _, plan := range plans {
		if plan.Key == payment.PlanID {
			selectedPlan = plan
			break
		}
	}

	if selectedPlan.Key == "" {
		c.String(200, "An internal server error occured. Please contact admin.")
		return
	}

	extent := 30 * selectedPlan.Duration * 24 * uint(time.Hour)
	expirateDate := time.Now().Add(time.Duration(extent))

	blockgroup.UpdateSingle("expiration_date", expirateDate.Format("2006-01-02"))

	c.String(200, "Payment was processed successfully. You can go back to your dashboard now.")
	return

}
