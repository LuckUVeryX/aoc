package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(23)

	graph := map[string][]string{}
	for _, line := range strings.Split(input, "\n") {
		var u, v string
		fmt.Sscanf(line, "%2s-%2s", &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	triangles := map[[3]string]bool{}

	for u, neighbors := range graph {
		for i := 0; i < len(neighbors); i++ {
			for j := i + 1; j < len(neighbors); j++ {
				v := neighbors[i]
				w := neighbors[j]
				isTriangle := false
				for _, n := range graph[v] {
					if n == w {
						isTriangle = true
						break
					}
				}
				if isTriangle {
					key := []string{u, v, w}
					sort.Strings(key)
					triangles[[3]string{key[0], key[1], key[2]}] = true
				}

			}
		}

	}

	count := 0
	for triangle := range triangles {
		for _, computer := range triangle {
			if strings.HasPrefix(computer, "t") {
				count++
				break
			}
		}
	}

	fmt.Println("Part 1:", count)

	cliques := [][]string{}

	var bronKerbosch func(r, p, x []string)
	bronKerbosch = func(r, p, x []string) {
		if len(p) == 0 && len(x) == 0 {
			cliques = append(cliques, append([]string{}, r...))
			return
		}

		for _, v := range p {
			nR := append(r, v)
			nP := intersect(p, graph[v])
			nX := intersect(x, graph[v])

			bronKerbosch(nR, nP, nX)

			p = remove(p, v)
			x = append(x, v)
		}
	}

	vertices := []string{}
	for v := range graph {
		vertices = append(vertices, v)
	}

	bronKerbosch([]string{}, vertices, []string{})

	largest := []string{}
	for _, clique := range cliques {
		if len(clique) > len(largest) {
			largest = clique
		}
	}
	sort.Strings(largest)
	fmt.Println("Part 2:", strings.Join(largest, ","))

}

func intersect(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	for _, v := range slice2 {
		set[v] = true
	}
	var intersection []string
	for _, v := range slice1 {
		if set[v] {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func remove(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
