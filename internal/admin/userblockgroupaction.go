package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserBlockGroupAction(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.dashboard.html", gin.H{})
}
