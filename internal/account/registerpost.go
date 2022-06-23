package account

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type RegisterData struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email" binding:"required"`
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func RegisterPost(c *gin.Context) {

	var json RegisterData

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user exists
	user := models.User{}
	user.Email = json.Email

	_ = user.Read()

	if user.ID != "" {
		c.JSON(200, gin.H{
			"success": false,
			"message": "email already in use.",
		})
		return
	}

	// set remaining params
	user.Name = json.Name

	hash, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 4)
	user.Password = string(hash)

	user.Phone = json.Phone

	if err := user.Create(); err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "an error occured. please try again later.",
		})
		return
	}

	c.JSON(200, gin.H{
		"success":  true,
		"message":  "account created successfully.",
		"redirect": "/account/login",
	})
}
