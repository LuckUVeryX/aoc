package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(3)

	result := 0

	mulPattern := []string{"m", "u", "l", "(", ",", ")"}

	mulIdx := 0

	do := true
	first := []string{}
	second := []string{}

	for i := 0; i < len(input); i++ {
		value := string(input[i])
		switch mulIdx {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			if value == "d" {
				// Check do
				if i+4 < len(input) && input[i:i+4] == "do()" {
					do = true
				}
				// Check don't
				if i+7 < len(input) && input[i:i+7] == "don't()" {
					do = false
				}
			}

			if value == mulPattern[mulIdx] {
				mulIdx++
			} else {
				mulIdx = 0
			}
		case 4:
			// Matches ","
			if value == mulPattern[mulIdx] && len(first) > 0 {
				mulIdx++
			} else {
				// Look for integers
				_, err := strconv.Atoi(string(value))
				if err != nil {
					mulIdx = 0
					first = []string{}
					second = []string{}
				} else {
					first = append(first, value)
				}
			}

		case 5:
			// Matches ")"
			if value == mulPattern[mulIdx] && len(second) > 0 {
				// Success
				val1, err := strconv.Atoi(strings.Join(first, ""))
				if err != nil {
					panic(err)
				}
				val2, err := strconv.Atoi(strings.Join(second, ""))
				if err != nil {
					panic(err)
				}
				if do {
					result += val1 * val2
				}

				mulIdx = 0
				first = []string{}
				second = []string{}
			} else {
				// Look for integers
				_, err := strconv.Atoi(value)
				if err != nil {
					mulIdx = 0
					first = []string{}
					second = []string{}
				} else {
					second = append(second, value)
				}

			}

		}

	}

	fmt.Println("Result:", result)
}
