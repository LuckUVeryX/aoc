package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(13)
	lines := strings.Split(input, "\n")

	graph := map[string]map[string]int{}

	attendeesMap := map[string]bool{}
	for _, line := range lines {
		var u, v, sign string
		var h int

		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &u, &sign, &h, &v)
		v = strings.TrimRight(v, ".")
		if sign == "lose" {
			h *= -1
		}
		attendeesMap[u] = true
		attendeesMap[v] = true
		if _, ok := graph[u]; !ok {
			graph[u] = map[string]int{}
		}
		graph[u][v] = h
	}

	attendees := []string{}
	for k := range attendeesMap {
		attendees = append(attendees, k)
	}

	first := ""

	happiness := 0
	var backtrack func(u string, attendees []string, h int)
	backtrack = func(u string, attendees []string, h int) {
		if len(attendees) == 0 {
			h += graph[u][first] + graph[first][u]
			if h > happiness {
				happiness = h
			}
			return
		}
		for i, v := range attendees {
			nAttendees := append([]string{}, attendees[:i]...)
			nAttendees = append(nAttendees, attendees[i+1:]...)
			backtrack(v, nAttendees, h+graph[u][v]+graph[v][u])
		}
	}

	backtrack(first, attendees, 0)
	fmt.Println("Happiness:", happiness)

}
