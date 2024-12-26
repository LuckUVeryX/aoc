package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(7)
	lines := strings.Split(input, "\n")

	res := 0
	for _, line := range lines {
		res += parse1(line)
	}

	fmt.Println("Result:", res)

	res = 0
	for _, line := range lines {
		res += parse2(line)
	}

	fmt.Println("Result:", res)

}

func parse1(line string) int {
	parts := strings.Split(line, ": ")
	target, _ := strconv.Atoi(parts[0])

	testValues := strings.Split(parts[1], " ")
	ints := make([]int, len(testValues))
	for i, v := range testValues {
		ints[i], _ = strconv.Atoi(v)
	}
	var dfs func(target, curr, idx int, values []int) bool

	dfs = func(target, curr, idx int, values []int) bool {
		if idx == len(values) {
			return curr == target
		}

		if curr > target {
			return false
		}

		return dfs(target, curr+values[idx], idx+1, values) || dfs(target, curr*values[idx], idx+1, values)
	}

	if dfs(target, 0, 0, ints) {
		return target
	}
	return 0
}

func parse2(line string) int {
	parts := strings.Split(line, ": ")
	target, _ := strconv.Atoi(parts[0])

	testValues := strings.Split(parts[1], " ")
	ints := make([]int, len(testValues))
	for i, v := range testValues {
		ints[i], _ = strconv.Atoi(v)
	}

	var dfs func(target, curr, idx int, values []int) bool

	dfs = func(target, curr, idx int, values []int) bool {
		if idx == len(values) {
			return curr == target
		}

		if curr > target {
			return false
		}

		concatenate := strconv.Itoa(curr) + strconv.Itoa(values[idx])
		value, _ := strconv.Atoi(concatenate)
		return dfs(target, value, idx+1, values) || dfs(target, curr+values[idx], idx+1, values) || dfs(target, curr*values[idx], idx+1, values)

	}
	if dfs(target, 0, 0, ints) {
		return target
	}
	return 0
}
