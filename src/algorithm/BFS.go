package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

var start string
var destination string
var path_found [][]string
var begin time.Time
var total_link_visited int
var single_path bool
var max_depth int

// Node untuk mencatat link dan kedalamannya
type Node struct {
	link  string
	depth int
}

func isInPath(mat [][]string, lis []string) bool {
	for i := len(mat) - 1; i >= 0; i-- {
		if reflect.DeepEqual(mat[i], lis) {
			return true
		}
	}
	return false
}

func validasiLinkBFS(queue *[]Node, m map[string]string, x *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	// Buat sebuah kolektor baru
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	var flag bool = false 
	var title string
	total_link_visited += 1

	// Ambil head dari queue
	x.Lock()
	l := (*queue)[0].link
	current := (*queue)[0].depth + 1
	*queue = (*queue)[1:]
	x.Unlock()

	c.OnError(func(e *colly.Response, err error) {
		log.Println("Terjadi error: ", err)
	})

	// Kasus jika judul artikel adalah judul destinasi akhir
	x.Lock()
	c.OnHTML(".firstHeading", func(e *colly.HTMLElement) {
		title = e.Text
		title = strings.ReplaceAll(e.Text, " ", "_")
		if title == destination {
			fmt.Println("Ketemu title", title)
			max_depth = current - 1
			tempPath := makePath(m, m[l], title)
			fmt.Println("Path ", tempPath)
			path_found = append(path_found, tempPath)
			fmt.Println("Path Found ", path_found)
			flag = true
		}
	})
	x.Unlock()

	// Scrap child link dari link saat ini jika judulnya tidak sama dengan destinasi
	if !flag {
		c.OnHTML("a[href^='/wiki']", func(e *colly.HTMLElement) {
			child := e.Attr("href")[6:]
			if !strings.Contains(child, ":") && !strings.Contains(child, "disambiguation") && child != "Main_Page" {
				// Cek apakah link mengandung #
				n := strings.Index(child, "#")
				if n != -1 {
					child = cutLink(child, n)
				}
	
				// Cek apakah link ada di array URL atau di map
				x.Lock()
				_, found := m[child]
				if !isInNode(*queue, child) && !found {
					if child == destination {
						// Batasi kedalaman maksimal
						if max_depth == 100 {
							max_depth = current
						}
						tempPath := makePath(m, l, child)
						// Batasi agar goroutine tidak memasukkan path yang sama berkali-kali
						if !isInPath(path_found, tempPath) {
							path_found = append(path_found, makePath(m, l, child))
						}
					} else {
						// Masukkan ke dalam map dan ke dalam queue
						m[child] = l
						var tempNode Node
						tempNode.link = child
						tempNode.depth = current
						*queue = append(*queue, tempNode)
					}
				}
				x.Unlock()
			}
		})
	}

	// Start scraping
	c.Visit("https://en.wikipedia.org/wiki/" + l)
}

// func BFS(awal string, akhir string) ([][]string, int, int, int64) {
func main() {
	total_link_visited = 0
	var wait sync.WaitGroup
	var mut sync.Mutex
	var unvisitedQueue []Node // Queue untuk menyimpan link yang belum dikunjungi
	var minPath int // Jalur minimum dari start ke tujuan
	// HAPUS
	single_path = false
	// HAPUS
	max_depth = 100

	// start = awal // nanti ini input start
	// destination = akhir // nanti ini input final
	var startNode Node
	// HAPUS
	startNode.link = "Neuroscience" // nanti ini input start
	startNode.depth = 0
	destination = "Springtail" // nanti ini input final
	// HAPUS
	unvisitedQueue = append(unvisitedQueue, startNode)

	// Jadikan link pertama sebagai start
	visitedMap := map[string]string{startNode.link: "start"}

	// hitung waktu proses BFS
	begin = time.Now()
	sekarang := begin
	// Loop berhenti jika queue habis atau waktu melebihi 4,5 menit atau kedalaman dengan rute terpendek sudah dicek semua
	for (sekarang.Sub(begin) <= 4*time.Minute+30*time.Second) && len(unvisitedQueue) > 0 && unvisitedQueue[0].depth < max_depth {
		// Cari jumlah proses yang dijalankan (maksimal 200)
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

		// Bereskan jika hanya mencari 1 path
		if single_path && len(path_found) > 0 {
			path_found = path_found[:1]
			break
		}

		sekarang = time.Now()
	}

	exTime := sekarang.Sub(begin).Milliseconds()
	if len(path_found) > 0 {
		minPath = len(path_found[0])
		for i := 0; i < len(path_found); i++ {
			if len(path_found[i]) < minPath {
				minPath = len(path_found[i])
			}
		}

		var sementara [][]string
		for i := 0; i < len(path_found); i++ {
			if len(path_found[i]) == minPath {
				sementara = append(sementara, path_found[i])
			}
		}
		path_found = sementara

		fmt.Println("List of paths:")
		fmt.Println(path_found)
		fmt.Printf("Found %d path(s), with minimum depth %d\n", len(path_found), minPath-1)
	} else {
		minPath = -1
		fmt.Println("No path found")
	}
	fmt.Println("Exec time:", exTime, "ms")
	fmt.Printf("Link visited: %d\n", total_link_visited)

	// return path, minPath, total_link, exTime
}
