package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(25)
	parts := strings.Split(input, "\n\n")

	locks := [][]int{}
	keys := [][]int{}

	for _, part := range parts {
		lines := strings.Split(part, "\n")
		grid := [][]string{}
		for _, line := range lines {
			grid = append(grid, strings.Split(line, ""))
		}
		if strings.Join(grid[0], "") == "#####" {
			lock := []int{}
			for c := 0; c < len(grid[0]); c++ {
				for r := 1; r < len(grid); r++ {
					if grid[r][c] == "." {
						lock = append(lock, r-1)
						break
					}
				}
			}
			locks = append(locks, lock)
		} else {
			key := []int{}
			for c := 0; c < len(grid[0]); c++ {
				for r := len(grid) - 2; r >= 0; r-- {
					if grid[r][c] == "." {
						key = append(key, len(grid)-1-r-1)
						break
					}
				}
			}
			keys = append(keys, key)
		}
	}

	counts := 0
	for _, lock := range locks {
		for _, key := range keys {
			if isValid(lock, key) {
				counts++
			}
		}
	}

	fmt.Println("Count:", counts)
}

func isValid(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}

	return true
}
