// A very basic go hello-world app
// for use in the minikube stage 1 example
// for Equal Experts.

// Authored: John Mylchreest 21/02/2019

package main

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
)

// Create response body struct
type helloWorldResponse struct {
    Message	 string `json:"message"`
    Hostname string `json:"hostname"`
}

func main() {
	// Listening socket
	port := ":8081"

	// Setup http router for /
	http.HandleFunc("/", RootHandler)

	// Launch and listen
	fmt.Printf("Listening on %s", port)
	http.ListenAndServe(port, nil)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch response body
	resp, err := rootResponse()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.WriteHeader(http.StatusForbidden)  // uncomment to force test fail
	w.Write(resp)
}

func rootResponse()([]byte, error) {
	// get hostname
	name, _ := os.Hostname()
	
	// set message
	message := "Hello world"

	// prepare response
	resp := helloWorldResponse{message, name}

	// Return response payload
    return json.Marshal(resp)
}