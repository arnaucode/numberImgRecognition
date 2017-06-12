package main

import "fmt"

func main() {
	//first, read the example images
	examplesImg := readExamples()
	//fmt.Println(examplesImg)

	//now, read the target image
	timg := readImage("images/test.png")
	//fmt.Println(timg)

	numCoincidences := analyzeImg(examplesImg, timg)
	_, result := printSortedMapStringInt(numCoincidences, 0)
	fmt.Print("The image contains the number: ")
	c.Green(result)
	//runServer()
}
