package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputEmbed embed.FS

func main() {
	input, err := inputEmbed.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	leftNumList := []int{}
	rightNumList := []int{}

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		lineSplit := strings.Split(line, "   ")

		leftNum, _ := strconv.Atoi(lineSplit[0])
		rightNum, _ := strconv.Atoi(lineSplit[1])

		leftNumList = append(leftNumList, leftNum)
		rightNumList = append(rightNumList, rightNum)
	}

	sort.Ints(leftNumList)
	sort.Ints(rightNumList)

	totalDiff := 0
	for i := range leftNumList {
		totalDiff += abs(leftNumList[i] - rightNumList[i])
	}

	fmt.Printf("The total difference between the two lists is: %d\n", totalDiff)

	similarity := 0
	for _, leftNum := range leftNumList {
		for _, rightNum := range rightNumList {
			if leftNum == rightNum {
				similarity += leftNum
			}
		}
	}

	fmt.Printf("The similarity between the two lists is: %d\n", similarity)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
