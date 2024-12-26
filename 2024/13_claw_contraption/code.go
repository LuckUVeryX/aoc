package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Machine struct {
	ax, ay, bx, by, px, py int
}

func main() {
	input := utils.GetInput(13)
	parts := strings.Split(input, "\n\n")

	machines := []Machine{}
	for _, part := range parts {
		lines := strings.Split(part, "\n")

		var ax, ay, bx, by, px, py int

		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)
		px += 10000000000000
		py += 10000000000000

		machine := Machine{ax, ay, bx, by, px, py}
		machines = append(machines, machine)

	}

	tokens := 0
	for _, machine := range machines {
		tokens += machine.solve()
	}

	fmt.Println("Tokens:", tokens)
}

func (m *Machine) solve() int {
	d := m.ax*m.by - m.ay*m.bx
	da := m.px*m.by - m.py*m.bx
	db := m.py*m.ax - m.px*m.ay

	if da%d == 0 && db%d == 0 {
		a := da / d
		b := db / d

		return a*3 + b
	}
	return 0
}
