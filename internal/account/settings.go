package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type UserSettings struct {
	Phone       string `form:"phone" json:"phone" xml:"phone"`
	OldPassword string `form:"oldpassword" json:"oldpassword" xml:"oldpassword" binding:"required"`
	NewPassword string `form:"newpassword" json:"newpassword" xml:"newpassword"`
}

func Settings(c *gin.Context) {

	user := models.User{}
	userId, _ := c.Get("userId")
	user.ID = userId.(string)
	user.Read()

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "account.settings.html", gin.H{
			"title":     "Profile Settings",
			"canonical": "/account/settings",
			"user":      user,
		})
		return
	}

	if c.Request.Method == "POST" {
		var json UserSettings
		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		// check if old password is okay
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.OldPassword))
		if err != nil {
			c.String(http.StatusBadRequest, "invalid password provided.")
			return
		}

		if json.Phone != "" {
			user.UpdateSingle("phone", json.Phone)
		}

		if json.NewPassword != "" {
			hash, _ := bcrypt.GenerateFromPassword([]byte(json.NewPassword), 4)
			user.UpdateSingle("password", string(hash))
		}

		c.String(http.StatusOK, "profile updated successfully.")
	}
}
