package account

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"bookieguardserver/internal/helpers"
	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		_token := reset[1] + "-" + reset[2]

		user.UpdateSingle("token", _token)

		var passwordResetLink = helpers.GetPasswordResetLink(c.Request.Host, _token)

		var message = "Hello " + user.Name + "<br/><br/>," +
			"Here is your password reset link. <a href='" + passwordResetLink + "'>" + passwordResetLink + "</a>.<br/><br/>" +
			"If you did not request for a password reset, please ignore this email.<br/><br/>" +
			"Thank you. <br/><br/>" +
			"Support Team,<br/>" +
			"<strong>Bookie Guard.</strong>"

		err := helpers.SendEmail("support", user.Email, user.Name+", Your Password Reset Link - Bookie Guard Support", message)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.Redirect(http.StatusMovedPermanently, "/account/forgot-password?successmessage="+url.QueryEscape("a reset link has been sent to your email."))
	}
}
