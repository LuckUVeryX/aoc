package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Node struct {
	op       string
	children [2]string
}

func main() {
	input := utils.GetInput(24)

	parts := strings.Split(input, "\n\n")

	keyMap := map[string]bool{}
	values := map[string]int{}
	graph := map[string]Node{}

	for _, line := range strings.Split(parts[0], "\n") {
		split := strings.Split(line, ": ")
		v, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		key := split[0]
		values[key] = v
		keyMap[key] = true
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var u, op, v, w string
		_, err := fmt.Sscanf(line, "%s %s %s -> %s", &u, &op, &v, &w)
		if err != nil {
			panic(err)
		}
		graph[w] = Node{op: op, children: [2]string{u, v}}
		keyMap[w] = true
	}

	memo := map[string]int{}
	var dfs func(curr string) int
	dfs = func(curr string) int {
		if v, ok := values[curr]; ok {
			return v
		}

		if v, ok := memo[curr]; ok {
			return v
		}

		node := graph[curr]

		left := dfs(node.children[0])
		right := dfs(node.children[1])

		var value int
		switch node.op {
		case "AND":
			value = left & right
		case "OR":
			value = left | right
		case "XOR":
			value = left ^ right
		}
		memo[curr] = value
		return value
	}

	keys := make([]string, 0, len(keyMap))
	for k := range keyMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	slices.Reverse(keys)

	bits := []int{}
	for _, k := range keys {
		if strings.HasPrefix(k, "z") {
			bits = append(bits, dfs(k))
		}
	}

	builder := strings.Builder{}
	for _, bit := range bits {
		builder.WriteString(strconv.Itoa(bit))
	}

	value, err := strconv.ParseInt(builder.String(), 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bits:", builder.String())
	fmt.Println("Value:", value)
}
