package account

import (
	"net/http"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func BlockGroupSettingsAction(c *gin.Context) {
	userId, _ := c.Get("userId")
	blockgroup_id, _ := c.Params.Get("blockgroup_id")
	action, _ := c.Params.Get("action")
	actionId, _ := c.Params.Get("action_id")

	// fetch this blockgroup
	bg := models.BlockGroup{}
	bg.ID = blockgroup_id
	bg.UserID = userId.(string)
	bg.Read()

	// make sure its the owner
	if bg.Title == "" {
		return
	}

	if action == "delete" {
		// delete from allowlist
		allowlist := models.Allowlist{}
		allowlist.ID = actionId
		allowlist.BlockGroupID = blockgroup_id

		allowlist.Delete()

		c.Redirect(http.StatusMovedPermanently, "/account/block-groups/"+blockgroup_id+"/settings")
		return
	}

	if action == "allow" {
		// add to allowlist
		blocklist := models.Blocklist{}
		blocklist.ID = actionId
		blocklist.Read()

		allowlist := models.Allowlist{}
		allowlist.Website = blocklist.Website
		allowlist.BlockGroupID = blockgroup_id

		allowlist.Create()

		c.Redirect(http.StatusMovedPermanently, "/account/block-groups/"+blockgroup_id+"/settings")
		return
	}
}
