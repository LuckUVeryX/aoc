package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Instruction struct {
	cycle int
	value int
}

func main() {
	input := utils.GetInput(10)

	instructions := []Instruction{}

	for _, s := range strings.Split(input, "\n") {
		instructions = append(instructions, parse(s))
	}

	cycle := 0
	x := 1
	result := 0
	crt := ""

	tick := func() {
		if cycle%40 >= x-1 && cycle%40 <= x+1 {
			crt += "#"
		} else {
			crt += " "
		}

		if cycle%40 == 39 {
			crt += "\n"
		}

		cycle++
		if (cycle+20)%40 == 0 {
			result += cycle * x
		}
	}

	for _, instruction := range instructions {
		for i := 0; i < instruction.cycle; i++ {
			tick()
		}
		x += instruction.value
	}

	fmt.Println("Result:", result)
	fmt.Println(crt)
}

func parse(s string) Instruction {
	if s == "noop" {
		return Instruction{1, 0}
	}

	v, _ := strconv.Atoi(strings.Split(s, " ")[1])
	return Instruction{2, v}
}
