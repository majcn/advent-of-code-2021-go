package main

import (
	"fmt"
	. "majcn.si/advent-of-code-2021/util"
	"strconv"
)

type DataType []byte

func parseData(data string) DataType {
	result := make(DataType, len(data) * 4)
	for i, char := range data {
		ui, _ := strconv.ParseUint(string(char), 16, 8)
		bin := fmt.Sprintf("%04b", ui)
		copy(result[(i*4):(i*4)+4], bin)
	}

	return result
}

type Package []byte
func (p Package) asInt() int {
	ui, _ := strconv.ParseUint(string(p), 2, 64)
	return int(ui)
}
func (p Package) getVersion() int{
	return p[:3].asInt()
}
func (p Package) getType() int {
	return p[3:6].asInt()
}
func (p Package) getLiteralValue() (int, int) {
	result := make(Package, 0)
	pointer := 6
	for true {
		result = append(result, p[pointer+1:pointer+5]...)
		pointer += 5

		if p[pointer-5] == '0' {
			break
		}
	}

	return result.asInt(), pointer
}
func (p Package) getOperatorValue() (int, int, int) {
	versionSum, packageLength, result := 0, 0, 0

	resultOfSubPackages := make([]int, 0)

	if p[6] == '0' {
		subPackageLength := p[7:22].asInt()
		offset := 0

		for offset < subPackageLength {
			subPackageVersionSum, subPackageResult, subPackagePackageLength := p[22+offset : 22+subPackageLength].parse()

			offset += subPackagePackageLength
			versionSum += subPackageVersionSum

			resultOfSubPackages = append(resultOfSubPackages, subPackageResult)
		}
		packageLength = 22 + subPackageLength
	} else {
		subPackageCount := p[7:18].asInt()
		offset := 0

		for i := 0; i < subPackageCount; i++ {
			subPackageVersionSum, subPackageResult, subPackagePackageLength := p[18+offset:].parse()

			offset += subPackagePackageLength
			versionSum += subPackageVersionSum

			resultOfSubPackages = append(resultOfSubPackages, subPackageResult)
		}
		packageLength = 18 + offset
	}

	switch p.getType() {
	case 0:
		result = Sum(resultOfSubPackages)
	case 1:
		result = 1
		for _, x := range resultOfSubPackages {
			result *= x
		}
	case 2:
		result = Min(resultOfSubPackages...)
	case 3:
		result = Max(resultOfSubPackages...)
	case 5:
		if resultOfSubPackages[0] > resultOfSubPackages[1] {
			result = 1
		} else {
			result = 0
		}
	case 6:
		if resultOfSubPackages[0] < resultOfSubPackages[1] {
			result = 1
		} else {
			result = 0
		}
	case 7:
		if resultOfSubPackages[0] == resultOfSubPackages[1] {
			result = 1
		} else {
			result = 0
		}
	}

	return versionSum, result, packageLength
}
func (p Package) parse() (int, int, int) {
	versionSum, result, packageLength := 0, 0, 0

	if p.getType() == 4 {
		result, packageLength = p.getLiteralValue()
	} else {
		versionSum, result, packageLength = p.getOperatorValue()
	}

	versionSum += p.getVersion()

	return versionSum, result, packageLength
}

func solvePart1(data DataType) (rc int) {
	rc, _, _ = Package(data).parse()
	return
}

func solvePart2(data DataType) (rc int) {
	_, rc, _ = Package(data).parse()
	return
}

func main() {
	data := parseData(FetchInputData(16))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
