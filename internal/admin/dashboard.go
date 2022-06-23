package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.dashboard.html", gin.H{
		"title":     "Dashboard",
		"canonical": "/admin/dashboard",
	})
}
