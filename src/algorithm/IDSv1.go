package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

var start string
var destination string
var path_found [][]string // list berisi jalur menuju tujuan
var begin time.Time
var total_link_visited int
var single_path bool

func main() {
	begin = time.Now()
	single_path = false

	fmt.Println("Starting WikiRace!\n")
	var wg sync.WaitGroup

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	start = "nasi_padang"
	destination = "Cooked_rice"

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
		fmt.Println("Runtime : ", end.Sub(begin).Milliseconds(), "ms")
	} else {
		var path_of_url []string // list berisi judul wiki yang dilewatin
		path_of_url = append(path_of_url, start)

		for {
			if len(path_found) > 0 {
				if !single_path && len(path_found) > 0 {
					var new_path_found [][]string
					depth := 100
					for i := 0; i < len(path_found); i++ {
						if len(path_found[i]) < depth {
							depth = len(path_found[i])
						}
					}
					for j := 0; j < len(path_found); j++ {
						if len(path_found[j]) == depth {
							new_path_found = append(new_path_found, path_found[j])
						}
					}
					path_found = new_path_found
				}
				end := time.Now()
				fmt.Println("Path found : ", path_found)
				fmt.Println("Number of links visited : ", total_link_visited)
				fmt.Println("Path length : ", len(path_found[0]))
				fmt.Println("Runtime : ", end.Sub(begin).Milliseconds(), "ms")
				break
			}

			// limit := time.Now()
			// duration := limit.Sub(begin)
			// if ((duration > 2*time.Minute+30*time.Second) && (len(path_found) > 0)) || (duration > 5*time.Minute) {
			// 	end := time.Now()
			// 	fmt.Println("Path found : ", path_found)
			// 	fmt.Println("Number of links visited : ", total_link_visited)
			// 	fmt.Println("Path length : ", len(path_found[0]))
			// 	fmt.Println("Runtime : ", end.Sub(begin))
			// 	break
			// }

			if iterasi == 5 {
				fmt.Println("Sengaja dibatesin 5, coba tanya Vanson")
				break
			}
			fmt.Print("No result found with depth ", iterasi, "! Attempting with depth ", iterasi+1, "...\n\n")

			total_link_visited = 0

			iterasi += 1
			fmt.Println("Depth ", iterasi)

			wg.Add(1)
			go func() {
				dls(c, path_of_url, start, iterasi, wg)
			}()
			wg.Wait()
		}
	}
}

func dls(c *colly.Collector, path_of_url []string, url_scraped string, iterasi int, wg *sync.WaitGroup) {
	defer wg.Done()
	sem := make(chan struct{}, 5)

	var page string = "https://en.wikipedia.org/wiki/"

	list_of_url := make([]string, 0)
	var link_to_visit string = page + url_scraped

	new_path_of_url := make([]string, 0)
	new_path_of_url = append(new_path_of_url, path_of_url...)
	validasiLinkIDS(&link_to_visit, &list_of_url, new_path_of_url)

	iterasi -= 1
	for j := 0; j < len(list_of_url); j++ {
		limit := time.Now()
		duration := limit.Sub(begin)
		if (single_path) && (len(path_found) > 0) {
			break
		}
		if duration > 4*time.Minute+45*time.Second {
			// fmt.Println("LIMIT")
			// fmt.Println("Path found : ", path_found)
			// fmt.Println("Number of links visited : ", total_link_visited)
			// fmt.Println("Path length : ", len(path_found[0]))
			// fmt.Println("Runtime : ", duration)
			// os.Exit(0)
			break
		}

		if list_of_url[j] == destination {
			new_path_of_url := make([]string, 0)
			new_path_of_url = append(new_path_of_url, path_of_url...)
			new_path_of_url = append(new_path_of_url, list_of_url[j])
			path_found = append(path_found, new_path_of_url)

			// print duration
			// end := time.Now()
			// fmt.Println("Path found : ", path_found[0])
			// fmt.Println("Number of links visited : ", total_link_visited)
			// fmt.Println("Path length : ", len(path_found[0]))
			// fmt.Println("Runtime : ", end.Sub(begin))
			// os.Exit(0)
		} else {
			if iterasi > 0 {

				new_path_of_url := make([]string, 0)
				new_path_of_url = append(new_path_of_url, path_of_url...)
				new_path_of_url = append(new_path_of_url, list_of_url[j])
				sem <- struct{}{}
				wg.Add(1)
				go func() {
					defer func() {
						<-sem
					}()
					dls(c, new_path_of_url, list_of_url[j], iterasi, wg)
				}()
			}
		}
	}
}