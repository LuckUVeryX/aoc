package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(5)
	parts := strings.Split(input, "\n\n")

	rules := map[string][]string{}

	for _, rule := range strings.Split(parts[0], "\n") {
		r := strings.Split(rule, "|")
		rules[r[0]] = append(rules[r[0]], r[1])
	}

	cmp := func(a, b string) int {
		for _, v := range rules[a] {
			if v == b {
				return -1
			}
		}
		return 0
	}

	ordered := 0
	for _, update := range strings.Split(parts[1], "\n") {
		if nums := strings.Split(update, ","); slices.IsSortedFunc(nums, cmp) {
			n, _ := strconv.Atoi(nums[len(nums)/2])
			ordered += n
		}
	}

	fmt.Println("Ordered:", ordered)

	unordered := 0
	for _, update := range strings.Split(parts[1], "\n") {
		if nums := strings.Split(update, ","); !slices.IsSortedFunc(nums, cmp) {
			slices.SortFunc(nums, cmp)
			n, _ := strconv.Atoi(nums[len(nums)/2])
			unordered += n
		}
	}

	fmt.Println("Unordered:", unordered)
}
