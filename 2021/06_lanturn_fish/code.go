package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(6)

	fishes := map[int]int{}
	for _, line := range strings.Split(input, ",") {
		age, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		fishes[age]++
	}

	for i := 0; i < 256; i++ {
		fishes = simulate(fishes)
	}

	total := 0
	for _, count := range fishes {
		total += count
	}

	fmt.Println("Total:", total)

}

func simulate(fishes map[int]int) map[int]int {
	result := map[int]int{}

	for age, count := range fishes {
		newFishes := tick(age)
		for _, newFish := range newFishes {
			result[newFish] += count
		}
	}

	return result
}

func tick(fish int) []int {
	if fish > 0 {
		return []int{fish - 1}
	}
	return []int{6, 8}
}
