package main

func main() {
	//first, read the example images
	examplesImg := readExamples()
	//fmt.Println(examplesImg)

	//now, read the target image
	timg := readImage("images/test.png")
	//fmt.Println(timg)

	numCoincidences := analyzeImg(examplesImg, timg)
	_ = printSortedMapStringInt(numCoincidences, 0)

}
