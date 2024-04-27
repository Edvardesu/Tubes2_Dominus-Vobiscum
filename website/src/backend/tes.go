// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	var sliceURL []string
// 	scrapeUrl := "https://en.wikipedia.org/wiki/Nasi_padang"
// 	// Instantiate a new collector
// 	// c := colly.NewCollector()
// 	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org/wiki"))

// 	// Visit the website and print its title
// 	c.OnHTML("p a[href]", func(e *colly.HTMLElement) {
// 		if strings.HasPrefix(e.Attr("href"), "/wiki/") {
// 			sliceURL = append(sliceURL, e.Attr("href"))
// 			fmt.Println(e.Attr("href"))
// 		}
// 	})

// 	// Start scraping
// 	c.Visit(scrapeUrl)

// 	for i := 0; i < len(sliceURL); i++ {
// 		fmt.Println(sliceURL[i])
// 	}
// 	fmt.Print("done")
// }
