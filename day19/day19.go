package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strings"
)

type LocationXYZ struct {
	X int
	Y int
	Z int
}

type Scanner map[LocationXYZ]bool

type DataType []Scanner

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n\n")

	result := make(DataType, len(dataSplit))
	for i, lines := range dataSplit {
		linesSplit := strings.Split(lines, "\n")[1:]
		result[i] = make(map[LocationXYZ]bool, len(linesSplit))
		for _, line := range linesSplit {
			lineSplit := strings.Split(line, ",")
			location := LocationXYZ{
				X: ParseInt(lineSplit[0]),
				Y: ParseInt(lineSplit[1]),
				Z: ParseInt(lineSplit[2]),
			}
			result[i][location] = true
		}
	}

	return result
}

func states() []LocationXYZ {
	XYZ := LocationXYZ{X: 1, Y: 2, Z: 3}
	matrixX := [][]int{{1,0,0},{0,0,-1},{0,1,0}}
	matrixY := [][]int{{0,0,1},{0,1,0},{-1,0,0}}
	matrixZ := [][]int{{0,-1,0},{1,0,0},{0,0,1}}

	rotate := func(location LocationXYZ, matrix [][]int) LocationXYZ {
		return LocationXYZ{
			X: matrix[0][0] * location.X + matrix[0][1] * location.Y + matrix[0][2] * location.Z,
			Y: matrix[1][0] * location.X + matrix[1][1] * location.Y + matrix[1][2] * location.Z,
			Z: matrix[2][0] * location.X + matrix[2][1] * location.Y + matrix[2][2] * location.Z,
		}
	}

	result := map[LocationXYZ]bool{XYZ: true}
	for i := 0; i < 4; i++ {
		newResult := make(map[LocationXYZ]bool)
		for r := range result {
			newResult[r] = true
			newResult[rotate(r, matrixX)] = true
			newResult[rotate(r, matrixY)] = true
			newResult[rotate(r, matrixZ)] = true
		}
		result = newResult
	}

	resultAsSlice := make([]LocationXYZ, 0, len(result))
	for r := range result {
		resultAsSlice = append(resultAsSlice, r)
	}

	return resultAsSlice
}

// parsed from states()
func options() []func(LocationXYZ)LocationXYZ {
	return []func(LocationXYZ)LocationXYZ {
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Z, Y: -l.X, Z: -l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Y, Y:  l.X, Z: -l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Y, Y: -l.X, Z:  l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Y, Y:  l.Z, Z: -l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Y, Y: -l.Z, Z: -l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Z, Y:  l.Y, Z: -l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Z, Y:  l.X, Z:  l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Y, Y: -l.X, Z: -l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.X, Y: -l.Y, Z:  l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Z, Y: -l.Y, Z: -l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Z, Y:  l.Y, Z:  l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.X, Y:  l.Z, Z:  l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Z, Y:  l.X, Z: -l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.X, Y:  l.Y, Z:  l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.X, Y:  l.Z, Z: -l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Y, Y:  l.X, Z:  l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.X, Y: -l.Z, Z:  l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.X, Y:  l.Y, Z: -l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.X, Y: -l.Y, Z: -l.Z}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Z, Y: -l.Y, Z:  l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Z, Y: -l.X, Z:  l.Y}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X:  l.Y, Y:  l.Z, Z:  l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.Y, Y: -l.Z, Z:  l.X}},
		func(l LocationXYZ) LocationXYZ {return LocationXYZ{X: -l.X, Y: -l.Z, Z: -l.Y}},
	}
}

func (xyz LocationXYZ) Add(offset LocationXYZ) LocationXYZ {
	return LocationXYZ{
		X: xyz.X + offset.X,
		Y: xyz.Y + offset.Y,
		Z: xyz.Z + offset.Z,
	}
}

func (scanner Scanner) transform(f func(LocationXYZ)LocationXYZ) Scanner {
	transformedScanner := make(Scanner, len(scanner))
	for el := range scanner {
		transformedScanner[f(el)] = true
	}
	return transformedScanner
}

func hasTwelveCommon(scanner1 Scanner, scanner2 Scanner, offset LocationXYZ) bool {
	c := 0
	for el := range scanner2 {
		if _, ok := scanner1[el.Add(offset)]; ok {
			c++

			if c >= 12 {
				return true
			}
		}
	}

	return false
}

func (scanner Scanner) findOffset(other Scanner) (Scanner, LocationXYZ) {
	for l1 := range scanner {
		for _, option := range options() {
			transformedOtherScanner := other.transform(option)
			for l2 := range transformedOtherScanner {
				offset := LocationXYZ{X: l1.X - l2.X, Y: l1.Y - l2.Y, Z: l1.Z - l2.Z}
				if hasTwelveCommon(scanner, transformedOtherScanner, offset) {
					return transformedOtherScanner, offset
				}
			}
		}
	}

	return nil, LocationXYZ{}
}

func solvePartX(data DataType) (Scanner, []LocationXYZ) {
	uberGrid := make(Scanner, len(data[0]))
	for el := range data[0] {
		uberGrid[el] = true
	}

	scannerOffsets := make(map[int]LocationXYZ, len(data))
	scannerOffsets[0] = LocationXYZ{X: 0, Y: 0, Z: 0}
	for len(scannerOffsets) < len(data) {
		for i, scanner := range data {
			transformedScanner, offset := uberGrid.findOffset(scanner)
			if transformedScanner != nil {
				scannerOffsets[i] = offset
				for el := range transformedScanner {
					uberGrid[el.Add(offset)] = true
				}
			}
		}
	}

	scannerOffsetsAsSlice := make([]LocationXYZ, len(scannerOffsets))
	for i := 0; i < len(scannerOffsets); i++ {
		scannerOffsetsAsSlice[i] = scannerOffsets[i]
	}

	return uberGrid, scannerOffsetsAsSlice
}

func solvePart1(data DataType) (rc int) {
	uberGrid, _ := solvePartX(data)
	return len(uberGrid)
}

func solvePart2(data DataType) (rc int) {
	_, offsets := solvePartX(data)
	for i := 0; i < len(offsets); i++ {
		for j := i + 1; j < len(offsets); j++ {
			dx := offsets[i].X - offsets[j].X
			dy := offsets[i].Y - offsets[j].Y
			dz := offsets[i].Z - offsets[j].Z
			r := Abs(dx) + Abs(dy) + Abs(dz)

			rc = Max(rc, r)
		}
	}
	return
}

func main() {
	data := parseData(FetchInputData(19))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
