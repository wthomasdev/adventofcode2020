package main

import (
	inputreader "adventOfCode2019/datautils"
	"fmt"
)

func main() {
	numbers := inputreader.ReadNumberfile("../data/day01a.txt", "\n")
	sortedNumbers := quickSort(numbers)
	fmt.Println("sortedNumbers", sortedNumbers)
	total := calculateThreeEntriesNiave(sortedNumbers)
	fmt.Println("total", total)

}

func calculateThreeEntriesNiave(numbers []int) int {
	targetSum := 2020
	numLength := len(numbers)
	for i := 0; i < numLength; i++ {
		firstNum := numbers[i]
		for j := i + 1; j < numLength; j++ {
			secondNum := numbers[j]
			for k := j + 1; k < numLength; k++ {
				thirdNum := numbers[k]
				if firstNum+secondNum+thirdNum == targetSum {
					fmt.Println("found combination", firstNum, secondNum, thirdNum)
					return firstNum * secondNum * thirdNum
				}
			}
		}
	}
	return 0
}

func quickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}
