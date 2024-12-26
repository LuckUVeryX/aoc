package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(10)
	lines := strings.Split(input, "\n")

	rows, cols := len(lines), len(lines[0])

	grid := make([][]int, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]int, cols)
		line := strings.Split(lines[r], "")
		for c := 0; c < cols; c++ {
			height, err := strconv.Atoi(line[c])
			if err != nil {
				height = -1
			}
			grid[r][c] = height
		}
	}

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	score := 0
	var dfs func(r, c, target int, peaks map[[2]int]bool)
	dfs = func(r, c, target int, peaks map[[2]int]bool) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if grid[r][c] != target {
			return
		}

		if target == 9 {
			peaks[[2]int{r, c}] = true
			return
		}

		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			dfs(nr, nc, target+1, peaks)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			peaks := make(map[[2]int]bool)
			dfs(r, c, 0, peaks)

			score += len(peaks)
		}

	}

	fmt.Println("Score:", score)

	rating := 0

	var dfs2 func(r, c, target int) int
	dfs2 = func(r, c, target int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return 0
		}
		if grid[r][c] != target {
			return 0
		}

		if target == 9 {
			return 1
		}

		count := 0
		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]
			count += dfs2(nr, nc, target+1)
		}
		return count
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			rating += dfs2(r, c, 0)
		}

	}

	fmt.Println("Rating:", rating)
}
