package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(4)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]string, len(lines))
	for r, line := range lines {
		grid[r] = make([]string, len(line))
		for c, char := range line {
			grid[r][c] = string(char)
		}
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	rows, cols := len(grid), len(grid[0])
	xmas := strings.Split("XMAS", "")
	xmasCount := 0

	searchXmas := func(r, c, dx, dy, idx int) {}
	searchXmas = func(r, c, dx, dy, idx int) {
		if idx == len(xmas) {
			xmasCount++
			return
		}

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}

		if grid[r][c] != xmas[idx] {
			return
		}

		searchXmas(r+dx, c+dy, dx, dy, idx+1)

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == xmas[0] {
				for _, dir := range dirs {
					dx, dy := dir[0], dir[1]
					searchXmas(r+dx, c+dy, dx, dy, 1)
				}
			}

		}
	}

	fmt.Println("Result:", xmasCount)

	diags := [][2][2]int{
		{{1, 1}, {-1, -1}},
		{{-1, 1}, {1, -1}},
	}

	xmasCount2 := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != "A" {
				continue
			}

			isXmas := true

			for _, diag := range diags {
				r1, c1 := r+diag[0][0], c+diag[0][1]
				r2, c2 := r+diag[1][0], c+diag[1][1]

				if r1 < 0 || r1 >= rows || c1 < 0 || c1 >= cols {
					isXmas = false
					break
				}

				if r2 < 0 || r2 >= rows || c2 < 0 || c2 >= cols {
					isXmas = false
					break
				}

				if (grid[r1][c1] == "M" && grid[r2][c2] == "S") || (grid[r1][c1] == "S" && grid[r2][c2] == "M") {
					continue
				}

				isXmas = false
				break

			}

			if isXmas {
				xmasCount2++
			}
		}
	}

	fmt.Println("Result:", xmasCount2)

}
