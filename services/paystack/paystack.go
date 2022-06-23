package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"bookieguardserver/config"
)

const initURL = "https://api.paystack.co/transaction/initialize"
const verifyURL = "https://api.paystack.co/transaction/verify/%v"

var headers = config.BodyStructure{
	"Authorization": "Bearer " + config.PaystackSecretKey,
	"Content-Type":  "application/json",
}

var client = &http.Client{}

func CreatePaymentLink(params map[string]any) config.BodyStructure {

	marshaled, _ := json.Marshal(params)

	req, _ := http.NewRequest("POST", initURL, bytes.NewBuffer(marshaled))

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, _ := client.Do(req)

	var res any

	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return config.BodyStructure{
			"success": "false",
			"message": err.Error(),
		}
	}

	decoded := res.(map[string]any)

	if decoded["status"] == false {
		return config.BodyStructure{
			"success": "false",
			"message": fmt.Sprint(decoded["message"]),
		}
	}

	d := decoded["data"].(map[string]interface{})

	paymentLink := d["authorization_url"]

	return config.BodyStructure{
		"success": "true",
		"link":    paymentLink.(string),
	}
}

func VerifyPayment(reference string) config.BodyStructure {

	url := fmt.Sprintf(verifyURL, reference)

	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

	req.Header.Set("Authorization", headers["Authorization"])

	resp, _ := client.Do(req)

	var res any

	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return config.BodyStructure{
			"success": "false",
			"message": err.Error(),
		}
	}

	decoded := res.(map[string]any)

	if decoded["status"] == false {
		return config.BodyStructure{
			"success": "false",
			"message": fmt.Sprint(decoded["message"]),
		}
	}

	d := decoded["data"].(map[string]interface{})

	// check if it failed
	if d["status"].(string) != "success" {
		return config.BodyStructure{
			"success": "false",
			"message": d["gateway_response"].(string),
		}
	}

	metadata := d["metadata"].(map[string]interface{})

	if metadata["UserID"].(string) == "" ||
		metadata["BlockGroupID"].(string) == "" ||
		metadata["PaymentReference"].(string) == "" {
		return config.BodyStructure{
			"success": "false",
			"message": "Vital details missing. Please contact admin.",
		}
	}

	// successful response
	return config.BodyStructure{
		"success":          "true",
		"amount":           fmt.Sprint(d["amount"]),
		"userID":           metadata["UserID"].(string),
		"blockGroupID":     metadata["BlockGroupID"].(string),
		"paymentReference": metadata["PaymentReference"].(string),
	}
}
