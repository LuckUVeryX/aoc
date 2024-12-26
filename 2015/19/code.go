package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(19)
	parts := strings.Split(input, "\n\n")

	s := parts[1]
	graph := map[string][]string{}
	for _, line := range strings.Split(parts[0], "\n") {
		var u, v string
		fmt.Sscanf(line, "%s => %s", &u, &v)
		graph[u] = append(graph[u], v)
	}

	distinct := map[string]bool{}
	for i := 0; i < len(s); i++ {
		ch := s[i : i+1]
		if _, ok := graph[ch]; ok {
			for _, v := range graph[ch] {
				res := s[:i] + v + s[i+1:]
				distinct[res] = true
			}
		}

		if i+1 < len(s) {
			ch := s[i : i+2]
			if _, ok := graph[ch]; ok {
				for _, v := range graph[ch] {
					res := s[:i] + v + s[i+2:]
					distinct[res] = true
				}
			}
		}
	}

	fmt.Println("Distincts:", len(distinct))

	reverse := map[string][]string{}
	for _, line := range strings.Split(parts[0], "\n") {
		var u, v string
		fmt.Sscanf(line, "%s => %s", &u, &v)
		reverse[v] = append(reverse[v], u)
	}
	target := "e"
	var bfs func(s string) int
	bfs = func(s string) int {
		if s == target {
			return 0
		}

		for u, v := range reverse {
			idxs := matchingIndexes(s, u)
			for _, idx := range idxs {
				for _, v := range v {
					str := s[:idx] + v + s[idx+len(u):]
					value := bfs(str)
					if value != math.MaxInt {
						return value + 1
					}
				}
			}
		}

		return math.MaxInt
	}
	fmt.Println(bfs(s))

}

func matchingIndexes(s string, pattern string) []int {
	indexes := []int{}

	length := len(pattern)

	for i := 0; i <= len(s)-length; i++ {
		if s[i:i+length] == pattern {
			indexes = append(indexes, i)
		}
	}

	return indexes

}
