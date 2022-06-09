package cmd

import (
	"fmt"
	"net/http"

	"github.com/onumahkalusamuel/bookieguardserver/internal"
)

func Update(resp http.ResponseWriter, req *http.Request) {

	// there are basically thef following types of update

	// admin panel update (design)

	// blocklist update (based on user settings and general settings)

	// program update

	// from the requestbody, we'll determine whose system is requesting for update
	// basically through the hashedID

	// then we'll check for program version

	body, err := internal.ProcessRequestBody(req)
	if err != nil {
		// return nil, err
	}
	fmt.Println(body)

	internal.SendBackResponse(resp, body)
}
