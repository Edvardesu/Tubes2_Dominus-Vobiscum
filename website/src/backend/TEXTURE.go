package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"

	// _ "image/svg"
	"log"
	"os"

	// "github.com/lucasb-eyer/go-colorful"
	"math"
	"sync"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func saveImageGray(filename string, pixels [][]color.Color)  {
	rect := image.Rect(0,0,len(pixels),len(pixels[0]))
	nImg := image.NewRGBA(rect)

	for x:=0; x<len(pixels);x++{
		for y:=0; y<len(pixels[0]);y++ {
			q:=pixels[x]
			if q==nil{
				continue
			}
			p := pixels[x][y]
			if p==nil{
				continue
			}
			original,ok := color.RGBAModel.Convert(p).(color.RGBA)
			if ok{
				nImg.Set(x,y,original)
			}
		}
	}

	fg, err:= os.Create("testing.jpg")
	if err!=nil{
		fmt.Println("Creating file:",err)
	}
	defer fg.Close()

	err = jpeg.Encode(fg, nImg, nil)
	if err != nil {
		fmt.Println("Encoding image:", err)
	}
}

func greyScale(pixels *[][]color.Color)  {
	ppixels := *pixels
	xLen:=len(ppixels)
	yLen := len(ppixels[0])

	//create new image
	newImage:=make([][]color.Color, xLen)
	for i:=0;i<len(newImage);i++{
	newImage[i] = make([]color.Color,yLen)
	}
	//idea is processing pixels in parallel
	wg := sync.WaitGroup{}

	for x:=0;x<xLen;x++{
		for y:=0;y<yLen; y++{
			wg.Add(1)
			go func(x,y int) {
				pixel :=ppixels[x][y]
				originalColor,ok := color.RGBAModel.Convert(pixel).(color.RGBA)
				if !ok{
					fmt.Println("type conversion went wrong")
				}
				grey := uint8(float64(originalColor.R)*0.299 + float64(originalColor.G)*0.587 + float64(originalColor.B)*0.114)
				col :=color.RGBA{
					grey,
					grey,
					grey,
					originalColor.A,
				}
				newImage[x][y] = col
				wg.Done()
			}(x,y)

		}
	}
	wg.Wait()
	*pixels = newImage
}

func greyScale2(matrix *[][]uint8, img *image.Image){
	bounds := (*img).Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			colorAt := (*img).At(x, y)

			c, _ := color.RGBAModel.Convert(colorAt).(color.RGBA)

			(*matrix)[x][y] = uint8(float64(c.R)*0.299 + float64(c.G)*0.587 + float64(c.B)*0.114)
		}
	}
}

func featureExtraction(fm *mat.Dense) []float64 {
	rows, _ := fm.Dims()

	// Create a matrix_index slice
	matrixIndex := make([]float64, rows)
	for i := range matrixIndex {
		matrixIndex[i] = float64(i)
	}

	// Convert matrix_index to a column vector (column-major order)
	var colVec mat.VecDense
	colVec.CloneFromVec(mat.NewVecDense(len(matrixIndex), matrixIndex))

	// Calculate m_square
	var mSquare mat.Dense
	mSquare.Apply(func(i, j int, v float64) float64 {
		return (v - colVec.At(i, 0)) * (v - colVec.At(j, 0))
	}, fm)


	rawMatrix := fm.RawMatrix()
	datafm := rawMatrix.Data

	raw_mSquare := mSquare.RawMatrix()
	data_mSquare := raw_mSquare.Data


	var contrast float64
	for i := 0; i < len(datafm); i++ {
		contrast += data_mSquare[i] * datafm[i]
	}

	var homogeneity float64
	for i := 0; i < len(datafm); i++ {
		homogeneity += datafm[i] / (1 + data_mSquare[i])
	}

	var entropy float64
	for i := 0; i < len(datafm); i++ {
		if datafm[i] == 0 {
			entropy += 0
		} else {
			entropy += datafm[i] * math.Log(datafm[i])
		}
	}
	entropy = -entropy

	return []float64{contrast, homogeneity, entropy}
}

func createCoOccurence(matrix [][]uint8)(*mat.Dense){
	occurence0 := mat.NewDense(256, 256, nil)

	// 0 angle
	// bounds := matrix.Bounds()
	xlen := len(matrix)
	ylen := len(matrix[0])

	// for y := 0; y < bounds.Max.Y- 1; y++{
	for y := 0; y < ylen - 1; y++{
		// for x := 0; x < bounds.Max.X; x++ {
		for x := 0; x < xlen; x++ {
			i := matrix[x][y]
			j := matrix[x][y+1]
			occurence0.Set(int(i), int(j), occurence0.At(int(i), int(j)) + 1)
		}
	}

	transpose0 := mat.DenseCopyOf(occurence0.T())

	var sum0 mat.Dense
	sum0.Add(occurence0, transpose0)
	sum0Matrix := mat.Matrix(&sum0)

	fmt.Println(sum0Matrix.At(0, 0))


	data0 := sum0.RawMatrix().Data 
	norm_val := floats.Sum(data0)
	
	// fmt.Println(sum0Matrix.   Dims())
	normalizedMatrix0 := mat.NewDense(256, 256, nil)
	// fmt.Println(normalizedMatrix0.Dims())
	normalizedMatrix0.Scale(1/norm_val, sum0Matrix)

	return normalizedMatrix0
}

func processPicture_texture(filename string) []float64 {
	// read file

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }

	xLen := img.Bounds().Max.X 
	yLen := img.Bounds().Max.Y

	matrix := make([][]uint8, xLen)
	for i := range matrix {
        matrix[i] = make([]uint8, yLen)
    }

	greyScale2(&matrix, &img)
	// fmt.Println(matrix)

	// saveImageGray("testing.jpeg", pixels)

	occ_matrix := createCoOccurence(matrix)

	vecktor := featureExtraction(occ_matrix)

	return vecktor
	// return matrix
    
	// fmt.Println(img)
}

func cosine_similarity_texture(vektor1 []float64, vektor2 []float64) float64{
	vecDense1 := mat.NewVecDense(len(vektor1), vektor1)
	vecDense2 := mat.NewVecDense(len(vektor2), vektor2)
	dotProduct := mat.Dot(vecDense1, vecDense2)
	return dotProduct / (mat.Norm(vecDense1,2 ) * mat.Norm(vecDense2,2))
}


// func main(){
// 	start := time.Now()

// 	vektor1 := processPicture("../images/lena.png")
// 	fmt.Println(vektor1)

// 	vektor2 := processPicture("../images/lena.png")

// 	var simmilarity float64
// 	simmilarity = cosine_similarity(vektor1, vektor2)

// 	elapsed := time.Since(start)
// 	fmt.Println("Execution time :", elapsed)
// 	fmt.Println("Simmilarity :", simmilarity )
// }