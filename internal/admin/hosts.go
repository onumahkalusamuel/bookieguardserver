package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

func Hosts(c *gin.Context) {

	// fetch all hosts
	h := models.Host{}
	_, hosts := h.ReadAll()

	// return
	c.HTML(http.StatusOK, "admin.hosts.html", gin.H{
		"title":     "Hosts",
		"canonical": "/admin/hosts",
		"hosts":     hosts,
	})
}
