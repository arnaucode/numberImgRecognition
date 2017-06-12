package main

func comparePixels(p []int, q []int) int {
	numCoincidences := 0
	for i := 0; i < len(p); i++ {
		if p[i] == q[i] {
			numCoincidences++
		}
	}
	return numCoincidences
}

func compareImgs(a imgRGBA, b imgRGBA) int {
	numCoincidences := 0

	for i := 0; i < len(a); i++ {
		numCoincidences = numCoincidences + comparePixels(a[i], b[i])
	}
	return numCoincidences
}

func analyzeImg(examplesImg map[string][]imgRGBA, timg imgRGBA) map[string]int {
	numCoincidences := make(map[string]int)

	for k, imgs := range examplesImg {
		numCoincidences[k] = 0
		//fmt.Println(k)
		for _, img := range imgs {
			numCoincidences[k] = numCoincidences[k] + compareImgs(img, timg)
		}
	}
	return numCoincidences
}
