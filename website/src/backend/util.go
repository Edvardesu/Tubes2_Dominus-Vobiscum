package main

import (
	"strings"
	"reflect"
	"slices"
	"log"
	"fmt"
	"sync"
	"time"
	"github.com/gocolly/colly/v2"
)

// Utility umum
/*
variable global agar mudah diakses di dalam fungsi
*/
var start string           // judul awal Wiki Race
var destination string     // judul tujuan Wiki Race
var path_found [][]string  // list berisi jalur menuju tujuan
var begin time.Time        // waktu mulai
var exTime int64	   // Waktu eksekusi program
var total_link_visited int // total link yang dikunjungi
var single_path bool       // true jika pemilihan IDS/BFS single_path

type Result struct {
    Paths            [][]string `json:"paths"`
    TotalLinks       int        `json:"total_links"`
    PathLength       int        `json:"path_length"`
    DurationInMS     int64      `json:"duration_in_ms"`
	SinglePath		bool		`json:"single_path"`
	PathAmount       int        `json:"path_amount"`
}

func isIn(lis []string, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i]) {return true}
	}
	return false
}

func cutLink(l string, i int) string {
	return l[:i]
}

// BFS
type Node struct {
	link  string
	depth int
}

func isInNode(lis []Node, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i].link) {return true}
	}
	return false
}

func isInPath(mat [][]string, lis []string) bool {
	for i := len(mat) - 1; i >= 0; i-- {
		if reflect.DeepEqual(mat[i], lis) {
			return true
		}
	}
	return false
}

func makePath(m map[string]string, now string, des string) []string {
	path := []string{now}
	parent := ""
	for parent != "start" {
		parent = m[now]
		path = append(path, parent)
		now = parent
	}
	slices.Reverse(path)
	path = append(path, des)
	return path[1:]
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

// IDS
func validasiLinkIDS(l *string, slice *[]string, new_path_of_url []string) {
	total_link_visited += 1
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))

	// checking title
	var flag bool = false // ini true jika destinasi nya adalah string yang divisit
	var title string
	c.OnHTML(".firstHeading", func(e *colly.HTMLElement) {
		title = e.Text
		title = strings.ReplaceAll(e.Text, " ", "_")
		if title == destination {
			new_path_of_url[len(new_path_of_url)-1] = title
			path_found = append(path_found, new_path_of_url)
			flag = true
		}
	})

	// scraping
	if !flag {
		scraping(c, slice)
	}
	// Start scraping
	c.Visit(*l)
}

func scraping(c *colly.Collector, slice *[]string) {
	c.OnHTML("a[href^='/wiki']", func(e *colly.HTMLElement) {
		link := e.Attr("href")[6:]
		if (!strings.Contains(link, ":")) && (!strings.Contains(link, "disambiguation")) && (!strings.Contains(link, "Main_Page")) {
			// cek apakah link mengandung #
			n := strings.Index(link, "#")
			if n != -1 {
				link = cutLink(link, n)
			}

			// cek apakah link ada di array URL
			if !isIn(*slice, link) {
				*slice = append(*slice, link)
			}
		}
	})
}