package main

import (
	"net/http"

	"github.com/onumahkalusamuel/bookieguardserver/cmd"
)

func main() {

	http.HandleFunc("/posts", cmd.Activation)
	http.HandleFunc("/hosts-upload", cmd.HostsUpload)
	http.HandleFunc("/update", cmd.Update)

	http.ListenAndServe("localhost:8888", nil)

}
