package main

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly/v2"
)

var start string
var destination string
var path_found [][]string // list berisi jalur menuju tujuan

func main() {
	fmt.Println("Starting WikiRace!")

	var wg sync.WaitGroup

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org/wiki"))
	start = "https://en.wikipedia.org/wiki/nasi_padang"
	destination = "Indonesia"

	var path_of_url []string // list berisi judul wiki yang dilewatin
	var url_scraped string   // string berisi judul yang mau discrape
	var iterasi int = 2

	path_of_url = append(path_of_url, "nasi_padang")
	url_scraped = "nasi_padang"

	wg.Add(1)
	scraping(c, path_of_url, url_scraped, iterasi, &wg)
	wg.Wait()

	fmt.Print(path_found)

	// Start scraping on http://example.com
	// var page string = "https://en.wikipedia.org/wiki/"
	// var link string
	// fmt.Scan(&link)
	// c.Visit(page)
}

func scraping(c *colly.Collector, path_of_url []string, url_scraped string, iterasi int, wg *sync.WaitGroup) {
	defer wg.Done()
	// On every a element which has href attribute call callback
	var page string = "en.wikipedia.org/wiki"
	fmt.Println("In")

	go func() {
		var list_of_url []string
		c.OnHTML("p a[href^='/wiki/']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			list_of_url = append(list_of_url, link[5:])
			if iterasi == 2 {
				fmt.Println("Link found:", link)
			}
		})

		c.Visit(page + url_scraped)

		iterasi -= 1
		for j := 0; j < len(list_of_url); j++ {
			var new_path_of_url []string = path_of_url
			new_path_of_url = append(new_path_of_url, list_of_url[j])

			if list_of_url[j][5:] == destination {
				path_found = append(path_found, new_path_of_url)
			} else {
				if iterasi > 0 {
					wg.Add(1)
					scraping(c, new_path_of_url, list_of_url[j], iterasi, wg)
				}
			}
		}
	}()
}
