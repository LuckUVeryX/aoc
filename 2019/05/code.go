package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

func main() {
	part1()
	part2()

}

func part1() {
	input := utils.GetInput(5)
	data := intcode.ReadProgram(input)
	in := make(chan int, 1)
	out := intcode.Run(data, in)
	in <- 1
	result := <-out
	fmt.Println("Part 1:", result)
}

func part2() {
	input := utils.GetInput(5)
	data := intcode.ReadProgram(input)
	in := make(chan int, 1)
	in <- 5
	out := intcode.Run(data, in)
	result := <-out
	fmt.Println("Part 2:", result)
}
