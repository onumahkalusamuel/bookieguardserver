package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

func UserDelete(c *gin.Context) {
	user_id, _ := c.Params.Get("user_id")

	// fetch this blocklist
	user := models.User{}
	user.ID = user_id
	user.UserType = "account"
	user.Read()

	//
	if user.ID == "" {
		c.String(http.StatusNotFound, "Not found")
	}

	user.Delete()

	c.Redirect(http.StatusMovedPermanently, "/admin/users")
	return
}
