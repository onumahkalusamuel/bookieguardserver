package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(http.StatusOK, "public.index.html", gin.H{
		"title":     "Home",
		"canonical": "/",
	})
}
