package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

func DownloadImage(url, outputDir string) error {
    // if outputDir not exist then create
    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return err
		}
	}
	
	
	// Make an HTTP GET request to the image URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to download image. Status code: %d", resp.StatusCode)
	}

	// Get the file name from the URL
	fileName := filepath.Base(url)
	// fmt.Println(fileName)
	// fileName = strings.Replace(fileName, "/", "_", -1)
	// fmt.Println("filename dari -------------", fileName)

	// Create the full path for the output file
	outputPath := filepath.Join(outputDir, fileName)

	// Create a new file to save the image
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body (image data) to the local file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}



func scraping_image(url string) {

    start := time.Now()

    c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()

		// Use Chromedp to navigate to the URL and wait for page load
		err := chromedp.Run(ctx,
			chromedp.Navigate(r.URL.String()),
			chromedp.WaitReady("body"), // Wait until the body is ready
		)
		if err != nil {
			log.Fatal(err)
		}
	})

    // Find and print all links
    c.OnHTML("img", func(e *colly.HTMLElement) {
		// fmt.Println(e)
        links := e.Attr("src")

		if(!strings.Contains(links, "gif") && !strings.Contains(links, "svg")){
			baseUrl := strings.Split(url, "/")[2]
			if(!strings.Contains(links, baseUrl)){
				links = "https://" + baseUrl + links
			}
			fmt.Println( "ini links", links)


			DownloadImage(links, "../../public/images/uploadDataset")

			folder := "../../public/images/uploadDataset/"
			fmt.Println(links)
			file := strings.Split(links, "/")
			filename := file[len(file)-1]
			if(strings.Contains(filename, "?")){
				filename = strings.Split(filename, "?")[0]
			}

			fmt.Println("ini dari sssss ", filename)
			// filename = strings.Replace(filename, "/", "_", -1)

			filename = folder + filename
			fmt.Println(filename)
			// file_hash(filename)

			var imgfeat ImageFeature
			vektor_color := processPicture(filename)
			vektor_texture := processPicture_texture(filename)
			imgfeat.filename = filename
			imgfeat.data_color = vektor_color
			imgfeat.data_texture = vektor_texture

			arrVecDence = append(arrVecDence, imgfeat)
		}

		
    })

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
    c.Visit(url)

    elapsed := time.Since(start)
    fmt.Println("Execution time :", elapsed)
}
// func file_hash(filename string) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	h := md5.New()
// 	if _, err := io.Copy(h, f); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%x\n", h.Sum(nil))
// }
