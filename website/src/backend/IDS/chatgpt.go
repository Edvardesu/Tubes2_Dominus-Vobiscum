package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go start_page target_page")
		os.Exit(1)
	}

	startPage := os.Args[1]
	targetPage := os.Args[2]

	// Create a new collector
	c := colly.NewCollector()

	// Create a wait group to keep track of scraping goroutines
	var wg sync.WaitGroup

	// Create a channel to communicate between goroutines
	ch := make(chan []string)

	// Start a goroutine to fetch links from each page
	go func() {
		// Start with the start page
		wg.Add(1)
		fetchLinks(c, startPage, ch, &wg)
	}()

	// Start a timer to limit the search time
	timer := time.NewTimer(10 * time.Second)

	// Start a goroutine to process links and search for the target page
	go func() {
		for {
			select {
			case links := <-ch:
				for _, link := range links {
					if link == targetPage {
						fmt.Println("Path found:", startPage, "->", link)
						os.Exit(0)
					}
					// Start a new goroutine to fetch links from the current page
					wg.Add(1)
					go fetchLinks(c, link, ch, &wg)
				}
			case <-timer.C:
				fmt.Println("Timeout: Target page not found within 10 seconds.")
				os.Exit(1)
			}
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Target page not found.")
}

func fetchLinks(c *colly.Collector, page string, ch chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Visit the page
	err := c.Visit("https://en.wikipedia.org/wiki/" + page)
	if err != nil {
		log.Printf("Error visiting page %s: %v\n", page, err)
		return
	}

	// Extract links from the page
	links := make([]string, 0)
	c.OnHTML("a[href^='/wiki/']", func(e *colly.HTMLElement) {
		link := strings.TrimPrefix(e.Attr("href"), "/wiki/")
		links = append(links, link)
	})

	// Send links to the channel
	ch <- links
}
