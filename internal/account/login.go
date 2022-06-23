package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	message, _ := c.GetQuery("message")

	c.HTML(http.StatusOK, "account.login.html", gin.H{
		"title":     "Login",
		"canonical": "/account/login/",
		"message":   message,
	})
}
