package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(9)

	graph := map[string]map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		var u, v string
		var dist int
		fmt.Sscanf(line, "%s to %s = %d", &u, &v, &dist)
		if _, ok := graph[u]; !ok {
			graph[u] = map[string]int{}
		}
		if _, ok := graph[v]; !ok {
			graph[v] = map[string]int{}
		}
		graph[u][v] = dist
		graph[v][u] = dist
	}
	cities := make([]string, 0, len(graph))
	for k := range graph {
		cities = append(cities, k)
	}

	minDist := math.MaxInt
	maxDist := 0
	var dfs func(cities []string, node string, d int)
	dfs = func(cities []string, node string, d int) {
		if len(cities) == 0 {
			if d > maxDist {
				maxDist = d
			}
			if d < minDist {
				minDist = d
			}
			return
		}

		for i := 0; i < len(cities); i++ {
			if dist, ok := graph[node][cities[i]]; ok {
				nCities := append([]string{}, cities[:i]...)
				nCities = append(nCities, cities[i+1:]...)
				dfs(nCities, cities[i], d+dist)

			}
		}
	}

	for i := 0; i < len(cities); i++ {
		nCities := append([]string{}, cities[:i]...)
		nCities = append(nCities, cities[i+1:]...)
		dfs(nCities, cities[i], 0)
	}

	fmt.Println("Min Distance:", minDist)
	fmt.Println("Max Distance:", maxDist)
}
