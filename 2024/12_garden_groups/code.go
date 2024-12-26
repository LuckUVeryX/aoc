package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(12)
	lines := strings.Split(input, "\n")
	rows, cols := len(lines), len(lines[0])

	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
	}
	for r, line := range lines {
		for c, v := range strings.Split(line, "") {
			grid[r][c] = v
		}
	}

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	price1, price2 := 0, 0
	visited := make(map[[2]int]bool)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if visited[[2]int{r, c}] {
				continue
			}
			area, perimeter, corners := 0, 0, 0
			var dfs func(r, c int, plant string)

			dfs = func(r, c int, plant string) {
				if r < 0 || r >= rows || c < 0 || c >= cols {
					return
				}
				if grid[r][c] != plant {
					return
				}
				if visited[[2]int{r, c}] {
					return
				}
				visited[[2]int{r, c}] = true
				area++

				// If neighbouring is out of bounds or a different plant, there is a perimeter.
				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						perimeter++
					} else if grid[nr][nc] != plant {
						perimeter++
					}
				}

				// Corners https://www.reddit.com/r/adventofcode/comments/1hcpgcp/comment/m1pxhb8/
				for i := 0; i < len(directions); i++ {
					nr1, nc1 := r+directions[i][0], c+directions[i][1]
					nr2, nc2 := r+directions[(i+1)%len(directions)][0], c+directions[(i+1)%len(directions)][1]
					diff1 := nr1 < 0 || nr1 >= rows || nc1 < 0 || nc1 >= cols || grid[nr1][nc1] != plant
					diff2 := nr2 < 0 || nr2 >= rows || nc2 < 0 || nc2 >= cols || grid[nr2][nc2] != plant
					if diff1 && diff2 {
						corners++ // outer corner
					}

					nr3, nc3 := r+directions[i][0]+directions[(i+1)%len(directions)][0], c+directions[i][1]+directions[(i+1)%len(directions)][1]
					diff3 := nr3 < 0 || nr3 >= rows || nc3 < 0 || nc3 >= cols || grid[nr3][nc3] != plant
					same1 := nr1 >= 0 && nr1 < rows && nc1 >= 0 && nc1 < cols && grid[nr1][nc1] == plant
					same2 := nr2 >= 0 && nr2 < rows && nc2 >= 0 && nc2 < cols && grid[nr2][nc2] == plant
					if same1 && same2 && diff3 {
						corners++ // inner corner
					}
				}

				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					dfs(nr, nc, plant)
				}

			}

			dfs(r, c, grid[r][c])
			price1 += area * perimeter
			price2 += area * corners
		}
	}
	fmt.Println("Price1:", price1)
	fmt.Println("Price2:", price2)
}
