package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(15)

	parts := strings.Split(input, "\n\n")

	lines := strings.Split(parts[0], "\n")

	rows, cols := len(lines), len(lines[0])
	grid := make([][]string, rows)
	robotRow, robotCol := -1, -1

	for i := range grid {
		grid[i] = []string{}
	}
	for r := 0; r < rows; r++ {
		line := strings.Split(lines[r], "")
		for _, ch := range line {
			switch ch {
			case ".":
				grid[r] = append(grid[r], []string{".", "."}...)
			case "#":
				grid[r] = append(grid[r], []string{"#", "#"}...)
			case "O":
				grid[r] = append(grid[r], []string{"[", "]"}...)
			case "@":
				grid[r] = append(grid[r], []string{"@", "."}...)
			}
		}
	}

	directions := map[string][2]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}
	instructions := [][2]int{}
	for _, line := range strings.Split(parts[1], "\n") {
		for _, instruction := range strings.Split(line, "") {
			instructions = append(instructions, directions[instruction])
		}
	}

	for _, instruction := range instructions {

		r, c := robotRow+instruction[0], robotCol+instruction[1]

		switch grid[r][c] {
		case ".":
			grid[robotRow][robotCol] = "."
			robotRow, robotCol = r, c
			grid[r][c] = "@"

		case "#":
			continue
		case "O":
			nr, nc := r, c
			for grid[nr][nc] == "O" {
				nr, nc = nr+instruction[0], nc+instruction[1]
			}
			switch grid[nr][nc] {
			case ".":
				grid[robotRow][robotCol] = "."
				robotRow, robotCol = r, c
				grid[r][c] = "@"
				grid[nr][nc] = "O"

			case "#":
				continue
			}
		}
	}

	gps := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == "O" {
				gps += 100*r + c
			}
		}
	}

	fmt.Println("GPS:", gps)
}
