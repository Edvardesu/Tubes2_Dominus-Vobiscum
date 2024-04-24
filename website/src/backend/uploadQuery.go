package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"net/http"
)

const PORT = "8080"

var queryImage ImageFeature

func Ping(w http.ResponseWriter, r *http.Request) {
	answer := map[string]interface{}{
		"messageType": "S",
		"message":     "",
		"data":        "PONG",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(answer)
}

// handler to handle the image upload
func UploadImages(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	files := r.MultipartForm.File["file"]

	var errNew string
	var httpStatus int

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		fileType := http.DetectContentType(buff)
		if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/jpg" {
			errNew = "The provided file format is not allowed. Please upload a JPEG, JPG, or PNG image"
			httpStatus = http.StatusBadRequest
			break
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		err = os.MkdirAll("../../public/images/uploads", os.ModePerm)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		f, err := os.Create(fmt.Sprintf("../../public/images/uploads%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		// Call processPicture with the uploaded image's path
		vektor1 := processPicture(f.Name())
		vektor2 := processPicture_texture(f.Name())
		queryImage.data_color = vektor1
		queryImage.data_texture = vektor2
		queryImage.filename = f.Name()
		fmt.Printf("Vektor1: %f\n", vektor1)
	}

	message := "file uploaded successfully"
	messageType := "S"

	if errNew != "" {
		message = errNew
		messageType = "E"
	}

	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}

	resp := map[string]interface{}{
		"messageType": messageType,
		"message":     message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(resp)
}

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/ping", Ping).Methods("GET")
// 	r.HandleFunc("/upload", UploadImages).Methods("POST")

// 	log.Printf("Server is running on http://localhost:%s", PORT)
// 	log.Println(http.ListenAndServe(":"+PORT, r))
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	// "fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
//     // "algos/color.go"
// )

// const PORT = "8080"

// func Ping(w http.ResponseWriter, r *http.Request) {
//     answer := map[string]interface{}{
//         "messageType": "S",
//         "message":     "",
//         "data":        "PONG",
//     }
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(200)
//     json.NewEncoder(w).Encode(answer)

// }

// // handler to handle the image upload
// func UploadImages(w http.ResponseWriter, r *http.Request) {
//     // 32 MB is the default used by FormFile() function
//     if err := r.ParseMultipartForm(32 << 20); err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     w.Header().Set("Access-Control-Allow-Origin", "*")

//     // Get a reference to the fileHeaders.
//     // They are accessible only after ParseMultipartForm is called
//     files := r.MultipartForm.File["file"]

//     var errNew string
//     var http_status int

//     for _, fileHeader := range files {
//         // Open the file
//         file, err := fileHeader.Open()
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusInternalServerError
//             break
//         }

//         defer file.Close()

//         buff := make([]byte, 512)
//         _, err = file.Read(buff)
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusInternalServerError
//             break
//         }

//         // checking the content type
//         // so we don't allow files other than images
//         filetype := http.DetectContentType(buff)
//         if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
//             errNew = "The provided file format is not allowed. Please upload a JPEG,JPG or PNG image"
//             http_status = http.StatusBadRequest
//             break
//         }

//         _, err = file.Seek(0, io.SeekStart)
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusInternalServerError
//             break
//         }

//         err = os.MkdirAll("./uploads", os.ModePerm)
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusInternalServerError
//             break
//         }

//         f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusBadRequest
//             break
//         }

//         defer f.Close()

//         _, err = io.Copy(f, file)
//         if err != nil {
//             errNew = err.Error()
//             http_status = http.StatusBadRequest
//             break
//         }
//     }
//     message := "file uploaded successfully"
//     messageType := "S"

//     if errNew != "" {
//         message = errNew
//         messageType = "E"
//     }

//     if http_status == 0 {
//         http_status = http.StatusOK
//     }

//     resp := map[string]interface{}{
//         "messageType": messageType,
//         "message":     message,
//     }
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http_status)
//     json.NewEncoder(w).Encode(resp)

// }

// func main() {
// 	r := mux.NewRouter()
//     // r.HandleFunc("/ping", nil).Methods("GET")
// 	r.HandleFunc("/ping", Ping).Methods("GET")
//     // r.HandleFunc("/upload", nil).Methods("POST")
// 	r.HandleFunc("/upload", UploadImages).Methods("POST")

//     // var vektor1 float64
//     vektor1 := processPicture("./uploads/1700031094273167400.png")
//     // cosine_similarity(vektor1, vektor1)
//     fmt.Printf("%f\n",vektor1)

//     log.Printf("Server is running on http://localhost:%s", PORT)
//     log.Println(http.ListenAndServe(":"+PORT, r))
// }
