package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(2)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	safeCount := 0
	for _, line := range lines {
		var levels []int
		for _, v := range strings.Split(line, " ") {
			value, _ := strconv.Atoi(v)
			levels = append(levels, value)
		}

		if isReportSafe(levels) {
			safeCount++
		}
	}

	fmt.Println("Number of safe reports:", safeCount)
}

func isReportSafe(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		if isSafeDampened(levels, i) {
			return true
		}
	}

	return false
}

func isSafeDampened(levels []int, i int) bool {
	modLevels := []int{}

	modLevels = append(modLevels, levels[:i]...)
	modLevels = append(modLevels, levels[i+1:]...)

	return isSafe(modLevels)
}

func isSafe(levels []int) bool {

	isIncreasing := levels[0] < levels[1]

	if isIncreasing {
		for i := 1; i < len(levels); i++ {
			diff := levels[i] - levels[i-1]
			if diff >= 1 && diff <= 3 {
				continue
			}
			return false
		}
	} else {
		for i := 1; i < len(levels); i++ {
			diff := levels[i-1] - levels[i]
			if diff >= 1 && diff <= 3 {
				continue
			}
			return false
		}
	}

	return true
}
