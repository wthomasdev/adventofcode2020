package main

import (
	"adventOfCode2020/inputreader"
	"fmt"
	"strings"
)

type stepInstructions struct {
	downSteps  int
	rightSteps int
}

func main() {
	inputPattern := inputreader.ReadStringFile("../data/day03.txt", "\n")
	fmt.Println("Part 1", partOne(inputPattern))
	fmt.Println("Part 2", partTwo(inputPattern))
}

func partOne(inputPattern []string) int {
	stepI := stepInstructions{
		rightSteps: 3,
		downSteps:  1,
	}
	result := make(chan int, 1)
	go countTrees(stepI, inputPattern, result)
	for i := 0; i < 1; i++ {
		select {
		case treeCount := <-result:
			fmt.Println("treeCount", treeCount)
			return treeCount
		}
	}
	return 0
}

func partTwo(inputPattern []string) int {
	routes := []stepInstructions{
		stepInstructions{
			downSteps:  1,
			rightSteps: 1,
		},
		stepInstructions{
			downSteps:  1,
			rightSteps: 3,
		},
		stepInstructions{
			downSteps:  1,
			rightSteps: 5,
		},
		stepInstructions{
			downSteps:  1,
			rightSteps: 7,
		},
		stepInstructions{
			downSteps:  2,
			rightSteps: 1,
		},
	}
	result := make(chan int, len(routes))

	totalTrees := 0
	for _, route := range routes {
		go countTrees(route, inputPattern, result)
	}

	for i := 0; i < len(routes); i++ {
		select {
		case treeCount := <-result:
			if totalTrees == 0 {
				totalTrees += treeCount
			} else {
				totalTrees *= treeCount
			}
		}
	}
	return totalTrees
}

func countTrees(stepI stepInstructions, input []string, done chan int) {
	treeCount := 0
	currRightIndex := 0
	inputLen := len(input)
	for i := 0; i < inputLen; i += stepI.downSteps {
		if i == 0 {
			currRightIndex += stepI.rightSteps
			continue
		}
		splitLine := strings.Split(input[i], "")
		currRightIndex = resetBounds(currRightIndex, splitLine)
		if isItATree(splitLine[currRightIndex]) {
			treeCount++
		}
		currRightIndex += stepI.rightSteps
	}
	done <- treeCount
}

func isItATree(location string) bool {
	if location == "#" {
		return true
	}
	return false
}

func resetBounds(index int, line []string) int {
	lineLength := len(line)
	if index+1 > lineLength {
		newIndex := index - lineLength
		return newIndex
	}
	return index
}
