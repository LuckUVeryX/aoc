package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Coordinate struct {
	X, Y int
}

type Node struct {
	c            Coordinate
	score, steps int
}

func main() {
	input := utils.GetInput(18)
	lines := strings.Split(input, "\n")

	bytes := make([]Coordinate, len(lines))
	for i, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		bytes[i] = Coordinate{x, y}
	}

	createGrid := func(length int) [][]rune {
		grid := make([][]rune, 71)
		for i := 0; i < len(grid); i++ {
			grid[i] = make([]rune, 71)
		}
		for i := 0; i < length; i++ {
			grid[bytes[i].Y][bytes[i].X] = '#'
		}
		return grid

	}

	isReachable := func(grid [][]rune) bool {
		directions := []Coordinate{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}

		q := []Node{{Coordinate{0, 0}, 0, 0}}
		visited := make(map[Coordinate]bool)

		for len(q) > 0 {
			sort.Slice(q, func(i, j int) bool {
				if q[i].score == q[j].score {
					return q[i].steps < q[j].steps
				}
				return q[i].score < q[j].score
			})
			var node Node
			node, q = q[0], q[1:]
			if node.c.X < 0 || node.c.X >= len(grid[0]) || node.c.Y < 0 || node.c.Y >= len(grid) {
				continue
			}
			if grid[node.c.Y][node.c.X] == '#' {
				continue
			}

			if (node.c == Coordinate{70, 70}) {
				fmt.Println("Part 1:", node.steps)
				return true
			}

			if visited[node.c] {
				continue
			}
			visited[node.c] = true

			for _, d := range directions {
				x, y := node.c.X+d.X, node.c.Y+d.Y
				q = append(q, Node{Coordinate{x, y}, node.steps + (70 - x) + (70 - y), node.steps + 1})
			}
		}

		return false
	}

	l, r := 0, len(bytes)-1
	for l < r {
		mid := (l + r) / 2
		fmt.Println("Mid:", mid)
		grid := createGrid(mid)
		if isReachable(grid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Println("Coordinate:", bytes[l-1])
}
