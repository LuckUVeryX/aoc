package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(17)
	parts := strings.Split(input, "\n\n")

	var a, b, c int
	fmt.Sscanf(parts[0], "Register A: %d\nRegister B: %d\nRegister C: %d", &a, &b, &c)
	program := []int{}
	for _, op := range strings.Split(strings.Split(parts[1], ": ")[1], ",") {
		val, _ := strconv.Atoi(op)
		program = append(program, val)
	}

	out := runProgram(a, b, c, program)
	fmt.Println("Output:", out)

	a = 0
	for pos := len(program) - 1; pos >= 0; pos-- {
		a <<= 3
		for !slices.Equal(runProgram(a, b, c, program), program[pos:]) {
			a++
		}
	}
	fmt.Println("A:", a)
}

func runProgram(a, b, c int, program []int) []int {
	out := []int{}

	for i := 0; i < len(program); i += 2 {
		opcode, operand := program[i], program[i+1]

		combo := operand
		switch operand {
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}

		switch opcode {
		case 0:
			a >>= combo
		case 1:
			b ^= operand
		case 2:
			b = combo % 8
		case 3:
			if a != 0 {
				i = operand - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, combo%8)
		case 6:
			b = a >> combo
		case 7:
			c = a >> combo
		}
	}
	return out
}
