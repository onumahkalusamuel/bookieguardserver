package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/helpers"
)

func Pricing(c *gin.Context) {

	plans := helpers.GetPlans()

	c.HTML(http.StatusOK, "public.pricing.html", gin.H{
		"title":     "Pricing",
		"canonical": "/pricing",
		"plans":     plans,
	})
}
