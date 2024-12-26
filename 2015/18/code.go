package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(18)
	lines := strings.Split(input, "\n")
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]string, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]string, cols)
		line := strings.Split(lines[r], "")
		for c := 0; c < cols; c++ {
			grid[r][c] = line[c]
		}
	}

	corners := [][]int{
		{0, 0},
		{0, cols - 1},
		{rows - 1, 0},
		{rows - 1, cols - 1},
	}

	for _, corner := range corners {
		grid[corner[0]][corner[1]] = "#"
	}

	for i := 0; i < 100; i++ {
		ngrid := clone(grid)
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {

				count := onCount(grid, r, c)

				if grid[r][c] == "#" {
					if count == 2 || count == 3 {
						ngrid[r][c] = "#"
					} else {
						ngrid[r][c] = "."
					}
				}
				if grid[r][c] == "." {
					if count == 3 {
						ngrid[r][c] = "#"
					} else {
						ngrid[r][c] = "."
					}
				}
			}
		}
		grid = ngrid
		for _, corner := range corners {
			grid[corner[0]][corner[1]] = "#"
		}
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == "#" {
				count++
			}
		}
	}
	fmt.Println("Count:", count)
}

func onCount(grid [][]string, r, c int) int {
	dirs := [][]int{
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	count := 0

	for _, dir := range dirs {
		nr := r + dir[0]
		nc := c + dir[1]
		if nr < 0 || nr >= len(grid) {
			continue
		}
		if nc < 0 || nc >= len(grid[0]) {
			continue
		}
		if grid[nr][nc] == "#" {
			count++
		}
	}
	return count
}

func clone(graph [][]string) [][]string {
	clone := make([][]string, len(graph))
	for i := 0; i < len(graph); i++ {
		clone[i] = append([]string{}, graph[i]...)
	}
	return clone
}
