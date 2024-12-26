package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(19)

	lines := strings.Split(input, "\n\n")

	patterns := []string{}
	for _, pattern := range strings.Split(lines[0], ", ") {
		patterns = append(patterns, pattern)
	}

	designs := []string{}
	for _, design := range strings.Split(lines[1], "\n") {
		designs = append(designs, design)
	}

	memo := make(map[string]int)

	var dfs func(design string) int
	dfs = func(design string) int {
		if design == "" {
			return 1
		}

		if val, exists := memo[design]; exists {
			return val
		}

		count := 0
		for _, pattern := range patterns {
			if strings.HasPrefix(design, pattern) {
				count += dfs(design[len(pattern):])
			}
		}

		memo[design] = count
		return count

	}

	count := 0
	for i, design := range designs {
		fmt.Println("Design:", i, "/", len(designs))
		count += dfs(design)
	}

	fmt.Println("Count:", count)
}

func part1() {
	input := utils.GetInput(19)

	lines := strings.Split(input, "\n\n")

	patterns := []string{}
	for _, pattern := range strings.Split(lines[0], ", ") {
		patterns = append(patterns, pattern)
	}

	designs := []string{}
	for _, design := range strings.Split(lines[1], "\n") {
		designs = append(designs, design)
	}

	var dfs func(p []rune) bool
	dfs = func(p []rune) bool {
		if len(p) == 0 {
			return true
		}

	pattern:
		for i := 0; i < len(patterns); i++ {
			pattern := []rune(patterns[i])
			if len(p) < len(pattern) {
				continue
			}

			for j := 0; j < len(pattern); j++ {
				if p[j] != pattern[j] {
					continue pattern
				}
			}
			result := dfs(p[len(pattern):])
			if result {
				return true
			}
		}
		return false
	}

	possible := 0
	for i, design := range designs {
		fmt.Println("Design:", i, "/", len(designs))
		p := []rune(design)
		if dfs(p) {
			possible++
		}
	}

	fmt.Println("Possible:", possible)
}
