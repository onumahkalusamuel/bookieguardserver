package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "account.register.html", gin.H{
		"title":     "Register",
		"canonical": "/account/register/",
	})
}
