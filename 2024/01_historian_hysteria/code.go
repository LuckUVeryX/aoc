package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(1)

	lines := strings.Split(strings.TrimSpace(input), "\n")

	var left, right []int

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		firstNum, _ := strconv.Atoi(parts[0])
		secondNum, _ := strconv.Atoi(parts[1])

		left = append(left, firstNum)
		right = append(right, secondNum)
	}

	// * Part 2
	rightMap := make(map[int]int)
	for i := 0; i < len(right); i++ {
		rightMap[right[i]]++
	}

	similarity := 0

	for i := 0; i < len(left); i++ {
		similarity += left[i] * rightMap[left[i]]
	}

	fmt.Println("Similarity:", similarity)

	// * Part 1
	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff *= -1
		}
		distance += diff
	}

	fmt.Println("Distance:", distance)

}
