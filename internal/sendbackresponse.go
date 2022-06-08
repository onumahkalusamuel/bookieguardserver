package internal

import (
	"encoding/json"
	"net/http"

	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/pkg"
)

func SendBackResponse(resp http.ResponseWriter, responseBody config.BodyStructure) (bool, error) {

	m, err := json.Marshal(responseBody)
	if err != nil {
		return false, err
	}

	sendback := config.BodyStructure{"data": pkg.Encrypt(string(m), config.Key)}

	resp.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(resp).Encode(&sendback)
	if err != nil {
		return false, err
	}

	return true, nil
}
