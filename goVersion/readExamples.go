package main

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
)

type imgRGBA [][]int
type CollectionImgs []imgRGBA
type ExamplesImg map[string]CollectionImgs

func readImage(path string) imgRGBA {
	//open image, and generate the histogram, made from r,g,b,a components
	reader, err := os.Open(path)
	check(err)
	defer reader.Close()

	m, _, err := image.Decode(reader)
	check(err)
	bounds := m.Bounds()

	//generate the histogram
	var histogram imgRGBA
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			var pixel []int
			pixel = append(pixel, int(r), int(g), int(b), int(a))
			histogram = append(histogram, pixel)
		}
	}
	return histogram
}
func readExamples() map[string][]imgRGBA {
	//var examplesImg ExamplesImg
	examplesImg := make(map[string][]imgRGBA)
	//for each image in the directory of example images, read the image and generate the histogram
	files, _ := ioutil.ReadDir("./images/numbers/")
	for _, f := range files {
		//fmt.Println("images/numbers/" + f.Name())
		file := readImage("images/numbers/" + f.Name())
		numberName := strings.Split(f.Name(), ".")[0]
		//fmt.Println(numberName)
		examplesImg[numberName] = append(examplesImg[numberName], file)
	}
	return examplesImg

}
