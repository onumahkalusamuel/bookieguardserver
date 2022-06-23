package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {

	c.HTML(http.StatusOK, "public.contact.html", gin.H{
		"title":     "Contact",
		"canonical": "/contact",
	})
}
