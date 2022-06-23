package account

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

type ForgotPasswordData struct {
	Email string `form:"email" json:"email" xml:"email"`
}

func ForgotPassword(c *gin.Context) {

	if c.Request.Method == "GET" {

		message, _ := c.GetQuery("message")
		successmessage, _ := c.GetQuery("successmessage")

		c.HTML(http.StatusOK, "account.forgot-password.html", gin.H{
			"title":          "Forgot Password",
			"canonical":      "/account/forgot-password",
			"message":        message,
			"successmessage": successmessage,
		})
	}

	if c.Request.Method == "POST" {
		var json ForgotPasswordData
		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		user := models.User{}
		user.Email = json.Email
		user.Read()

		if user.ID == "" {
			c.Redirect(http.StatusMovedPermanently, "/account/forgot-password?message="+url.QueryEscape("account not found."))
			return
		}

		// generate password token
		token, _ := uuid.NewRandom()

		reset := strings.Split(strings.ToUpper(token.String()), "-")

		user.UpdateSingle("token", reset[1]+"-"+reset[2])

		c.Redirect(http.StatusMovedPermanently, "/account/forgot-password?successmessage="+url.QueryEscape("a reset link has been sent to your email."))
	}
}
