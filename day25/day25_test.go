package main

import (
	. "../util"
	"testing"
)

func ExpectedPart1() int {
	return 471
}

var data DataType
func init() {
	data = InitData(func(data string) interface{} { return parseData(data) }).(DataType)
}

func TestPart1(t *testing.T) {
	AssertTestPartX(t, ExpectedPart1(), solvePart1(data))
}
