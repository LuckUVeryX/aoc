package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(2)
	computer := NewIntCodeComputer(input, map[int]int{1: 12, 2: 2})
	computer.Execute()
	fmt.Println("Part 1:", computer.GetMemory()[0])

	target := 19690720
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			computer := NewIntCodeComputer(input, map[int]int{1: i, 2: j})
			computer.Execute()
			if computer.GetMemory()[0] == target {
				fmt.Println("Part 2:", i, j, 100*i+j)
				return
			}
		}
	}

}
