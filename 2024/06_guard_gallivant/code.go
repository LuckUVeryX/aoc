// package main

// import (
// 	"github.com/luckuveryx/aoc/utils"

// 	"fmt"
// 	"strings"
// 	"time"
// )

// func main() {
// 	start := time.Now()
// 	input := utils.GetInput(6)
// 	lines := strings.Split(strings.TrimSpace(input), "\n")

// 	dirs := [4][2]int{
// 		{-1, 0},
// 		{0, 1},
// 		{1, 0},
// 		{0, -1},
// 	}
// 	rows, cols := len(lines), len(lines[0])
// 	var startR, startC int
// 	grid := make([][]rune, rows)
// 	for r, line := range lines {
// 		grid[r] = make([]rune, cols)
// 		for c, char := range line {
// 			grid[r][c] = char
// 			if char == '^' {
// 				startR = r
// 				startC = c
// 			}
// 		}
// 	}

// 	visit := make(map[[2]int]bool)
// 	move := func(r, c, dirIdx int) {}
// 	move = func(r, c, dirIdx int) {
// 		if r < 0 || r >= rows || c < 0 || c >= cols {
// 			return
// 		}
// 		visit[[2]int{r, c}] = true

// 		nr, nc := r+dirs[dirIdx][0], c+dirs[dirIdx][1]
// 		if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '#' {
// 			// turn 90
// 			nDirIdx := (dirIdx + 1) % 4
// 			move(r+dirs[nDirIdx][0], c+dirs[nDirIdx][1], nDirIdx)
// 			return
// 		}
// 		move(nr, nc, dirIdx)
// 	}

// 	move(startR, startC, 0)

// 	obstacles := 0
// 	for r := 0; r < rows; r++ {
// 		fmt.Println("Row:", r, "/", rows)
// 		for c := 0; c < cols; c++ {
// 			if r == startR && c == startC {
// 				continue
// 			}
// 			if grid[r][c] == '#' {
// 				continue
// 			}

// 			grid[r][c] = '#'

// 			seen := make(map[[3]int]bool)

// 			search := func(r, c, dirIdx int) {}
// 			search = func(r, c, dirIdx int) {
// 				if r < 0 || r >= rows || c < 0 || c >= cols {
// 					return
// 				}
// 				if _, ok := seen[[3]int{r, c, dirIdx}]; ok {
// 					obstacles++
// 					return
// 				}
// 				seen[[3]int{r, c, dirIdx}] = true

// 				nr, nc := r+dirs[dirIdx][0], c+dirs[dirIdx][1]
// 				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '#' {
// 					// turn 90
// 					nDirIdx := (dirIdx + 1) % 4
// 					search(r+dirs[nDirIdx][0], c+dirs[nDirIdx][1], nDirIdx)
// 					return
// 				}
// 				search(nr, nc, dirIdx)
// 			}

// 			search(startR, startC, 0)
// 			grid[r][c] = '.'
// 		}
// 	}
// 	fmt.Println("Obstacles:", obstacles)
// 	fmt.Println("Visited:", len(visit))

// 	elapsed := time.Since(start)
// 	fmt.Println("Elapsed:", elapsed)
// }

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	start := time.Now()
	input := utils.GetInput(6)
	lines := strings.Split(strings.TrimSpace(input), "\n")

	dirs := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	rows, cols := len(lines), len(lines[0])
	var startR, startC int
	grid := make([][]rune, rows)
	for r, line := range lines {
		grid[r] = make([]rune, cols)
		for c, char := range line {
			grid[r][c] = char
			if char == '^' {
				startR = r
				startC = c
			}
		}
	}

	visited := make(map[[2]int]bool)
	var simulate func(r, c, dirIdx int)

	simulate = func(r, c, dirIdx int) {
		for {
			if r < 0 || r >= rows || c < 0 || c >= cols {
				return
			}
			visited[[2]int{r, c}] = true
			nr, nc := r+dirs[dirIdx][0], c+dirs[dirIdx][1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '#' {
				dirIdx = (dirIdx + 1) % 4
			} else {
				r, c = nr, nc
			}
		}
	}
	simulate(startR, startC, 0)

	obstacles := 0
	simulate = func(r, c, dirIdx int) {
		seen := make(map[[3]int]bool)
		for {
			if r < 0 || r >= rows || c < 0 || c >= cols {
				return
			}
			if _, exists := seen[[3]int{r, c, dirIdx}]; exists {
				obstacles++
				return
			}
			seen[[3]int{r, c, dirIdx}] = true
			nr, nc := r+dirs[dirIdx][0], c+dirs[dirIdx][1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '#' {
				dirIdx = (dirIdx + 1) % 4
			} else {
				r, c = nr, nc
			}
		}
	}
	simulate(startR, startC, 0)

	for r := 0; r < rows; r++ {
		fmt.Println("Rows:", r, "/", rows)
		for c := 0; c < cols; c++ {
			if r == startR && c == startC {
				continue
			}
			if grid[r][c] == '#' {
				continue
			}
			grid[r][c] = '#'

			simulate(startR, startC, 0)
			grid[r][c] = '.'
		}
	}

	fmt.Println("Visited:", len(visited))
	fmt.Println("Obstacles:", obstacles)

	elapsed := time.Since(start)
	fmt.Println("Elapsed:", elapsed)
}
