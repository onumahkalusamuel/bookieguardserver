package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Downloads(c *gin.Context) {

	c.HTML(http.StatusOK, "public.downloads.html", gin.H{
		"title":     "Downloads",
		"canonical": "/downloads",
	})
}
