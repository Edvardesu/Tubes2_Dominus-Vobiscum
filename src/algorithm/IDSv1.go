package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

var start string
var destination string
var path_found [][]string // list berisi jalur menuju tujuan
var begin time.Time
var total_link_visited int

func main() {
	begin = time.Now()

	fmt.Println("Starting WikiRace!\n")
	var wg sync.WaitGroup

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	start = "Joko_Widodo"
	destination = "Cadmium"

	central(c, &wg)
	fmt.Println("Program finished!")
}

func central(c *colly.Collector, wg *sync.WaitGroup) {
	begin = time.Now()
	var iterasi int = 0

	fmt.Println("Depth ", iterasi)
	if start == destination {
		end := time.Now()
		fmt.Println("Path found : [", start, "]")
		fmt.Println("Number of links visited : 0")
		fmt.Println("Path length : 1")
		fmt.Println("Runtime : ", end.Sub(begin))
	} else {
		var path_of_url []string // list berisi judul wiki yang dilewatin
		path_of_url = append(path_of_url, start)

		for {

			if len(path_found) > 0 {
				end := time.Now()
				fmt.Println("Path found : ", path_found)
				fmt.Println("Number of links visited : ", total_link_visited)
				fmt.Println("Path length : ", len(path_found[0]))
				fmt.Println("Runtime : ", end.Sub(begin))
				break
			}
			if iterasi == 5 {
				fmt.Println("Sengaja dibatesin 5, coba tanya Vanson")
				break
			}
			fmt.Print("No result found with depth ", iterasi, "! Attempting with depth ", iterasi+1, "...\n\n")

			total_link_visited = 0

			iterasi += 1
			fmt.Println("Depth ", iterasi)

			wg.Add(1)
			dls(c, path_of_url, start, iterasi, wg)
			wg.Wait()
		}
	}
}

func dls(c *colly.Collector, path_of_url []string, url_scraped string, iterasi int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println("IN")

	var page string = "https://en.wikipedia.org/wiki/"

	// var list_of_url []string
	list_of_url := make([]string, 0)
	var link_to_visit string = page + url_scraped

	new_path_of_url := make([]string, 0)
	new_path_of_url = append(new_path_of_url, path_of_url...)
	validasiLinkIDS(&link_to_visit, &list_of_url, new_path_of_url)
	// fmt.Println(list_of_url)
	// fmt.Println("OUT")
	iterasi -= 1
	for j := 0; j < len(list_of_url); j++ {

		if list_of_url[j] == destination {
			new_path_of_url := make([]string, 0)
			new_path_of_url = append(new_path_of_url, path_of_url...)
			new_path_of_url = append(new_path_of_url, list_of_url[j])
			path_found = append(path_found, new_path_of_url)

			// print duration
			end := time.Now()
			fmt.Println("Path found : ", path_found[0])
			fmt.Println("Number of links visited : ", total_link_visited)
			fmt.Println("Path length : ", len(path_found[0]))
			fmt.Println("Runtime : ", end.Sub(begin))
			os.Exit(0)
		}
		if iterasi > 0 {
			new_path_of_url := make([]string, 0)
			new_path_of_url = append(new_path_of_url, path_of_url...)
			new_path_of_url = append(new_path_of_url, list_of_url[j])
			wg.Add(1)
			go func() {
				dls(c, new_path_of_url, list_of_url[j], iterasi, wg)
			}()
		}
	}
	// for j := 0; j < len(list_of_url); j++ {
	// 	title := getValidLink(list_of_url[j])
	// 	if title == destination {
	// 		new_path_of_url := make([]string, 0)
	// 		new_path_of_url = append(new_path_of_url, path_of_url...)
	// 		new_path_of_url = append(new_path_of_url, list_of_url[j])
	// 		path_found = append(path_found, new_path_of_url)

	// 		// print duration
	// 		end := time.Now()
	// 		fmt.Println("Path found : ", path_found[0])
	// 		fmt.Println("Number of links visited : ", total_link_visited)
	// 		fmt.Println("Path length : ", len(path_found[0]))
	// 		fmt.Println("Runtime : ", end.Sub(begin))
	// 		os.Exit(0)
	// 	}
	// }
}
