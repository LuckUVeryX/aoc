package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/utils"
)

const basement = -1

func main() {
	input := utils.GetInput(1)
	count := 0
	for _, v := range input {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
	}

	fmt.Println("Count:", count)

	floor := 0
	for i, v := range input {
		if v == '(' {
			floor++
		} else if v == ')' {
			floor--
		}
		if floor == basement {
			fmt.Println("Position:", i+1)
			break
		}
	}

}
