// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	// Set a lower memory limit for multipart forms (default is 32 MiB)
// 	router.MaxMultipartMemory = 8 << 20 // 8 MiB

// 	// Specify the destination folder to save the uploaded files
// 	dst := "uploadDataset"

// 	var uploadedFilePaths []string

// 	router.POST("/uploadDataset", func(c *gin.Context) {
// 		// Multipart form
// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
// 			return
// 		}

// 		files := form.File["files"]

// 		for _, file := range files {
// 			log.Println(file.Filename)

// 			// Create the destination path for the file
// 			destPath := filepath.Join(dst, file.Filename)

// 			// Upload the file to the specified destination
// 			if err := c.SaveUploadedFile(file, destPath); err != nil {
// 				c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %s", err.Error()))
// 				return
// 			}

// 			// Save the uploaded file path
// 			uploadedFilePaths = append(uploadedFilePaths, destPath)
// 		}

// 		// Call processPictures with the paths of all uploaded images
// 		vektor2 := processPictures(uploadedFilePaths)
// 		fmt.Printf("Vektor: %f\n", vektor2)

// 		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
// 	})

// 	// Run the server on port 8080
// 	router.Run(":8080")
// }

// // processPictures takes a slice of file paths and processes each image
// func processPictures(filePaths []string) float64 {
// 	// Your logic to process the images and calculate the vektor
// 	// For simplicity, let's assume you have a function named processPicture
// 	// that takes a file path and returns a float64
// 	// You may need to adjust this based on your actual implementation
// 	var totalVektor float64
// 	for _, filePath := range filePaths {
// 		vektor := processPict(filePath)
// 		// Accumulate the vektors
// 		totalVektor += vektor
// 	}
// 	// Calculate the average or use any other logic based on your requirement
// 	averageVektor := totalVektor / float64(len(filePaths))
// 	return averageVektor
// }

// // processPicture takes a single file path and processes the image
// func processPict(filePath string) float64 {
// 	// Your logic to process a single image and return a float64 vektor
// 	// Placeholder implementation, replace with your actual logic
// 	return float64(len(filePath))
// }

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	// "os"
	"encoding/json"
	"path/filepath"
	"strings"

	// "gonum.org/v1/gonum/mat"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gonum.org/v1/gonum/mat"
	// "gonum.org/v1/gonum/mat"
)

type Employee struct {
	Vec string
}

type ImageFeature struct {
	filename     string
	data_color   [4][4]*mat.VecDense
	data_texture []float64
}

var arrVecDence []ImageFeature

func checkColorSimilarity() ([]string, []float64, time.Duration) {
	start := time.Now()
	var filename []string
	var result []float64
	var length = len(arrVecDence)
	for y := 0; y < length; y++ {
		temp := cosine_similarity(queryImage.data_color, arrVecDence[y].data_color)
		fname := arrVecDence[y].filename
		fmt.Println("filename ", fname)
		fmt.Println("simmilarity: ", temp)
		if temp > 0.6 {
			fname = strings.Split(fname, "/")[len(strings.Split(fname, "/"))-1]
			filename = append(filename, fname)
			result = append(result, temp)
		}
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	return filename, result, elapsed
}

func checkTextureSimmilarity() ([]string, []float64, time.Duration) {
	start := time.Now()
	var filename []string
	var result []float64
	var length = len(arrVecDence)
	for y := 0; y < length; y++ {
		temp := cosine_similarity_texture(queryImage.data_texture, arrVecDence[y].data_texture)
		fname := arrVecDence[y].filename
		fmt.Println("filename ", fname)
		fmt.Println("simmilarity: ", temp)
		if temp > 0.6 {
			fname = strings.Split(fname, "/")[len(strings.Split(fname, "/"))-1]
			filename = append(filename, fname)
			result = append(result, temp)
		}
	}
	elapsed := time.Since(start)
	return filename, result, elapsed
}

type JsonRequest struct {
	URL string `json:"url"`
}

func runGin() {
	router := gin.Default()

	// Enable CORS middleware
    router.Use(cors.Default())

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Specify the destination folder to save the uploaded files
	dst := "../../public/images/uploadDataset"

	router.POST("/uploadDataset", func(c *gin.Context) {
		// clear array everytime dataset uploaded
		arrVecDence = nil

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
			return
		}

		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			replaced_string := strings.Replace(file.Filename, " ", "-", -1)

			// Create the destination path for the file
			destPath := filepath.Join(dst, replaced_string)
			// fmt.Printf("destPath: %s\n", destPath)
			bakulGorengan := "../../public/images/uploadDataset/" + replaced_string
			fmt.Printf("bakulGorengan: %s\n", bakulGorengan)

			// Upload the file to the specified destination
			if err := c.SaveUploadedFile(file, destPath); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving file: %s", err.Error()))
				return
			}

			var imgfeat ImageFeature

			vektor2 := processPicture(bakulGorengan) // ./Folder/namfile
			// fmt.Printf("Vektor: %f\n", vektor2)

			imgfeat.filename = bakulGorengan
			imgfeat.data_color = vektor2

			vektor2_texture := processPicture_texture(bakulGorengan)
			imgfeat.data_texture = vektor2_texture

			arrVecDence = append(arrVecDence, imgfeat)

			a := Employee{Vec: bakulGorengan}
			// o := Employee{Vec:"Orange"}

			var fs []Employee
			fs = append(fs, a)
			// fs = append(fs, o)
			log.Println(fs)

			file, _ := json.MarshalIndent(fs, "", " ")
			_ = ioutil.WriteFile("test.json", file, 0644)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

		// checking all images

	})

	router.GET("/colorResult", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		file, result, elapsed := checkColorSimilarity()
		c.JSON(200, gin.H{
			"data":   file,
			"result": result,
			"time":   elapsed,
		})
	})

	router.GET("/textureResult", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		file, result, elapsed := checkTextureSimmilarity()
		c.JSON(200, gin.H{
			"data":   file,
			"result": result,
			"time":   elapsed,
		})
	})

	router.POST("/scrapping", func(c *gin.Context) {
		var requestData JsonRequest

		c.BindJSON(&requestData)

		url := requestData.URL

		c.Header("Access-Control-Allow-Origin", "*")
		fmt.Println("ini dari scrapping")
		fmt.Println(url)
		scraping_image(url)

		c.JSON(200, gin.H{
			// "data": file,
			// "result" : result,
			"name": "edbert",
		})
	})

	// Run the server on port 8080
	router.Run(":8081")
}
