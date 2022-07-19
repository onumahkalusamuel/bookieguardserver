package account

import (
	"net/http"
	"strings"

	"bookieguardserver/internal/models"

	"github.com/gin-gonic/gin"
)

func BlockGroup(c *gin.Context) {
	userId, _ := c.Get("userId")
	blockgroup_id, _ := c.Params.Get("blockgroup_id")

	// fetch this blockgroup
	bg := models.BlockGroup{}
	bg.ID = blockgroup_id
	bg.UserID = userId.(string)
	bg.Read()

	// make sure its the owner
	if bg.Title == "" {
		return
	}

	if c.Request.Method == "GET" {

		var blocklistHolder []models.Blocklist
		var allowlistHolder []models.Allowlist
		var allowlistHolder2 []string

		// fetch all blocklists
		b := models.Blocklist{}
		_, blocklists := b.ReadAllFull()

		// fetch all allowlists
		a := models.Allowlist{BlockGroupID: bg.ID}
		_, allowlistHolder = a.ReadAll()

		// put allowlist.website into a string concatenation
		for _, allowlist := range allowlistHolder {
			allowlistHolder2 = append(allowlistHolder2, allowlist.Website)
		}

		allowed := strings.Join(allowlistHolder2, ",")

		// loop through blocklists
		for _, blocklist := range blocklists {
			// check if website exists in allowlist and skip if it does
			if strings.Contains(allowed, blocklist.Website) {
				continue
			}

			// add blocklist to blocklistHolder
			blocklistHolder = append(blocklistHolder, blocklist)
		}

		// filter out any blocklist that is in allowed list
		c.HTML(http.StatusOK, "account.blockgroup.html", gin.H{
			"title":      "Block Groups",
			"canonical":  "/account/block-groups/",
			"blockgroup": bg,
			"blocklist":  blocklistHolder,
			"allowlist":  allowlistHolder,
		})

		return
	}

	if c.Request.Method == "POST" {

		website := c.PostForm("website")

		if website != "" {
			allowlist := models.Allowlist{}
			allowlist.BlockGroupID = bg.ID
			allowlist.Website = website

			if a := allowlist.Create(); a != nil {
				c.String(200, "Unable to create allowlist. Please try again later.")
				return
			}

		}

		c.Request.Method = "GET"
		c.Redirect(http.StatusMovedPermanently, "/account/block-groups/"+blockgroup_id)

	}
}
