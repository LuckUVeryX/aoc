package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(24)
	packages := []int{}
	for _, line := range strings.Split(input, "\n") {
		pkg, _ := strconv.Atoi(line)
		packages = append(packages, pkg)
	}

	totalWeight := 0
	for _, pkg := range packages {
		totalWeight += pkg
	}
	targetWeight := totalWeight / 4
	fmt.Println("Target Weight:", targetWeight)

	sort.Sort(sort.Reverse(sort.IntSlice(packages)))

	subsets := [][]int{}
	// Finds all the valid packages combinations that make target weight
	var backtrack func(start, target int, curr []int)
	backtrack = func(start, target int, curr []int) {
		if target == 0 {
			subsets = append(subsets, append([]int(nil), curr...))
			return
		}

		for i := start; i < len(packages); i++ {
			if packages[i] > target {
				continue
			}

			backtrack(i+1, target-packages[i], append(curr, packages[i]))
		}
	}

	backtrack(0, targetWeight, []int{})

	minPackages := len(packages) - 1
	QE := math.MaxInt
	for _, subset := range subsets {
		if len(subset) < minPackages {
			minPackages = len(subset)
			QE = quantumEntanglement(subset)
		} else if len(subset) == minPackages {
			qe := quantumEntanglement(subset)
			if qe < QE {
				QE = qe
			}
		}
	}

	fmt.Println("QE:", QE)
}

func quantumEntanglement(group []int) int {
	quantum := 1
	for _, w := range group {
		quantum *= w
	}
	return quantum
}
