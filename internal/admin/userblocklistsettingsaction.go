package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserBlockGroupSettingsAction(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.dashboard.html", gin.H{})
}
