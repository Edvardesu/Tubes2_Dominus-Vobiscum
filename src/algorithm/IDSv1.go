package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

/*
variable global agar mudah diakses di dalam fungsi
*/
var start string           // judul awal Wiki Race
var destination string     // judul tujuan Wiki Race
var path_found [][]string  // list berisi jalur menuju tujuan
var begin time.Time        // waktu mulai
var total_link_visited int // total link yang dikunjungi
var single_path bool       // true jika pemilihan IDS single_path

func main() {
	fmt.Println("Starting WikiRace!\n")

	// inisialisasi variable
	single_path = false
	var wg sync.WaitGroup // waitgroup untuk keep track konkurensi

	// Create a new collector
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.CacheDir("./Cache"))
	start = "Neuroscience"
	destination = "Springtail"

	// memanggil fungsi central
	central(c, &wg)

	fmt.Println("Program finished!")
}

/*
Fungsi central yang mengatur iterasi depth dan pemberhentian
*/
func central(c *colly.Collector, wg *sync.WaitGroup) {
	begin = time.Now() // catet waktu mulai
	var iterasi int = 0

	fmt.Println("Depth ", iterasi)
	if start == destination { // check apabila judul awal = judul tujuan
		end := time.Now()
		fmt.Println("Path found : [", start, "]")
		fmt.Println("Number of links visited : 0")
		fmt.Println("Path length : 1")
		fmt.Println("Runtime : ", end.Sub(begin).Milliseconds(), "ms")
	} else {
		var path_of_url []string // list berisi judul wiki yang dilewatin
		path_of_url = append(path_of_url, start)

		// iterasi sampai di break dari dalam
		for {
			if len(path_found) > 0 { // break jika sudah ditemukan solusi
				if !single_path && len(path_found) > 0 { // jika multiple path, path_found akan dirapikan (membuang path2 dengan depth+1)
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

			// ini failsafe biar ketahan aja iterasinya, biasanya kalo udh sampe iterasi 5 berarti somting wong
			if iterasi == 5 {
				fmt.Println("Kemungkinan error!")
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
			wg.Wait() // wait sampai seluruh konkurensi selesai
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
	sem := make(chan struct{}, 5) // konkurensi dibatasi 5 setiap pembangkitan. Adjustable

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
