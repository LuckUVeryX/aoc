package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

const (
	size = 1000
)

type Instruction struct {
	op           string
	rStart, rEnd int
	cStart, cEnd int
}

func main() {
	input := utils.GetInput(6)
	lines := strings.Split(input, "\n")

	instructions := []Instruction{}
	for _, line := range lines {
		var op string
		var rStart, rEnd, cStart, cEnd int

		if strings.HasPrefix(line, "turn") {
			_, err := fmt.Sscanf(line, "turn %s %d,%d through %d,%d", &op, &rStart, &cStart, &rEnd, &cEnd)
			if err != nil {
				panic(err)
			}
		}

		if strings.HasPrefix(line, "toggle") {
			_, err := fmt.Sscanf(line, "%s %d,%d through %d,%d", &op, &rStart, &cStart, &rEnd, &cEnd)
			if err != nil {
				panic(err)
			}
		}

		instructions = append(instructions, Instruction{
			op:     op,
			rStart: rStart,
			rEnd:   rEnd,
			cStart: cStart,
			cEnd:   cEnd,
		})
	}

	lights := [size][size]int{}
	for _, instruction := range instructions {
		switch instruction.op {
		case "on":
			for r := instruction.rStart; r <= instruction.rEnd; r++ {
				for c := instruction.cStart; c <= instruction.cEnd; c++ {
					lights[r][c]++
				}
			}
		case "off":
			for r := instruction.rStart; r <= instruction.rEnd; r++ {
				for c := instruction.cStart; c <= instruction.cEnd; c++ {
					lights[r][c]--
					if lights[r][c] < 0 {
						lights[r][c] = 0
					}
				}
			}
		case "toggle":
			for r := instruction.rStart; r <= instruction.rEnd; r++ {
				for c := instruction.cStart; c <= instruction.cEnd; c++ {
					lights[r][c] += 2
				}
			}
		}
	}
	count := 0

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			count += lights[r][c]
		}
	}

	fmt.Println("Count:", count)

}
