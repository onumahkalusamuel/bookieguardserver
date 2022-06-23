package admin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "admin.login.html", gin.H{
			"title":     "Login",
			"canonical": "/admin/login",
		})
	}

	if c.Request.Method == "POST" {

		var json LoginData

		if err := c.ShouldBind(&json); err != nil {
			c.String(http.StatusBadRequest, "%v", gin.H{"error": err.Error()})
			return
		}

		// check if user exists

		if json.Username != "adminuser" {
			c.String(http.StatusBadRequest, "invalid login details.")
			return
		}

		if json.Password != "adminuser" {
			c.String(http.StatusBadRequest, "invalid login details.")
			return
		}

		// set session
		session := sessions.Default(c)
		session.Options(sessions.Options{Path: "/"})
		session.Set("userId", "adminuser")
		session.Set("userType", "adminuser")
		session.Set("email", "adminuser@example.com")
		session.Save()

		c.Request.Method = "GET"

		c.Redirect(http.StatusMovedPermanently, "/admin/dashboard")

	}
}
