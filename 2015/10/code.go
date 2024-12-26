package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Node struct {
	ch    string
	count int
}

func main() {
	input := utils.GetInput(10)

	iter := func(s string) string {
		chars := strings.Split(s, "")
		curr := chars[0]
		counter := 0
		stack := []Node{}

		for _, ch := range chars {
			if ch == curr {
				counter++
				continue
			}
			stack = append(stack, Node{curr, counter})
			curr = ch
			counter = 1
		}
		stack = append(stack, Node{curr, counter})

		builder := strings.Builder{}
		for _, node := range stack {
			builder.WriteString(fmt.Sprintf("%d%s", node.count, node.ch))

		}
		return builder.String()
	}

	out := input
	for i := 0; i < 50; i++ {
		out = iter(out)
	}

	fmt.Println(len(out))

}
