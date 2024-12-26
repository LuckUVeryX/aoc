package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	part1()
	part2()

}

func part1() {
	input := utils.GetInput(5)
	computer := NewIntcodeComputer(input, 1)
	computer.Run()
	fmt.Println("Part 1:", computer.output)
}

func part2() {
	input := utils.GetInput(5)
	computer := NewIntcodeComputer(input, 5)
	computer.Run()
	fmt.Println("Part 2:", computer.output)
}
