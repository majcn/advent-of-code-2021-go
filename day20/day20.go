package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strconv"
	"strings"
)

type DataType struct {
	algorithm []int
	image map[Location]int
	imageSize Location
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n\n")

	dataSplitAlgorithm := dataSplit[0]
	algorithm := make([]int, len(dataSplitAlgorithm))
	for i, v := range dataSplitAlgorithm {
		if v == '#' {
			algorithm[i] = 1
		} else {
			algorithm[i] = 0
		}
	}

	dataSplitImage := strings.Split(dataSplit[1], "\n")
	image := make(map[Location]int)
	for y, line := range dataSplitImage {
		for x, v := range line {
			if v == '#' {
				image[Location{X: x, Y: y}] = 1
			} else {
				image[Location{X: x, Y: y}] = 0
			}
		}
	}

	imageSize := Location{X: len(dataSplitImage[0]), Y: len(dataSplitImage)}

	return DataType{
		algorithm: algorithm,
		image:     image,
		imageSize: imageSize,
	}
}

func getAlgorithmIndex(grid map[Location]int, location Location, background int) int {
	neighbours := []Location{
		{X: -1, Y: -1},
		{X:  0, Y: -1},
		{X:  1, Y: -1},
		{X: -1, Y:  0},
		{X:  0, Y:  0},
		{X:  1, Y:  0},
		{X: -1, Y:  1},
		{X:  0, Y:  1},
		{X:  1, Y:  1},
	}

	kernelBin := make([]byte, 9)
	for i, n := range neighbours {
		if el, ok := grid[location.Add(n)]; ok {
			kernelBin[i] = strconv.Itoa(el)[0]
		} else {
			kernelBin[i] = strconv.Itoa(background)[0]
		}
	}

	ui, _ := strconv.ParseUint(string(kernelBin), 2, 9)
	return int(ui)
}

func solvePartX(data DataType, enhanceCounter int) (rc int) {
	algorithm, image, imageSize := data.algorithm, data.image, data.imageSize

	for t := 1; t <= enhanceCounter; t++ {
		newImage := make(map[Location]int)
		for x := -2; x < imageSize.X + 2; x++ {
			for y := -2; y < imageSize.Y + 2; y++ {
				imageLocation := Location{X: x, Y: y}
				newImageLocation := imageLocation.Add(Location{X: 2, Y: 2})
				backgroundColor := (t + 1) % 2

				algorithmIndex := getAlgorithmIndex(image, imageLocation, backgroundColor)
				newImage[newImageLocation] = algorithm[algorithmIndex]
			}
		}
		image = newImage
		imageSize = imageSize.Add(Location{X: 4, Y: 4})
	}

	for _, v := range image {
		if v == 1 {
			rc++
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	return solvePartX(data, 2)
}

func solvePart2(data DataType) (rc int) {
	return solvePartX(data, 50)
}

func main() {
	data := parseData(FetchInputData(20))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
