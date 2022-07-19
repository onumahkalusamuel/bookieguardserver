package admin

import (
	"bookieguardserver/internal/helpers"
	"bookieguardserver/internal/models"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserBlockGroups(c *gin.Context) {
	userId, _ := c.Params.Get("user_id")

	bg := models.BlockGroup{}
	bg.UserID = userId

	user := models.User{}
	user.ID = userId
	user.Read()

	if c.Request.Method == "GET" {

		_, blockgroups := bg.ReadAll()

		for index, blockgroup := range blockgroups {
			t, _ := time.Parse("2006-01-02T00:00:00Z", blockgroup.ExpirationDate)
			blockgroups[index].ExpirationDate = t.Format("02-Jan-2006")
		}

		c.HTML(http.StatusOK, "admin.userblockgroups.html", gin.H{
			"title":       "Users",
			"canonical":   "/admin/users",
			"user":        user,
			"blockgroups": blockgroups,
			"plans":       helpers.GetPlans(),
		})
	}

	if c.Request.Method == "POST" {

		// set session
		session := sessions.Default(c)
		session.Options(sessions.Options{Path: "/"})
		session.Set("userId", user.ID)
		session.Set("userType", user.UserType)
		session.Set("email", user.Email)
		session.Save()

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/account/dashboard")
	}
}
