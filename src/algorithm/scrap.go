package main

import (
	"strings"
	"github.com/gocolly/colly/v2"
	"slices"
)


func isIn(lis[] string, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i]) {return true}
	}
	return false
}

func makePath(m map[string]string, f string) []string {
	path := []string{f}
	parent := ""
	for parent != "start" {
		parent = m[f]
		path = append(path, parent)
		f = parent
	}
	slices.Reverse(path)
	return path[1:]
}

func checkHastag(l string) int {
	for i := 0; i < len(l); i++ {
		if l[i] == '#' {
			return i
		}
	}
	return -1
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
				link = cutLink(link,n)
			}

			// cek apakah link ada di array URL
			if !isIn(*slice, link) { 
				*slice = append(*slice, link)
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

// 	c.OnRequest(func (e *colly.Request) {
// 		fmt.Println("Mengunjungi ", e.URL)
// 	})

// 	c.OnError(func (e *colly.Response, err error) {
// 		log.Println("Terjadi error: ", err)
// 	})

// 	// Visit the website and print its title
// 	validasiLink(&scrapeUrl, c, &sliceURL)

// 	for i:=0; i < len(sliceURL); i++ {
// 		fmt.Println(sliceURL[i])
// 	}
// }