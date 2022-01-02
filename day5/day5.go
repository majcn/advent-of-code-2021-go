package main

import (
	. "../util"
	"fmt"
	"regexp"
	"strings"
)

type Vent struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

type DataType []Vent

func fromLine(s string) Vent {
	re := `^(\d+),(\d+) -> (\d+),(\d+)$`
	m := regexp.MustCompile(re).FindStringSubmatch(s)

	return Vent{
		X1: ParseInt(m[1]),
		X2: ParseInt(m[3]),
		Y1: ParseInt(m[2]),
		Y2: ParseInt(m[4]),
	}
}

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, v := range dataSplit {
		result[i] = fromLine(v)
	}

	return result
}

func getGridHorizontalAndVertical(data DataType) map[Location]int {
	grid := make(map[Location]int)

	for _, d := range data {
		if d.X1 == d.X2 {
			start := Min(d.Y1, d.Y2)
			end := Max(d.Y1, d.Y2)
			for y := start; y <= end; y++ {
				grid[Location{X: d.X1, Y: y}]++
			}
		}

		if d.Y1 == d.Y2 {
			start := Min(d.X1, d.X2)
			end := Max(d.X1, d.X2)
			for x := start; x <= end; x++ {
				grid[Location{X: x, Y: d.Y1}]++
			}
		}
	}

	return grid
}

func getGridDiagonal(data DataType) map[Location]int {
	grid := make(map[Location]int)

	for _, d := range data {
		if d.X1 < d.X2 && d.Y1 < d.Y2 {
			x, y := d.X1, d.Y1
			for i := 0; i <= d.X2-d.X1; i++ {
				grid[Location{X: x, Y: y}]++
				x++; y++
			}
		}

		if d.X1 > d.X2 && d.Y1 < d.Y2 {
			x, y := d.X1, d.Y1
			for i := 0; i <= d.X1-d.X2; i++ {
				grid[Location{X: x, Y: y}]++
				x--; y++
			}
		}

		if d.X1 < d.X2 && d.Y1 > d.Y2 {
			x, y := d.X1, d.Y1
			for i := 0; i <= d.X2-d.X1; i++ {
				grid[Location{X: x, Y: y}]++
				x++; y--
			}
		}

		if d.X1 > d.X2 && d.Y1 > d.Y2 {
			x, y := d.X1, d.Y1
			for i := 0; i <= d.X1-d.X2; i++ {
				grid[Location{X: x, Y: y}]++
				x--; y--
			}
		}
	}

	return grid
}

func solvePart1(data DataType) (rc int) {
	grid := getGridHorizontalAndVertical(data)
	for _, v := range grid {
		if v > 1 {
			rc++
		}
	}
	return
}

func solvePart2(data DataType) (rc int) {
	grid := make(map[Location]int)
	for k, v := range getGridHorizontalAndVertical(data) {
		grid[k] += v
	}
	for k, v := range getGridDiagonal(data) {
		grid[k] += v
	}

	for _, v := range grid {
		if v > 1 {
			rc++
		}
	}

	return
}

func main() {
	data := parseData(FetchInputData(5))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
