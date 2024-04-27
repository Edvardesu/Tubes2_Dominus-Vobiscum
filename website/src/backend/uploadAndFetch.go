package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"net/http"
)

const PORT = "8080"

func Ping(w http.ResponseWriter, r *http.Request) {
	answer := map[string]interface{}{
		"messageType": "S",
		"message":     "",
		"data":        "PONG",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(answer)
}

func FetchIDSResults(w http.ResponseWriter, r *http.Request) {
	
}

func UploadTextBFS(w http.ResponseWriter, r *http.Request) {
	// Define a struct to match the expected input
	type RequestData struct {
		Start       string `json:"start"`
		Destination string `json:"destination"`
	}

	// Create an instance of the struct
	var data RequestData

	// Decode the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// If there is an error in decoding, return an error message
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If there is space in string, change to underscore
	data.Start = strings.ReplaceAll(data.Start, " ", "_")
	data.Destination = strings.ReplaceAll(data.Destination, " ", "_")

	// You can now use data.Start and data.Destination in your application logic
	// For example, log the received data (you could also implement any logic needed)
	fmt.Printf("Received start: %s, destination: %s\n", data.Start, data.Destination)

	// Send a success response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"messageType": "S",
		"message":     "Data received successfully",
	})

	// BFS(data.Start, data.Destination)
}

func UploadTextIDS(w http.ResponseWriter, r *http.Request) {
	// Define a struct to match the expected input
	type RequestData struct {
		Start       string `json:"start"`
		Destination string `json:"destination"`
	}

	// Create an instance of the struct
	var data RequestData

	// Decode the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// If there is an error in decoding, return an error message
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// If there is space in string, change to underscore
	data.Start = strings.ReplaceAll(data.Start, " ", "_")
	data.Destination = strings.ReplaceAll(data.Destination, " ", "_")

	// You can now use data.Start and data.Destination in your application logic
	// For example, log the received data (you could also implement any logic needed)
	fmt.Printf("Received start: %s, destination: %s\n", data.Start, data.Destination)

	// Send a success response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"messageType": "S",
		"message":     "Data received successfully",
		// "pathFound": 
	})

	IDS(data.Start, data.Destination)
}



