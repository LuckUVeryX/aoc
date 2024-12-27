package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(6)
	lines := strings.Split(input, "\n")

	graph := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, ")")
		graph[parts[0]] = append(graph[parts[0]], parts[1])
		graph[parts[1]] = append(graph[parts[1]], parts[0])
	}

	visited := map[string]bool{}
	var dfsCount func(node string, depth int) int
	dfsCount = func(node string, depth int) int {
		if visited[node] {
			return 0
		}
		visited[node] = true
		total := depth
		for _, child := range graph[node] {
			total += dfsCount(child, depth+1)
		}
		return total
	}

	fmt.Println("Part 1:", dfsCount("COM", 0))

	visited = map[string]bool{}
	q := []string{"YOU"}
	distance := 0

bfs:
	for len(q) > 0 {
		nq := []string{}
		for _, node := range q {
			if node == "SAN" {
				break bfs
			}
			if visited[node] {
				continue
			}
			visited[node] = true
			for _, child := range graph[node] {
				nq = append(nq, child)
			}
		}
		q = nq
		distance++

	}
	fmt.Println("Part 2:", distance-1-1)
}
