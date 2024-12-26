package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(2)

	totalSquareFeet := 0
	for _, line := range strings.Split(input, "\n") {
		totalSquareFeet += squareFeet(line)
	}

	fmt.Println("Square Feet:", totalSquareFeet)

	totalRibbon := 0
	for _, line := range strings.Split(input, "\n") {
		totalRibbon += ribbon(line)
	}

	fmt.Println("Ribbon:", totalRibbon)
}

func squareFeet(s string) int {
	var l, w, h int
	fmt.Sscanf(s, "%dx%dx%d", &l, &w, &h)

	sides := []int{l, w, h}
	sort.Ints(sides)

	area := 2*l*w + 2*w*h + 2*h*l
	area += sides[0] * sides[1]
	return area
}

func ribbon(s string) int {
	var l, w, h int
	fmt.Sscanf(s, "%dx%dx%d", &l, &w, &h)

	sides := []int{l, w, h}
	sort.Ints(sides)

	ribbon := (sides[0] + sides[1]) * 2
	ribbon += l * w * h

	return ribbon
}
