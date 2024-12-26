package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {

	fmt.Println("Complexity:", day21(utils.GetInput(21)))
}

type Coord struct {
	x, y int
}

func day21(input string) int {
	numMap := map[string]Coord{
		"A": {2, 0},
		"0": {1, 0},
		"1": {0, 1},
		"2": {1, 1},
		"3": {2, 1},
		"4": {0, 2},
		"5": {1, 2},
		"6": {2, 2},
		"7": {0, 3},
		"8": {1, 3},
		"9": {2, 3},
	}

	dirMap := map[string]Coord{
		"A": {2, 1},
		"^": {1, 1},
		"<": {0, 0},
		"v": {1, 0},
		">": {2, 0},
	}

	robots := 2
	robots = 25

	return getSequence(strings.Split(input, "\n"), numMap, dirMap, robots)
}

func getSequence(input []string, numMap, dirMap map[string]Coord, robotCount int) int {
	total := 0
	cache := make(map[string][]int)

	for _, line := range input {
		chars := strings.Split(line, "")
		moves := getNumPadSequence(chars, "A", numMap)
		length := countSequences(moves, robotCount, 1, cache, dirMap)
		code, _ := strconv.Atoi(strings.TrimLeft(line[:3], "0"))
		total += length * code
	}

	return total
}

func getNumPadSequence(input []string, start string, numMap map[string]Coord) []string {
	curr := numMap[start]
	seq := []string{}

	for _, char := range input {
		dest := numMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		if dx > 0 {
			for i := 0; i < dx; i++ {
				horiz = append(horiz, ">")
			}
		} else {
			for i := 0; i < -dx; i++ {
				horiz = append(horiz, "<")
			}
		}

		if dy > 0 {
			for i := 0; i < dy; i++ {
				vert = append(vert, "^")
			}
		} else {
			for i := 0; i < -dy; i++ {
				vert = append(vert, "v")
			}
		}

		// Order moves based on position
		if curr.y == 0 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if curr.x == 0 && dest.y == 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func countSequences(input []string, maxRobots, robot int, cache map[string][]int, dirMap map[string]Coord) int {
	key := strings.Join(input, "")
	if val, ok := cache[key]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	if _, ok := cache[key]; !ok {
		cache[key] = make([]int, maxRobots)
	}

	seq := getDirPadSequence(input, "A", dirMap)
	if robot == maxRobots {
		return len(seq)
	}

	steps := splitSequence(seq)
	count := 0
	for _, step := range steps {
		c := countSequences(step, maxRobots, robot+1, cache, dirMap)
		count += c
	}

	cache[key][robot-1] = count
	return count
}

func getDirPadSequence(input []string, start string, dirMap map[string]Coord) []string {
	curr := dirMap[start]
	seq := []string{}

	for _, char := range input {
		dest := dirMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		if dx > 0 {
			for i := 0; i < dx; i++ {
				horiz = append(horiz, ">")
			}
		} else {
			for i := 0; i < -dx; i++ {
				horiz = append(horiz, "<")
			}
		}

		if dy > 0 {
			for i := 0; i < dy; i++ {
				vert = append(vert, "^")
			}
		} else {
			for i := 0; i < -dy; i++ {
				vert = append(vert, "v")
			}
		}

		if curr.x == 0 && dest.y == 1 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if curr.y == 1 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func splitSequence(input []string) [][]string {
	var result [][]string
	var current []string

	for _, char := range input {
		current = append(current, char)
		if char == "A" {
			result = append(result, current)
			current = []string{}
		}
	}
	return result
}
