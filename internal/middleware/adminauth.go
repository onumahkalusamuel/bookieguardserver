package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)

		if session.Get("userType") != "adminuser" {
			c.Redirect(301, "/admin/login")
			c.Abort()
		}

		c.Set("userId", session.Get("userId"))
		c.Set("email", session.Get("email"))
		c.Set("userType", session.Get("userType"))

		c.Next()
	}
}
