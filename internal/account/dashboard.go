package account

import (
	"bookieguardserver/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	userId, _ := c.Get("userId")

	user := models.User{}
	user.ID = userId.(string)
	user.Read()

	c.HTML(http.StatusOK, "account.dashboard.html", gin.H{
		"title":     "Dashboard",
		"canonical": "/account/dashboard",
		"user":      user,
	})

}
