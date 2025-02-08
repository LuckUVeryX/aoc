package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(3)
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]byte, rows)
	for r := 0; r < rows; r++ {
		line := lines[r]
		grid[r] = make([]byte, cols)
		for c := 0; c < cols; c++ {
			grid[r][c] = line[c]
		}
	}

	r, c := 0, 0

	count := 0
	for r < rows {
		if grid[r][c] == '#' {
			count++
		}
		r += 1
		c += 3
		if c >= cols {
			c %= cols
		}
	}

	fmt.Println("part1:", count)

	rSlope := []int{1, 1, 1, 1, 2}
	cSlope := []int{1, 3, 5, 7, 1}
	counts := []int{0, 0, 0, 0, 0}

	for i := 0; i < len(counts); i++ {
		r, c = 0, 0
		for r < rows {
			if grid[r][c] == '#' {
				counts[i]++
			}
			r += rSlope[i]
			c += cSlope[i]
			if c >= cols {
				c %= cols
			}
		}
	}

	res := 1
	for _, c := range counts {
		res *= c
	}

	fmt.Println("part2:", res)

}
