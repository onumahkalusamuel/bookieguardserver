package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HowItWorks(c *gin.Context) {
	c.HTML(http.StatusOK, "public.how-it-works.html", gin.H{
		"title":     "How It Works",
		"canonical": "/how-it-works",
	})
}
