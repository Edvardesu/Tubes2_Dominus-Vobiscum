package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func runMuxMain(wg *sync.WaitGroup) {
	defer wg.Done()
    r := mux.NewRouter()

    r.HandleFunc("/ping", Ping).Methods("GET")
    r.HandleFunc("/uploadbfs", UploadTextBFS).Methods("POST")
	r.HandleFunc("/uploadids", UploadTextIDS).Methods("POST")
	// r.HandleFunc("/getids", FetchIDSResults).Methods("GET")

    // Set up CORS
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

    // Wrap your router with CORS middleware
    handler := handlers.CORS(originsOk, headersOk, methodsOk)(r)

    log.Printf("Server is running on http://localhost:%s", PORT)
    log.Println(http.ListenAndServe(":"+PORT, handler))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	// go runGinMain(&wg)
	go runMuxMain(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
