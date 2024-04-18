package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func runGinMain(wg *sync.WaitGroup) {
	defer wg.Done()
	runGin()
}

func runMuxMain(wg *sync.WaitGroup) {
	defer wg.Done()
	r := mux.NewRouter()

	r.HandleFunc("/ping", Ping).Methods("GET")
	r.HandleFunc("/upload", UploadImages).Methods("POST")

	log.Printf("Server is running on http://localhost:%s", PORT)
	log.Println(http.ListenAndServe(":"+PORT, r))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go runGinMain(&wg)
	go runMuxMain(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
