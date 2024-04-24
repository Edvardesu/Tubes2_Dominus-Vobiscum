package main

import (
	"fmt"
	"time"
	"strings"
	"log"
	"github.com/gocolly/colly/v2"
)

type Node struct {
	link string
	depth int
}

func isInNode(lis []Node, s string) bool {
	for i := 0; i < len(lis); i++ {
		if (s == lis[i].link) {return true}
	}
	return false
}

func validasiLinkBFS(queue *[]Node, m map[string]string, c *colly.Collector) {
	domain := "https://en.wikipedia.org/wiki/"
	l := (*queue)[0].link
	current := (*queue)[0].depth + 1
	// pointer := &current
	*queue = (*queue)[1:]
	
	c.OnError(func (e *colly.Response, err error) {
		log.Println("Terjadi error: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL) 
	})
	
	c.OnHTML("p a[href^='/wiki']", func(e *colly.HTMLElement) {
		child := e.Attr("href")[6:]
		if !strings.Contains(child, "File:") {
			// cek apakah link mengandung #
			n := checkHastag(child)
			if n != -1 {
				child = cutLink(child,n)
			}
			
			// cek apakah link ada di array URL atau di map
			_, found := m[child]
			if !isInNode(*queue, child) && !found {
				log.Printf("Found link %s from %s current depth is %d\n", child, l, current)
				m[child] = l
				if (current < 2) {
					var tempNode Node
					tempNode.link = child
					tempNode.depth = current
					*queue = append(*queue, tempNode)
				}
			}
		}
	})
	
	// Start scraping
	c.Visit(domain + l)
}

func main() {
	var unvisitedQueue []Node
	var listOfPaths [][]string
	var start Node
	// scrapeUrl := "Frikadelle" // nanti ini input start
	final := "Yolk" // nanti ini input final

	start.link = "Frikadelle"
	start.depth = 0
	unvisitedQueue = append(unvisitedQueue, start)

	visitedMap := map[string]string {start.link: "start"}

	// Instantiate a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))

	// Visit the website and print its title
	for len(unvisitedQueue) > 0 {
		validasiLinkBFS(&unvisitedQueue, visitedMap, c)
		_, finalFound := visitedMap[final]
		if finalFound {
			listOfPaths = append(listOfPaths, makePath(visitedMap, final))
			delete(visitedMap, final)
		}
		fmt.Println(unvisitedQueue)
		fmt.Println(visitedMap)
		time.Sleep(2 * time.Second)
	}

	fmt.Println(listOfPaths)
}