package main

import (
	"encoding/json"
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

func UploadTextBFS(w http.ResponseWriter, r *http.Request) {
	// Define a struct to match the expected input
	type RequestData struct {
		Start        string `json:"start"`
		Destination  string `json:"destination"`
		SinglePath   bool   `json:"single_path"`
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

	 // Call BFS with the processed start and destination
	 result := BFS(data.Start, data.Destination , data.SinglePath)

	// Prepare the JSON response
	response := map[string]interface{}{
	//  "messageType": "S",
	"message":     "This is BFS !!!",
	"paths": result.Paths,
	"single_path": result.SinglePath,
	"total_links": result.TotalLinks,
	"path_length": result.PathLength,
	"exec_time": result.DurationInMS,
	"path_amount": result.PathAmount,
	}

	 // Encode the result into JSON and send it back to the client
	 if err := json.NewEncoder(w).Encode(response); err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
	 }

	// BFS(data.Start, data.Destination)
}

func UploadTextIDS(w http.ResponseWriter, r *http.Request) {
	// Define a struct to match the expected input
	type RequestData struct {
		Start        string `json:"start"`
		Destination  string `json:"destination"`
		SinglePath   bool   `json:"single_path"`
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

	 // Call IDS with the processed start and destination
	 result := IDS(data.Start, data.Destination , data.SinglePath)

	 // Prepare the JSON response
	 response := map[string]interface{}{
		//  "messageType": "S",
		"message":     "This is IDS !!!!",
		"paths": result.Paths,
		"single_path": result.SinglePath,
		"total_links": result.TotalLinks,
		"path_length": result.PathLength,
		"exec_time": result.DurationInMS,
		"path_amount": result.PathAmount,
	 }
 
	 // Set content type and write the status code
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
 
	 // Encode the result into JSON and send it back to the client
	 if err := json.NewEncoder(w).Encode(response); err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
	 }
}