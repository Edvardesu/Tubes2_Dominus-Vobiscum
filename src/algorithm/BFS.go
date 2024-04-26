package main

import (
	"fmt"
	"strings"
	"reflect"
	"log"
	"sync"
	"time"
	"github.com/gocolly/colly/v2"
)

var start string
var destination string
var path_found [][]string
var total_link int
var single_path bool
var max_depth int

type Node struct {
	link string
	depth int
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
	total_link += 1
	defer wg.Done()
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	domain := "https://en.wikipedia.org/wiki/"
	
	x.Lock()
	l := (*queue)[0].link
	current := (*queue)[0].depth + 1
	*queue = (*queue)[1:]
	x.Unlock()

	c.OnError(func (e *colly.Response, err error) {
		log.Println("Terjadi error: ", err)
	})
	
	c.OnHTML("a[href^='/wiki']", func(e *colly.HTMLElement) {
		child := e.Attr("href")[6:]
		if !strings.Contains(child, ":") && !strings.Contains(child, "disambiguation") && child != "Main_Page" {
			// cek apakah link mengandung #
			n := strings.Index(child, "#")
			if n != -1 {
				child = cutLink(child,n)
			}
			
			// cek apakah link ada di array URL atau di map
			x.Lock()
			_, found := m[child]
			if !isInNode(*queue, child) && !found {
				if child == destination {
					max_depth = current
					tempPath := makePath(m, l, child)
					if !isInPath(path_found, tempPath) {
						path_found = append(path_found, makePath(m, l, child))
					}
				} else {
					m[child] = l
					if (current < max_depth) {
						var tempNode Node
						tempNode.link = child
						tempNode.depth = current
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

// func BFS(awal string, akhir string) ([][]string, int, int, int64) {
func main() {
	total_link = 0
	var wait sync.WaitGroup
	var mut sync.Mutex
	var unvisitedQueue []Node
	var minPath int
	single_path = false
	max_depth = 100
	
	// start = awal // nanti ini input start
	// destination = akhir // nanti ini input final
	var startNode Node
	startNode.link = "Nasi_padang" // nanti ini input start
	startNode.depth = 0
	destination = "Indonesia" // nanti ini input final
	unvisitedQueue = append(unvisitedQueue, startNode)
	
	visitedMap := map[string]string {start: "start"}
	
	// hitung waktu proses BFS
	begin := time.Now().UnixNano() / int64(time.Millisecond)
	sekarang := begin
	// Loop berhenti jika queue habis atau (waktu melebihi 4,5 menit dan ada jalur ditemukan) atau waktu melebihi 5 menit
	for (sekarang - begin <= 270000) && len(unvisitedQueue) > 0 {
		// Cari jumlah proses yang dijalankan (maksimal 100)

		n_child := len(unvisitedQueue)
		if n_child > 200 {
			n_child = 200
		}

		// Lakukan proses BFS
		for i := 0; i < n_child; i++ {
			wait.Add(1)
			go validasiLinkBFS(&unvisitedQueue, visitedMap, &mut, &wait)
		}
		wait.Wait()
		if single_path && len(path_found) > 0 {
			path_found = path_found[:1]
			break
		}

		sekarang = time.Now().UnixNano() / int64(time.Millisecond)
	}

	exTime := sekarang - begin
	if (len(path_found) > 0) {
		minPath = len(path_found[0])-1
		fmt.Println("List of paths:")
		fmt.Println(path_found)
		fmt.Printf("Found %d path(s), with minimum depth %d\n", len(path_found), minPath)
	} else {
		minPath = -1
		fmt.Println("No path found")
	}
	fmt.Printf("Execution time: %dm %.2fs\nLink visited: %d\n", exTime/60000, float64(exTime%60000)/1000, total_link)

	// return path, minPath, total_link, exTime 
}