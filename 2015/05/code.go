package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(5)
	nice1 := 0

	for _, line := range strings.Split(input, "\n") {
		if isNice1(line) {
			nice1++
		}
	}
	fmt.Println("Nice1:", nice1)

	nice2 := 0
	for _, line := range strings.Split(input, "\n") {
		if isNice2(line) {
			nice2++
		}
	}
	fmt.Println("Nice2:", nice2)
}

func isNice2(s string) bool {

	disjointPairs := func(s string) bool {
		for i := 0; i < len(s)-1; i++ {
			xy := s[i:i+1] + s[i+1:i+2]
			if strings.Contains(s[i+2:], xy) {
				return true
			}
		}
		return false
	}

	hasXyx := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				return true
			}
		}
		return false
	}

	return disjointPairs(s) && hasXyx(s)

}

func isNice1(s string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	vowelsCount := 0
	for _, ch := range s {
		for _, v := range vowels {
			if ch == v {
				vowelsCount++
			}
		}
	}
	if vowelsCount < 3 {
		return false
	}

	isLetterTwice := false
	runes := []rune(s)
	substrings := []string{"ab", "cd", "pq", "xy"}
	for i := 0; i < len(s)-1; i++ {
		ch1 := runes[i]
		ch2 := runes[i+1]

		if ch1 == ch2 {
			isLetterTwice = true
		}

		double := string(ch1) + string(ch2)

		for _, sub := range substrings {
			if double == sub {
				return false
			}
		}

	}

	return isLetterTwice
}
