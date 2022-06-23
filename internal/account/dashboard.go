package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "account.dashboard.html", gin.H{
		"title":     "Dashboard",
		"canonical": "/account/dashboard",
	})

}
