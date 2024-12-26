package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(1)
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		v, _ := strconv.Atoi(line)

		sum += v/3 - 2
	}

	fmt.Println("Part 1:", sum)

	sum = 0

	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		sum += fuel(v)
	}
	fmt.Println("Part 2:", sum)

}

func fuel(mass int) int {
	fuel := 0
	for mass > 0 {
		mass = mass/3 - 2
		if mass >= 0 {
			fuel += mass
		}
	}

	return fuel

}
