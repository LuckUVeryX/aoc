package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/utils"
)

type Point struct {
	x, y int
}

func main() {
	part1()
	part2()
}

func part1() {
	visited := map[Point]int{}

	curr := Point{0, 0}
	visited[curr] = 1

	input := utils.GetInput(3)
	for _, ch := range input {
		switch ch {
		case '>':
			curr = Point{curr.x + 1, curr.y}
		case '<':
			curr = Point{curr.x - 1, curr.y}
		case '^':
			curr = Point{curr.x, curr.y - 1}
		case 'v':
			curr = Point{curr.x, curr.y + 1}
		}
		visited[curr]++
	}

	count := 0
	for _, v := range visited {
		if v >= 1 {
			count++
		}
	}

	fmt.Println("Count:", count)
}

func part2() {
	visited := map[Point]int{}

	curr := Point{0, 0}
	visited[curr] = 2

	input := utils.GetInput(3)

	isSanta := true
	santa := Point{0, 0}
	robot := Point{0, 0}

	for _, ch := range input {
		if isSanta {
			switch ch {
			case '>':
				santa = Point{santa.x + 1, santa.y}
			case '<':
				santa = Point{santa.x - 1, santa.y}
			case '^':
				santa = Point{santa.x, santa.y - 1}
			case 'v':
				santa = Point{santa.x, santa.y + 1}
			}
			visited[santa]++
		} else {
			switch ch {
			case '>':
				robot = Point{robot.x + 1, robot.y}
			case '<':
				robot = Point{robot.x - 1, robot.y}
			case '^':
				robot = Point{robot.x, robot.y - 1}
			case 'v':
				robot = Point{robot.x, robot.y + 1}
			}
			visited[robot]++
		}

		isSanta = !isSanta
	}

	count := 0
	for _, v := range visited {
		if v >= 1 {
			count++
		}
	}

	fmt.Println("Count:", count)
}
