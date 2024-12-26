package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	target := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	input := utils.GetInput(16)
	lines := strings.Split(input, "\n")
	aunts := []map[string]int{}
	for _, line := range lines {

		re := regexp.MustCompile(`Sue (\d+): (.*)`)
		matches := re.FindStringSubmatch(line)
		parts := strings.Split(matches[2], ", ")
		aunt := map[string]int{}
		for _, part := range parts {
			var k string
			var v int
			fmt.Sscanf(part, "%s %d", &k, &v)
			k = strings.TrimRight(k, ":")
			aunt[k] = v
		}
		aunts = append(aunts, aunt)
	}

	candidates := []int{}
	for i, aunt := range aunts {
		isPossible := true
		for k, v := range target {
			if !isPossible {
				break
			}
			if _, ok := aunt[k]; ok {
				if k == "cats" || k == "trees" {
					if aunt[k] <= v {
						isPossible = false
					}
				} else if k == "pomeranians" || k == "goldfish" {
					if aunt[k] >= v {
						isPossible = false
					}
				} else {
					if aunt[k] != v {
						isPossible = false
					}
				}
			}

		}

		if isPossible {
			candidates = append(candidates, i+1)
		}
	}
	fmt.Println("Candidates:", candidates)
}
