package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Node struct {
	pos    Point
	score  int
	dirIdx int
	path   []Point
}

const (
	Wall     = "#"
	Start    = "S"
	End      = "E"
	TurnCost = 1000
	MoveCost = 1
)

func main() {
	input := utils.GetInput(16)
	lines := strings.Split(input, "\n")

	dirs := []Direction{
		{1, 0},  // East
		{0, 1},  // South
		{-1, 0}, // West
		{0, -1}, // North
	}

	var start, end Point
	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = strings.Split(line, "")
		for x, ch := range grid[y] {
			switch ch {
			case Start:
				start = Point{x, y}
			case End:
				end = Point{x, y}
			}
		}
	}

	// dijkstra
	var targetScore int
	q := []Node{{start, 0, 0, []Point{start}}}
	visited := map[string]int{}

	for len(q) > 0 {
		sort.Slice(q, func(i, j int) bool {
			return q[i].score < q[j].score
		})
		node := q[0]
		q = q[1:]
		key := node.pos.key(node.dirIdx)
		if visited[key] > 0 {
			continue
		}
		visited[key] = 1

		if node.pos == end {
			fmt.Println("Score:", node.score)
			targetScore = node.score
			break
		}

		next := node.pos.add(dirs[node.dirIdx])
		if next.x >= 0 && next.x < len(grid) && next.y >= 0 && next.y < len(grid[0]) && grid[next.y][next.x] != Wall {
			q = append(q, Node{next, node.score + MoveCost, node.dirIdx, append(node.path, next)})
		}

		left := Node{node.pos, node.score + TurnCost, (node.dirIdx + 3) % 4, node.path}
		right := Node{node.pos, node.score + TurnCost, (node.dirIdx + 1) % 4, node.path}
		q = append(q, left, right)
	}

	paths := [][]Point{}
	q = []Node{{start, 0, 0, []Point{start}}}
	visited = map[string]int{}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.score > targetScore {
			continue
		}

		key := node.pos.key(node.dirIdx)
		if score, exists := visited[key]; exists && score < node.score {
			continue
		}
		visited[key] = node.score

		if node.pos == end && node.score == targetScore {
			paths = append(paths, node.path)
			continue
		}

		next := node.pos.add(dirs[node.dirIdx])
		if next.x >= 0 && next.x < len(grid) && next.y >= 0 && next.y < len(grid[0]) && grid[next.y][next.x] != Wall {
			path := make([]Point, len(node.path))
			// ! Copy not use the same path
			copy(path, node.path)
			q = append(q, Node{next, node.score + MoveCost, node.dirIdx, append(path, next)})
		}

		left := Node{node.pos, node.score + TurnCost, (node.dirIdx + 3) % 4, node.path}
		right := Node{node.pos, node.score + TurnCost, (node.dirIdx + 1) % 4, node.path}
		q = append(q, left, right)
	}

	tiles := map[Point]bool{}
	for _, path := range paths {
		for _, tile := range path {
			tiles[tile] = true
		}
	}

	fmt.Println("Tiles:", len(tiles))

}

func (p Point) key(dirIdx int) string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, dirIdx)
}

func (p Point) add(dir Direction) Point {
	return Point{p.x + dir.dx, p.y + dir.dy}
}
