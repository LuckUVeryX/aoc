package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(11)
	stones := map[string]int{}

	for _, line := range strings.Split(input, " ") {
		stones[line]++
	}

	for i := 0; i < 75; i++ {
		stones = simulate(stones)
	}

	result := 0
	for _, count := range stones {
		result += count
	}
	fmt.Println("Result:", result)

}

func blink(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}
	if len(stone)%2 == 0 {
		first := strings.TrimLeft(stone[:len(stone)/2], "0")
		second := strings.TrimLeft(stone[len(stone)/2:], "0")
		if first == "" {
			first = "0"
		}
		if second == "" {
			second = "0"
		}
		return []string{first, second}
	}
	value, _ := strconv.Atoi(stone)
	return []string{strconv.Itoa(value * 2024)}
}

func simulate(stones map[string]int) map[string]int {
	result := map[string]int{}

	for stone, count := range stones {
		for _, newStone := range blink(stone) {
			result[newStone] += count
		}
	}

	return result
}
