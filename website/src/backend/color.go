package main

import (
	"fmt"
	"image"
	_"image/png"
	_"image/jpeg"
	_"image/gif"
	"image/color"
	// "github.com/lucasb-eyer/go-colorful"
	// "gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math"
	// "time"
	"os"
)


func rgbToHSV(rgb color.RGBA) (float64, float64, float64){
	r := float64(rgb.R) / 255.0
	g := float64(rgb.G) / 255.0
	b := float64(rgb.B) / 255.0

	max := max(max(r, g), b)
	min := min(min(r, g), b)

	var h, s, v float64

	// Hue calculation
	if max == min {
		h = 0
	} else if max == r {
		h = math.Mod((60 * ((g - b) / (max - min)) + 360), 360)
	} else if max == g {
		h = math.Mod((60 * ((b - r) / (max - min)) + 120), 360)
	} else if max == b {
		h = math.Mod((60 * ((r - g) / (max - min)) + 240), 360)
	}

	// Saturation calculation
	if max == 0 {
		s = 0
	} else {
		s = (max - min) / max
	}

	// Value calculation
	v = max

	return h, s, v
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func processPixel(rgba color.RGBA) (float64, float64, float64) {

	// Convert RGB to HSV
	h, s, v := rgbToHSV(rgba)

	// Print the result
	// fmt.Printf("RGB(%d, %d, %d) converted to HSV: H=%.2f, S=%.2f, V=%.2f\n",
	// 	rgb.R, rgb.G, rgb.B, hsv.H, hsv.S, hsv.V)
	
	return h, s, v
}

func checkH(h float64)(int){
	if ((316 <= h && h <= 360) || h == 0){
		return 0
	} else if (1 <= h && h <= 25){
		return 1
	} else if (26 <= h  && h <= 40){
		return 2
	} else if (41 <= h && h <= 120){
		return 3
	} else if (121 <= h && h <= 190){
		return 4
	} else if (191 <= h && h <= 270){
		return 5
	} else if (271 <= h && h <= 295){
		return 6
	} else if (296 <= h && h <= 315){
		return 7
	}
	return 0;
}


func checkS(s float64)(int){
	if (0 <= s && s < 0.2){
		return 0
	} else if (0.2 <= s && s < 0.7){
		return 1
	} else if (0.7 <= s && s <= 1){
		return 2
	}
	return 0;
}

func checkV(v float64)(int){
	if (0 <= v && v < 0.2){
		return 0
	} else if (0.2 <= v && v < 0.7){
		return 1
	} else if (0.7 <= v && v <= 1){
		return 2
	}
	return 0;
}

func processRegion(img image.Image, xStart int, yStart int, xEnd int , yEnd int)(*mat.VecDense){
	var array_h [8]float64
	var array_s [3]float64
	var array_v [3]float64

	// Iterate over each row and column
	for y := yStart; y < yEnd; y++ {

		for x := xStart; x < xEnd; x++ {
			// Get the color at the current pixel
			colorAt := img.At(x, y)

			// Convert RGB to HSV using the go-colorful library
			c, _ := color.RGBAModel.Convert(colorAt).(color.RGBA)
			// fmt.Println(c.R)
			h, s, v := processPixel(c)
			h = math.Round(h)
			// check h s v value
			index_h := checkH(h)
			index_s := checkS(s)
			index_v := checkV(v)
			array_h[index_h] += 1
			array_s[index_s] += 1
			array_v[index_v] += 1
		}
	}
	
	
	// create histogram
	// fmt.Println(array_h)
	// fmt.Println(array_s)
	// fmt.Println(array_v)

	newValues:= []float64{array_h[0], array_h[1], array_h[2], array_h[3], array_h[4], array_h[5], array_h[6], array_h[7], array_s[0], array_s[1], array_s[2], array_v[0], array_v[1], array_v[2]}

	return mat.NewVecDense(14, newValues)
}

func processPicture(filename string)([4][4]*mat.VecDense){

	var arr_vector [4][4]*mat.VecDense

	// Open the image file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return arr_vector
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return arr_vector
	}

	// split picture


	// Get the image bounds
	bounds := img.Bounds()

	// Define the size of each region
	regionSizeX := bounds.Dx() / 4
	regionSizeY := bounds.Dy() / 4

	// Iterate over each region
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			// Define the boundaries of the current region
			xStart := i * regionSizeX
			yStart := j * regionSizeY
			xEnd := (i + 1) * regionSizeX
			yEnd := (j + 1) * regionSizeY

			// Process the current region
			arr_vector[i][j] = processRegion(img, xStart, yStart, xEnd, yEnd)
		}
	}
	
	return arr_vector
	
}

func cosine_similarity(vektor1 [4][4]*mat.VecDense, vektor2 [4][4]*mat.VecDense)(float64){
	var total float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			total = total + mat.Dot(vektor1[i][j], vektor2[i][j]) / (mat.Norm(vektor1[i][j], 2) * mat.Norm(vektor2[i][j], 2))
		}
	}

	return  total / 16
}

// func main(){
// 	start := time.Now()

// 	vektor1 := processPicture("../images/marvel_kucing1.jpg")

// 	vektor2 := processPicture("../images/marvel_kucing3.jpg")

// 	var simmilarity float64
// 	simmilarity = cosine_similarity(vektor1, vektor2)

// 	elapsed := time.Since(start)
// 	fmt.Println("Execution time :", elapsed)

// 	fmt.Println("Simmilarity :", simmilarity )
// }