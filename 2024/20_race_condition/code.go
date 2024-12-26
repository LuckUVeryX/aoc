package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Point struct {
	r, c int
}

type Node struct {
	p Point
	n int
}

func main() {
	input := utils.GetInput(20)
	lines := strings.Split(input, "\n")

	rows, cols := len(lines), len(lines[0])
	grid := make([][]rune, rows)
	for r := 0; r < rows; r++ {
		grid[r] = []rune(lines[r])
	}

	var start, end Point
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				start = Point{r, c}
			}
			if grid[r][c] == 'E' {
				end = Point{r, c}
			}
		}
	}

	directions := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	q := []Node{{start, 0}}
	distances := make(map[Point]int)

	for len(q) > 0 {
		var node Node
		node, q = q[0], q[1:]

		if node.p == end {
			break
		}

		for _, direction := range directions {
			p := node.p
			p.r += direction.r
			p.c += direction.c
			if p.r < 0 || p.r >= rows || p.c < 0 || p.c >= cols {
				continue
			}
			if grid[p.r][p.c] == '#' {
				continue
			}
			if distances[p] > 0 {
				continue
			}
			distances[p] = node.n
			q = append(q, Node{p, node.n + 1})
		}

	}
	cheats1 := 0
	cheats2 := 0
	for p, pVal := range distances {
		for q, qVal := range distances {
			d := absInt(p.r-q.r) + absInt(p.c-q.c)
			if d == 2 && pVal-qVal-d >= 100 {
				cheats1++
			}
			if d < 21 && pVal-qVal-d >= 100 {
				cheats2++
			}
		}
	}

	fmt.Println("Cheats1:", cheats1)
	fmt.Println("Cheats2:", cheats2)

}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
