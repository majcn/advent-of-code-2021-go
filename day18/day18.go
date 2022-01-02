package main

import (
	. "../util"
	"fmt"
	"strings"
)

type DataType []string

func parseData(data string) DataType {
	dataSplit := strings.Split(data, "\n")

	result := make(DataType, len(dataSplit))
	for i, line := range dataSplit {
		result[i] = line
	}

	return result
}

type Node struct {
	elements [2]*Node
	value int
}

func NewNode(s string) (node *Node) {
	var createTreeFromStringInner func(string)(*Node,int)
	createTreeFromStringInner = func(s string)(node *Node, i int) {
		node = &Node{}
		nodeElementsIndex := 0

		for i = 1; i < len(s); i++ {
			if s[i] == '[' {
				newNode, newI := createTreeFromStringInner(s[i:])
				i += newI

				node.elements[nodeElementsIndex] = newNode
				nodeElementsIndex++
				continue
			}

			if s[i] == ']' {
				return
			}

			if s[i] >= '0' && s[i] <= '9' {
				newNode := &Node{value: ParseInt(s[i])}

				node.elements[nodeElementsIndex] = newNode
				nodeElementsIndex++
			}
		}

		return
	}

	result, _ := createTreeFromStringInner(s)
	return result
}

func (node *Node) isValueNode() bool {
	return node.elements[0] == nil && node.elements[1] == nil
}

func (node *Node) walkForExplode() (*Node, *Node, *Node) {
	var targetNode, prevValueNode, nextValueNode *Node

	var walkForExplodeInner func(*Node, int)
	walkForExplodeInner = func (node *Node, depth int) {
		if targetNode == nil && depth == 4 {
			targetNode = node
			return
		}

		for _, lfNode := range node.elements {
			if lfNode.isValueNode() {
				if targetNode == nil {
					prevValueNode = lfNode
				} else {
					nextValueNode = lfNode
				}
			} else {
				walkForExplodeInner(lfNode, depth + 1)
			}

			if nextValueNode != nil {
				return
			}
		}
	}

	walkForExplodeInner(node, 0)
	return targetNode, prevValueNode, nextValueNode
}

func (node *Node) explode() bool {
	targetNode, prevValueNode, nextValueNode := node.walkForExplode()

	if targetNode == nil {
		return false
	}

	if prevValueNode != nil {
		prevValueNode.value += targetNode.elements[0].value
	}
	if nextValueNode != nil {
		nextValueNode.value += targetNode.elements[1].value
	}
	*targetNode = Node{value: 0}

	return true
}

func (node *Node) walkForSplit() *Node {
	var targetNode *Node

	var walkForSplitInner func(*Node)
	walkForSplitInner = func (node *Node) {
		if targetNode != nil {
			return
		}

		for _, lfNode := range node.elements {
			if lfNode.isValueNode() {
				if targetNode == nil && lfNode.value >= 10 {
					targetNode = lfNode
					return
				}
			} else {
				walkForSplitInner(lfNode)
			}
		}
	}

	walkForSplitInner(node)
	return targetNode
}

func (node *Node) split() bool {
	targetNode := node.walkForSplit()

	if targetNode == nil {
		return false
	}

	leftNode := Node{value: targetNode.value / 2}
	rightNode := Node{value: targetNode.value - leftNode.value}

	*targetNode = Node{elements: [2]*Node{&leftNode, &rightNode}}

	return true
}

func (node *Node) reduce() {
	exploded := false
	for node.explode() {
		exploded = true
	}

	if node.split() || exploded {
		node.reduce()
	}
}

func (node *Node) magnitude() (rc int) {
	multiplier := [2]int{3, 2}

	for i, el := range node.elements {
		if el.isValueNode() {
			rc += multiplier[i] * el.value
		} else {
			rc += multiplier[i] * el.magnitude()
		}
	}

	return
}

func solvePart1(data DataType) (rc int) {
	node := NewNode(data[0])
	for i := 1; i < len(data); i++ {
		node = &Node{elements: [2]*Node{node, NewNode(data[i])}}
		node.reduce()
	}

	return node.magnitude()
}

func solvePart2(data DataType) (rc int) {
	for _, pair := range StringPermutations(data, 2) {
		node := &Node{elements: [2]*Node{NewNode(pair[0]), NewNode(pair[1])}}
		node.reduce()
		rc = Max(rc, node.magnitude())
	}

	return
}

func main() {
	data := parseData(FetchInputData(18))
	fmt.Println(solvePart1(data))
	fmt.Println(solvePart2(data))
}
