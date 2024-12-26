package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Command struct {
	op   string
	args []string
}

func main() {
	input := utils.GetInput(7)

	commands := map[string]Command{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		wire := parts[1]

		parts = strings.Split(parts[0], " ")

		switch len(parts) {
		case 1:
			// x -> y
			commands[wire] = Command{args: []string{parts[0]}}
		case 2:
			// NOT x -> y
			commands[wire] = Command{
				op:   "NOT",
				args: []string{parts[1]},
			}
		default:
			switch parts[1] {
			case "AND":
				commands[wire] = Command{
					op:   "AND",
					args: []string{parts[0], parts[2]},
				}
			case "OR":
				commands[wire] = Command{
					op:   "OR",
					args: []string{parts[0], parts[2]},
				}
			case "LSHIFT":
				commands[wire] = Command{
					op:   "LSHIFT",
					args: []string{parts[0], parts[2]},
				}
			case "RSHIFT":
				commands[wire] = Command{
					op:   "RSHIFT",
					args: []string{parts[0], parts[2]},
				}
			}
		}

	}

	memo := map[string]uint16{}
	var dfs func(curr string) uint16
	dfs = func(curr string) uint16 {
		if v, ok := memo[curr]; ok {
			return v
		}

		node := commands[curr]
		if node.op == "" {
			v, err := strconv.ParseUint(curr, 10, 16)
			if err == nil {
				value := uint16(v)
				memo[curr] = value
				return value
			}
			return dfs(node.args[0])
		}

		var value uint16
		switch node.op {
		case "AND":
			value = dfs(node.args[0]) & dfs(node.args[1])
		case "OR":
			value = dfs(node.args[0]) | dfs(node.args[1])
		case "LSHIFT":
			v, _ := strconv.ParseUint(node.args[1], 10, 16)
			value = dfs(node.args[0]) << v
		case "RSHIFT":
			v, _ := strconv.ParseUint(node.args[1], 10, 16)
			value = dfs(node.args[0]) >> v
		case "NOT":
			value = ^dfs(node.args[0])
		}
		memo[curr] = value
		return value

	}

	fmt.Println("a:", dfs("a"))
}
