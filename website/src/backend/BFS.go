package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "net/http"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	// "github.com/gorilla/mux"

)

// const PORT = "8080"

type Node struct {
    Link  string `json:"Link"`
    Depth int    `json:"Depth"`
}

type Request struct {
    Start       string `json:"start"`
    Destination string `json:"destination"`
}

type Response struct {
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
    Paths   [][]string `json:"paths"`
    TotalLinks int      `json:"totalLinks"`
    ExecutionTime int64 `json:"executionTime"`
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
    var req Request
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Log for debugging
    log.Printf("Received Start: %s, Destination: %s", req.Start, req.Destination)

    // You should integrate BFS here, for simplicity we return a placeholder response
    response := Response{
        Message: "Received start and destination.",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}

var start Node
var destination string
var path [][]string
var total_link int

func isInNode(lis []Node, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i].Link) {return true}
	}
	return false
}

func isInPath(mat [][]string, lis []string) bool {
	for i := len(mat)-1; i >= 0; i-- {
		if reflect.DeepEqual(mat[i],lis) {
			return true
		}
	}
	return false
}

func validasiLinkBFS(queue *[]Node, m map[string]string, x *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	domain := "https://en.wikipedia.org/wiki/"
	
	x.Lock()
	l := (*queue)[0].Link
	current := (*queue)[0].Depth + 1
	*queue = (*queue)[1:]
	x.Unlock()

	c.OnError(func (e *colly.Response, err error) {
		log.Println("Terjadi error: ", err)
	})
	
	c.OnHTML("a[href^='/wiki']", func(e *colly.HTMLElement) {
		child := e.Attr("href")[6:]
		if !(strings.Contains(child, "File:") || strings.Contains(child, "Category:") || strings.Contains(child, "Help:") || strings.Contains(child, "Wikipedia:")) {
			total_link += 1
			// cek apakah Link mengandung #
			n := strings.Index(child, "#")
			if n != -1 {
				child = cutLink(child,n)
			}
			
			// cek apakah Link ada di array URL atau di map
			x.Lock()
			_, found := m[child]
			if !isInNode(*queue, child) && !found {
				if child == destination {
					tempPath := makePath(m, l, child)
					if !isInPath(path, tempPath) {
						path = append(path, makePath(m, l, child))
					}
				} else {
					m[child] = l
					// fixed Depth 4
					if (current < 4) {
						var tempNode Node
						tempNode.Link = child
						tempNode.Depth = current
						*queue = append(*queue, tempNode)
					}
				}
			}
			x.Unlock()
		}
	})
	
	// Start scraping
	c.Visit(domain + l)
}

func BFS(awal string, akhir string) ([][]string, int, int, int64) {
// func main() {
	fmt.Println("Start BFS")

	total_link = 0
	var wait sync.WaitGroup
	var mut sync.Mutex
	var unvisitedQueue []Node
	var minPath int
	
	// start.Link = "Joko_Widodo" // nanti ini input start
	start.Depth = 0
	// destination = "Cadmium" // nanti ini input final
	unvisitedQueue = append(unvisitedQueue, start)
	
	visitedMap := map[string]string {start.Link: "start"}
	
	// hitung waktu proses BFS
	begin := time.Now().UnixNano() / int64(time.Millisecond)
	sekarang := begin
	// Loop berhenti jika queue habis atau (waktu melebihi 2,5 menit dan ada jalur ditemukan) atau waktu melebihi 5 menit
	for (len(unvisitedQueue) > 0) && (sekarang - begin <= 150000 || len(path) < 0) && (sekarang - begin <= 300000) {
		// Cari jumlah proses yang dijalankan (maksimal 100)
		n_child := len(unvisitedQueue)
		if n_child > 200 {
			n_child = 200
		}

		// Lakukan proses BFS
		for i := 0; i < n_child; i++ {
			wait.Add(1)
			go func() {	
				validasiLinkBFS(&unvisitedQueue, visitedMap, &mut, &wait)
			}()
		}
		wait.Wait()
		sekarang = time.Now().UnixNano() / int64(time.Millisecond)
	}

	exTime := sekarang - begin
	if (len(path) > 0) {
		minPath := len(path[0])
		for i := 1; i < len(path); i++ {
			if (len(path[i]) < minPath) {
				minPath = len(path[i])
			}
		}
		fmt.Println("List of paths:")
		fmt.Println(path)
		fmt.Printf("Found %d path(s), with minimum Depth %d\n", len(path), minPath)
	} else {
		fmt.Println("No path found")
	}
	fmt.Printf("Execution time: %dm %.2fs\nLink visited: %d\n", exTime/60000, float64(exTime%60000)/1000, total_link)

	return path, minPath, total_link, exTime 
}


// func Ping(w http.ResponseWriter, r *http.Request) {
// 	answer := map[string]interface{}{
// 		"messageType": "S",
// 		"message":     "",
// 		"data":        "PONG",
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	json.NewEncoder(w).Encode(answer)
// }


// func main() {
// 	// http.HandleFunc("/path", pathHandler)
//     // log.Fatal(http.ListenAndServe(":8080", nil))

// 	// mux := http.NewServeMux()
//     // mux.HandleFunc("/path", pathHandler)
//     // handler := enableCORS(mux)
//     // log.Fatal(http.ListenAndServe(":8080", handler))

// 	mux := http.NewServeMux()
//     mux.HandleFunc("/path", pathHandler)
//     handler := enableCORS(mux)
//     log.Println("Server started on :8080")
//     log.Fatal(http.ListenAndServe(":8080", handler))


// 	// BFS("Joko_Widodo", "Cadmium")
// 	// http.HandleFunc("/startBFS", startBFSHandler)
//     // fmt.Println("Server started on :8080")
//     // log.Fatal(http.ListenAndServe(":8080", nil))
// }