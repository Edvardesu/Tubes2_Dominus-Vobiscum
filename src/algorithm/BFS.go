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

type Node struct {
	link string
	depth int
}

var start Node
var destination string
var path [][]string
var total_link int

func isInNode(lis []Node, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i].link) {return true}
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
	l := (*queue)[0].link
	current := (*queue)[0].depth + 1
	*queue = (*queue)[1:]
	x.Unlock()

	c.OnError(func (e *colly.Response, err error) {
		log.Println("Terjadi error: ", err)
	})
	
	c.OnHTML("a[href^='/wiki']", func(e *colly.HTMLElement) {
		child := e.Attr("href")[6:]
		if !(strings.Contains(child, "File:") || strings.Contains(child, "Category:") || strings.Contains(child, "Help:") || strings.Contains(child, "Wikipedia:")) {
			total_link += 1
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
					tempPath := makePath(m, l, child)
					if !isInPath(path, tempPath) {
						path = append(path, makePath(m, l, child))
					}
				} else {
					m[child] = l
					// fixed depth 4
					if (current < 4) {
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

func BFS(awal string, akhir string) ([][]string, int, int, int64) {
// func main() {
	total_link = 0
	var wait sync.WaitGroup
	var mut sync.Mutex
	var unvisitedQueue []Node
	var minPath int
	
	start.link = "Joko_Widodo" // nanti ini input start
	start.depth = 0
	destination = "Cadmium" // nanti ini input final
	unvisitedQueue = append(unvisitedQueue, start)
	
	visitedMap := map[string]string {start.link: "start"}
	
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
		fmt.Printf("Found %d path(s), with minimum depth %d\n", len(path), minPath)
	} else {
		fmt.Println("No path found")
	}
	fmt.Printf("Execution time: %dm %.2fs\nLink visited: %d\n", exTime/60000, float64(exTime%60000)/1000, total_link)

	return path, minPath, total_link, exTime 
}