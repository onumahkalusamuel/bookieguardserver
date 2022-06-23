package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AccountAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)

		if session.Get("userType") != "account" {
			c.Redirect(301, "/account/login")
			c.Abort()
		}

		c.Set("userId", session.Get("userId"))
		c.Set("email", session.Get("email"))
		c.Set("userType", session.Get("userType"))

		c.Next()
	}
}
