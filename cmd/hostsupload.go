package cmd

import (
	"fmt"
	"net/http"

	"github.com/onumahkalusamuel/bookieguardserver/internal"
)

func HostsUpload(resp http.ResponseWriter, req *http.Request) {

	body, err := internal.ProcessRequestBody(req)
	if err != nil {
		// return nil, err
	}
	fmt.Println(body)

	internal.SendBackResponse(resp, body)
}
