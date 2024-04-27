package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func IDS(startUrl, destinationUrl string, single_pathTes bool) Result {
	begin = time.Now()
	single_path = single_pathTes
	path_found = nil // Reset path_found for each run
	total_link_visited = 0

	fmt.Println("Starting WikiRace!\n")
	var wg sync.WaitGroup

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	start = startUrl
	destination = destinationUrl

	central(c, &wg)

	duration := time.Since(begin).Milliseconds()
	pathLength := 0
	if len(path_found) > 0 {
		pathLength = len(path_found[0])
	}

	return Result{
		Paths:        path_found,
		TotalLinks:   total_link_visited,
		PathLength:   pathLength,
		DurationInMS: duration,
	}
}

/*
Fungsi central yang mengatur iterasi depth dan pemberhentian
*/
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

/*
Depth Limit Search
Fungsi rekursif buat nyari hasil sesuai dengan batas iterasi yang diberikan oleh central.
Fungsi akan mulai dari depth 0 sampai n
*/
func dls(c *colly.Collector, path_of_url []string, url_scraped string, iterasi int, wg *sync.WaitGroup) {
	defer wg.Done()
	sem := make(chan struct{}, 10) // konkurensi dibatasi 5 setiap pembangkitan. Adjustable

	var page string = "https://en.wikipedia.org/wiki/"

	list_of_url := make([]string, 0)
	var link_to_visit string = page + url_scraped

	// membuat list jalur url yang baru persis dengan parameter. Berpotensi digunakan dalam fungsi scraping
	new_path_of_url := make([]string, 0)
	new_path_of_url = append(new_path_of_url, path_of_url...)
	// memanggil fungsi scraping judul
	validasiLinkIDS(&link_to_visit, &list_of_url, new_path_of_url)

	iterasi -= 1
	for j := 0; j < len(list_of_url); j++ { // validasi dan membangkitkan node anak-anak hasil scraping
		limit := time.Now()
		duration := limit.Sub(begin)
		if (single_path) && (len(path_found) > 0) { // batasan untuk single path, jika sudah ditemukan hasil akan break
			break
		}
		if duration > 4*time.Minute+45*time.Second { // batasan waktu sebelum lima menit, dapat diadjust
			break
		}

		if list_of_url[j] == destination { // jika hasil ditemukan, masukkan pada path_found
			new_path_of_url := make([]string, 0)
			new_path_of_url = append(new_path_of_url, path_of_url...)
			new_path_of_url = append(new_path_of_url, list_of_url[j])
			path_found = append(path_found, new_path_of_url)

		} else {
			if iterasi > 0 { // jika masih dalam jangkauan iterasi limit, membangkitkan node anak

				// membuat list jalur url yang telah ditambahi judul anak
				new_path_of_url := make([]string, 0)
				new_path_of_url = append(new_path_of_url, path_of_url...)
				new_path_of_url = append(new_path_of_url, list_of_url[j])

				sem <- struct{}{} // acquire semaphore
				wg.Add(1)
				go func() {
					defer func() { // release semaphore if function finishes
						<-sem
					}()
					dls(c, new_path_of_url, list_of_url[j], iterasi, wg) // rekursif Depth Limit Search
				}()
			}
		}
	}
}

// func IDS() {
// 	fmt.Println("Starting WikiRace!")

// 	// inisialisasi variable
// 	// HAPUS
// 	single_path = false
// 	// HAPUS
// 	var wg sync.WaitGroup // waitgroup untuk keep track konkurensi

// 	// Create a new collector
// 	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
// 	// HAPUS
// 	start = "Neuroscience"
// 	destination = "Springtail"
// 	// HAPUS

// 	// memanggil fungsi central
// 	central(c, &wg)

// 	fmt.Println("Program finished!")
// }
