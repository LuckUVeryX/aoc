package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

const target = 2020

func main() {
	part1 := twoSum()
	part2 := threeSum()

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}

func twoSum() int {
	input := utils.GetInput(1)
	lines := strings.Split(input, "\n")

	nums := map[int]int{}

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		if n, ok := nums[num]; ok {
			return n * num
		} else {
			nums[target-num] = num
		}
	}
	panic("no solution")
}

func threeSum() int {
	input := utils.GetInput(1)
	lines := strings.Split(input, "\n")

	nums := make([]int, len(lines))
	for i, lines := range lines {
		num, _ := strconv.Atoi(lines)
		nums[i] = num
	}

	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == target {
				return nums[i] * nums[j] * nums[k]
			}
			if nums[i]+nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}

	return 0
}
