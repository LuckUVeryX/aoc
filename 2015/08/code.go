package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	encoded := 0
	characters := 0
	memory := 0

	input := utils.GetInput(8)
	for _, line := range strings.Split(input, "\n") {
		length := len(line)
		characters += length
		space := length - 2
		encode := length + 2

		for i := 1; i < len(line)-1; i++ {
			if line[i:i+2] == "\\x" {
				space -= 3
				i += 3
			} else if line[i:i+1] == "\\" {
				space -= 1
				i += 1
			}
		}
		memory += space

		for i := 0; i < len(line); i++ {
			if line[i:i+1] == "\\" {
				encode++
			} else if line[i:i+1] == "\"" {
				encode++
			}
		}
		encoded += encode
	}

	fmt.Println("Characters:", characters)
	fmt.Println("Memory:", memory)
	fmt.Println("Encoded:", encoded)
	fmt.Println(characters - memory)
	fmt.Println(encoded - characters)
}
