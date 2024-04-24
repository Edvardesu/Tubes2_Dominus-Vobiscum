package main

import (
	"fmt"
	"os"
	"strings"
	"time"

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

func addUnderScore(title string) string {
	var result string
	for i := 0; i < len(title); i++ {
		if title[i] == ' ' {
			result += "_"
		} else {
			result += string(title[i])
		}
	}
	return result
}

func getValidLink(link string) string {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))
	var title string
	c.OnHTML(".firstHeading", func(e *colly.HTMLElement) {
		// Extract the title of the Wikipedia page
		title = strings.ToLower(e.Text)
		// title = addUnderScore(title)
		title = strings.ReplaceAll(title, " ", "_")
		if title == destination {

		}
	})
	c.Visit("https://en.wikipedia.org/wiki/" + link)
	return title
}
func validasiLinkIDS(l *string, slice *[]string, new_path_of_url []string) {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))

	// checking title
	var title string
	c.OnHTML(".firstHeading", func(e *colly.HTMLElement) {
		// Extract the title of the Wikipedia page
		// title = strings.ToLower(e.Text)
		// title = addUnderScore(title)
		title = strings.ReplaceAll(e.Text, " ", "_")
		if title == destination {
			new_path_of_url = append(new_path_of_url, title)
			path_found = append(path_found, new_path_of_url)

			end := time.Now()
			fmt.Println("Path found : ", path_found[0])
			fmt.Println("Number of links visited : ", total_link_visited)
			fmt.Println("Path length : ", len(path_found[0]))
			fmt.Println("Runtime : ", end.Sub(begin))
			os.Exit(0)
		}
	})
	// fmt.Println(*l)
	// scraping
	c.OnHTML("p a[href^='/wiki']", func(e *colly.HTMLElement) {
		total_link_visited += 1
		link := e.Attr("href")[6:]
		if !strings.Contains(link, "File:") {
			// cek apakah link mengandung #
			n := checkHastag(link)
			if n != -1 {
				link = cutLink(link, n)
			}

			// cek apakah link ada di array URL
			if !isIn(*slice, link) {
				// title := getValidLink(link)
				// fmt.Println("Title:", title)

				*slice = append(*slice, link)
				// fmt.Println("Link found :", link)
			}
		}
	})

	// Start scraping
	c.Visit(*l)
}
