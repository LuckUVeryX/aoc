package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(9)
	data := intcode.ReadProgram(input)

	in := make(chan int, 1)
	in <- 1
	out := intcode.Run(data, in)
	for o := range out {
		fmt.Println("Part 1:", o)
	}

	in <- 2
	out = intcode.Run(data, in)
	for o := range out {
		fmt.Println("Part 2:", o)
	}
}
