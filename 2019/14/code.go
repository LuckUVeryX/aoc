package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

const (
	oreLimit = 1000000000000
)

type Chemical struct {
	qty int
	id  string
}

type Reaction struct {
	in  []Chemical
	out Chemical
}

func main() {
	input := utils.GetInput(14)
	lines := strings.Split(input, "\n")

	reactions := map[string]Reaction{}
	for _, line := range lines {
		reaction := parseReaction(line)
		reactions[reaction.out.id] = reaction
	}

	getOre := func(qty int) int {
		surplus := map[string]int{}
		var dfs func(chem string, qty int) int
		dfs = func(chem string, qty int) int {
			if chem == "ORE" {
				return qty
			}

			if surplus[chem] > 0 {
				used := min(qty, surplus[chem])
				surplus[chem] -= used
				qty -= used
			}

			reaction := reactions[chem]
			batches := (qty + reaction.out.qty - 1) / reaction.out.qty
			surplus[chem] += batches*reaction.out.qty - qty

			totalOre := 0
			for _, input := range reaction.in {
				totalOre += dfs(input.id, input.qty*batches)
			}
			return totalOre
		}
		return dfs("FUEL", qty)
	}

	fmt.Println("Part 1:", getOre(1))

	l := 1
	r := oreLimit
	for l < r {
		m := (l + r) / 2
		ore := getOre(m)
		if ore <= oreLimit {
			l = m + 1
		} else {
			r = m
		}
	}

	fmt.Println("Part 2:", l-1)

}

func parseReaction(s string) Reaction {
	parts := strings.Split(s, " => ")
	out := parseChemical(parts[1])
	in := []Chemical{}
	for _, p := range strings.Split(parts[0], ", ") {
		in = append(in, parseChemical(p))
	}
	return Reaction{in, out}
}

func parseChemical(s string) Chemical {
	var qty int
	var id string
	fmt.Sscanf(s, "%d %s", &qty, &id)
	return Chemical{qty, id}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
