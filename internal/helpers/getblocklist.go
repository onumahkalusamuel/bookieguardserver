package helpers

import (
	"strings"

	"github.com/onumahkalusamuel/bookieguardserver/internal/models"
)

func GetBlockList(BlockGroupID uint) string {

	blocklistHolder := []string{}
	allowlistHolder := []string{}

	// fetch all blocklists
	b := models.Blocklist{}
	_, blocklists := b.ReadAll()

	// fetch all allowlists
	a := models.Allowlist{
		BlockGroupID: BlockGroupID,
	}
	_, allowlists := a.ReadAll()

	// put allowlist.website into a string concatenation
	for _, allowlist := range allowlists {
		allowlistHolder = append(allowlistHolder, allowlist.Website)
	}
	allowed := strings.Join(allowlistHolder, ",")

	// loop through blocklists
	for _, blocklist := range blocklists {
		// check if website exists in allowlist and skip if it does
		if strings.Contains(allowed, blocklist.Website) {
			continue
		}

		// add blocklist.Website to blocklistHolder
		blocklistHolder = append(blocklistHolder, blocklist.Website)
	}

	return strings.Join(blocklistHolder, ",")
}
