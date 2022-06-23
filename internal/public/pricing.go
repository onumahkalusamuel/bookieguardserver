package public

import (
	"net/http"

	"bookieguardserver/internal/helpers"

	"github.com/gin-gonic/gin"
)

func Pricing(c *gin.Context) {

	plans := helpers.GetPlans()

	c.HTML(http.StatusOK, "public.pricing.html", gin.H{
		"title":     "Pricing",
		"canonical": "/pricing",
		"plans":     plans,
	})
}
