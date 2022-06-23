package account

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type LoginData struct {
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func LoginPost(c *gin.Context) {

	var json LoginData

	if err := c.ShouldBind(&json); err != nil {
		c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	user := models.User{}
	user.Email = json.Email
	user.Read()

	if user.ID == "" {
		c.Redirect(http.StatusMovedPermanently, "/account/login?message="+url.QueryEscape("invalid login details."))
		return
	}

	// validate password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/account/login?message="+url.QueryEscape("invalid login details."))
		return
	}

	// set session
	session := sessions.Default(c)
	session.Options(sessions.Options{Path: "/"})
	session.Set("userId", user.ID)
	session.Set("userType", user.UserType)
	session.Set("email", user.Email)
	session.Save()

	c.Request.Method = "GET"
	c.Redirect(http.StatusMovedPermanently, "/account/dashboard")
}
