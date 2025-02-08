package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc

func main() {
	input := utils.GetInput(2)
	lines := strings.Split(input, "\n")

	part1 := 0
	part2 := 0
	for _, line := range lines {

		var a, b int

		var c byte
		var s string
		fmt.Sscanf(line, "%d-%d %c: %s", &a, &b, &c, &s)

		// part 1
		count1 := 0
		for _, letter := range s {
			if c == byte(letter) {
				count1++
			}
		}
		if (count1 >= a) && (count1 <= b) {
			part1++
		}

		count2 := 0
		chars := []byte(s)
		if chars[a-1] == c {
			count2++
		}
		if chars[b-1] == c {
			count2++
		}
		if count2 == 1 {
			part2++
		}

	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}
