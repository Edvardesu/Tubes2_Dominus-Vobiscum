package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func isIn(l []string, s string) bool {
	for i := 0; i < len(l); i++ {
		if s == l[i] {
			return true
		}
	}
	return false
}

func checkHastag(l string) int {
	ctr := -1
	for i := 0; i < len(l); i++ {
		if l[i] == '#' {
			ctr = i
		}
	}
	return ctr
}

func cutLink(l string, i int) string {
	return l[:i]
}

func validasiLinkIDS(l *string, c *colly.Collector, slice *[]string) {
	c.OnHTML("p a[href^='/wiki']", func(e *colly.HTMLElement) {
		link := e.Attr("href")[6:]
		if !strings.Contains(link, "File:") {
			// cek apakah link mengandung #
			n := checkHastag(link)
			if n != -1 {
				link = cutLink(link, n)
			}

			// cek apakah link ada di array URL
			if !isIn(*slice, link) {
				*slice = append(*slice, link)
				fmt.Println("Link found :", link)
			}
		}
	})

	// Start scraping
	c.Visit(*l)
}

// func main() {
// 	var sliceURL []string
// 	scrapeUrl := "https://en.wikipedia.org/wiki/Frikadelle"
// 	// Instantiate a new collector
// 	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

// 	c.OnRequest(func(e *colly.Request) {
// 		fmt.Println("Mengunjungi ", e.URL)
// 	})

// 	c.OnError(func(e *colly.Response, err error) {
// 		log.Println("Terjadi error: ", err)
// 	})

// 	// Visit the website and print its title
// 	validasiLink(&scrapeUrl, c, &sliceURL)

// 	for i := 0; i < len(sliceURL); i++ {
// 		fmt.Println(sliceURL[i])
// 	}
// }
