package main

import (
	. "../util"
	"testing"
)

func ExpectedPart1() int {
	return 3497
}

func ExpectedPart2() int {
	return 93686
}

var data DataType
func init() {
	data = InitData(func(data string) interface{} { return parseData(data) }).(DataType)
}

func TestPart1(t *testing.T) {
	AssertTestPartX(t, ExpectedPart1(), solvePart1(data))
}

func TestPart2(t *testing.T) {
	AssertTestPartX(t, ExpectedPart2(), solvePart2(data))
}