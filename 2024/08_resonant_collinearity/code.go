package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(8)
	grid := strings.Split(input, "\n")
	rows := len(grid)
	cols := len(grid[0])

	antennaMaps := make(map[byte][][2]int)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '.' {
				antennaMaps[grid[r][c]] = append(antennaMaps[grid[r][c]], [2]int{r, c})
			}
		}
	}

	antinodes := map[[2]int]bool{}
	dfs := func(antennas [][2]int) {}
	dfs = func(antennas [][2]int) {
		for i := 1; i < len(antennas); i++ {
			dr := antennas[0][0] - antennas[i][0]
			dc := antennas[0][1] - antennas[i][1]

			a1r, a1c := antennas[0][0], antennas[0][1]
			for {
				if a1r < 0 || a1r >= rows || a1c < 0 || a1c >= cols {
					break
				}
				antinodes[[2]int{a1r, a1c}] = true
				a1r += dr
				a1c += dc
			}

			a2r, a2c := antennas[i][0], antennas[i][1]

			for {
				if a2r < 0 || a2r >= rows || a2c < 0 || a2c >= cols {
					break
				}
				antinodes[[2]int{a2r, a2c}] = true
				a2r -= dr
				a2c -= dc
			}

		}

		if len(antennas) > 0 {
			dfs(antennas[1:])
		}
	}

	for _, antennas := range antennaMaps {
		dfs(antennas)
	}

	fmt.Println("antinodes:", len(antinodes))
}
