package cmd

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/onumahkalusamuel/bookieguardserver/internal"
)

func Activation(resp http.ResponseWriter, req *http.Request) {

	requestBody, err := internal.ProcessRequestBody(req)
	if err != nil {
		return
	}

	socialnetworks := []string{"facebook", "linkedin", "twitter", "youtube", "pinterest", "instagram", "tumblr", "flickr", "reddit", "snapchat", "whatsapp", "quora", "tiktok", "vimeo", "bigsugar", "mix.com", "medium.com", "digg.com", "viber", "wechat"}
	others := []string{"google.com", "adservices", "googlead", "altavista", "devilfinder", "bing.com", "nairaland", "livescore", "futbol24", "predictz", "forebet", "footballpredictions", "betensured", "tips180", "victorpredict", "stakegains", "freesupertips", "a;"}
	bookies := []string{"bet9ja", "msport", "sportybet", "betking", "1xbet", "betway", "betwinner", "bet365", "22bet", "nairabet", "melbet", "wazobet", "parimatch", "betbonanza", "bigibet", "merrybet", "paripesa", "winnersgoldenbet", "lionsbet", "accessbet", "naijabet", "surebet247", "1960bet", "betfair", "betwinner", "888sport", "betway", "netbet", "betbiga", "wazo-bet", "supabets", "betpawa", "cloudbet", "betfarm", "zebet.ng", "parimatch", "bangbet", "db-bet", "frapapa", "easybet.ng", "n1bigbet", "blackbet", "afribet", "allcitybet", "irokobet", "betwin9ja"}
	blocklist := []string{}

	for _, b := range socialnetworks {
		blocklist = append(blocklist, b)
	}
	for _, b := range others {
		blocklist = append(blocklist, b)
	}
	for _, b := range bookies {
		if strings.ToLower(requestBody["shop"]) == b {
			continue
		}
		blocklist = append(blocklist, b)
	}

	requestBody["success"] = "true"
	requestBody["activated"] = "true"
	requestBody["blocklist"] = strings.Join(blocklist, ",")
	requestBody["expirationDate"] = time.Now().String()

	send, err := internal.SendBackResponse(resp, requestBody)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	if send {

	}

}
