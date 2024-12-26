package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(4)
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		numbers := strings.Split(line, ":")[1]
		parts := strings.Split(strings.TrimSpace(numbers), "|")

		winningSet := make(map[string]bool)
		for _, v := range strings.Split(strings.TrimSpace(parts[0]), " ") {
			if v != "" {
				winningSet[v] = true
			}
		}

		points := 0
		for _, v := range strings.Split(strings.TrimSpace(parts[1]), " ") {
			if v == "" {
				continue
			}
			if winningSet[v] {
				if points == 0 {
					points++
				} else {
					points *= 2
				}
			}
		}
		total += points
	}
	fmt.Println("Points:", total)

	cards := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cards[i] = 1
	}

	parseLine := func(line string) {
		parts := strings.Split(line, ":")
		prefix := strings.Split(parts[0], " ")
		id, _ := strconv.Atoi(prefix[len(prefix)-1])
		parts = strings.Split(parts[1], "|")

		winningSet := make(map[string]bool)
		for _, v := range strings.Split(parts[0], " ") {
			if v != "" {
				winningSet[v] = true
			}
		}
		matches := 0
		for _, v := range strings.Split(parts[1], " ") {
			if v != "" && winningSet[v] {
				matches++
			}
		}

		count := cards[id-1]
		for i := 0; i < matches; i++ {
			cards[id+i] += count
		}
	}

	for _, line := range lines {
		parseLine(line)
	}

	count := 0
	for _, card := range cards {
		count += card
	}
	fmt.Println("Count:", count)
}
