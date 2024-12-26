package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(23)
	lines := strings.Split(input, "\n")

	i := 0
	a, b := 1, 0
	for i < len(lines) {
		line := lines[i]
		var v int

		parts := strings.Split(line, ", ")
		// jio or jie
		if len(parts) == 2 {
			v, _ = strconv.Atoi(parts[1])
		}

		parts = strings.Split(parts[0], " ")
		switch parts[0] {
		case "jmp":
			v, _ = strconv.Atoi(parts[1])
			i += v - 1
		case "hlf":
			if parts[1] == "a" {
				a /= 2
			}
			if parts[1] == "b" {
				b /= 2
			}
		case "tpl":
			if parts[1] == "a" {
				a *= 3
			}
			if parts[1] == "b" {
				b *= 3
			}
		case "inc":
			if parts[1] == "a" {
				a++
			}
			if parts[1] == "b" {
				b++
			}
		case "jie":
			if parts[1] == "a" {
				if a%2 == 0 {
					i += v - 1
				}
			}
			if parts[1] == "b" {
				if b%2 == 0 {
					i += v - 1
				}
			}
		case "jio":
			if parts[1] == "a" {
				if a == 1 {
					i += v - 1
				}
			}
			if parts[1] == "b" {
				if b == 1 {
					i += v - 1
				}
			}
		}
		i++
	}

	fmt.Println("A:", a)
	fmt.Println("B:", b)

}
