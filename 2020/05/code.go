package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

const rows = 128
const cols = 8

func main() {
	input := utils.GetInput(5)
	lines := strings.Split(input, "\n")

	maxId := 0
	ids := []int{}

	for _, line := range lines {
		pass := []byte(line)
		row := pass[:7]
		col := pass[7:]

		f, b := 0, rows-1
		for _, r := range row {
			m := (f + b) / 2
			switch r {
			case 'F':
				b = m
			case 'B':
				f = m + 1
			}
		}

		l, r := 0, cols-1
		for _, c := range col {
			m := (l + r) / 2
			switch c {
			case 'L':
				r = m
			case 'R':
				l = m + 1
			}

		}
		id := f*8 + l
		ids = append(ids, id)
		if id > maxId {
			maxId = id
		}
	}
	fmt.Println("part1:", maxId)

	slices.Sort(ids)

	for i := 1; i < len(ids)-1; i++ {
		prev, curr := ids[i-1], ids[i]
		if curr-prev > 1 {
			fmt.Println("part2:", curr-1)
			break
		}

	}
}
