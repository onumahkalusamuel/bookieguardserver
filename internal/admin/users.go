package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type NewUserData struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Email    string `form:"email" json:"email" xml:"email"  binding:"required"`
	Phone    string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Users(c *gin.Context) {

	if c.Request.Method == "GET" {

		// fetch all users
		b := models.User{}
		_, users := b.ReadAll()

		// filter out any blocklist that is in allowed list
		c.HTML(http.StatusOK, "admin.users.html", gin.H{
			"title":     "Users",
			"canonical": "/admin/users",
			"users":     users,
		})

		return
	}

	if c.Request.Method == "POST" {

		var newUser NewUserData

		if err := c.ShouldBind(&newUser); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), 4)

		user := models.User{
			Name:     newUser.Name,
			Email:    newUser.Email,
			Phone:    newUser.Phone,
			Password: string(hash),
		}

		user.Create()

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/admin/users")
	}
}
