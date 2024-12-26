package main

import (
	"fmt"
	"strconv"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(20)
	target, _ := strconv.Atoi(input)
	fmt.Println("Target:", target)

	house := 1
	for {
		total := 0

		for i := 1; i*i <= house; i++ {
			if house%i == 0 {
				total += i * 10
				if i*i != house {
					total += (house / i) * 10
				}
			}
		}
		if total >= target {
			break
		}
		house++

	}

	fmt.Println("House:", house)

	maxHouse := target / 11
	houses := make([]int, maxHouse+1)

	for elf := 1; elf <= maxHouse; elf++ {
		for house := elf; house <= maxHouse && house <= elf*50; house += elf {
			houses[house] += elf * 11
		}
	}

	// Find the first house with at least the required number of presents
	for house, presents := range houses {
		if presents >= target {
			fmt.Println("House:", house)
			return
		}
	}
}
