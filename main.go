package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PrintBody(w http.ResponseWriter, r *http.Request) {
	var payload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Print("Could not decode request body: is it JSON?")
		return
	}

	prettyPayload, err := json.MarshalIndent(payload, "", "    ")
	if err != nil {
		log.Print("Error marshalling JSON for printing. This probably shouldn't happen.")
	}

	log.Print("Decoded a JSON payload:")
	fmt.Println(string(prettyPayload))
}

func main() {
	http.HandleFunc("/", PrintBody)
	http.ListenAndServe(":8080", nil)
}
