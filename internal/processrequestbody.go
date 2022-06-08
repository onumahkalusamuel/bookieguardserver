package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/onumahkalusamuel/bookieguardserver/config"
	"github.com/onumahkalusamuel/bookieguardserver/pkg"
)

func ProcessRequestBody(req *http.Request) (config.BodyStructure, error) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return config.BodyStructure{}, fmt.Errorf(err.Error())
	}
	var unm config.BodyStructure

	err = json.Unmarshal(body, &unm)
	if err != nil {
		return config.BodyStructure{}, fmt.Errorf(err.Error())
	}

	data := pkg.Decrypt(unm["data"], config.Key)

	var holder config.BodyStructure

	json.Unmarshal([]byte(data), &holder)

	return holder, nil
}
