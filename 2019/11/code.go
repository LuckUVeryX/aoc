package main

import (
	"fmt"
	"math"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

const (
	left Rotate = iota
	right
)

type Rotate int

type Phase int

type Robot struct {
	point  Point
	dirIdx int
}

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

func main() {
	input := utils.GetInput(11)
	data := intcode.ReadProgram(input)

	directions := []Direction{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	visited := map[Point]int{}

	robot := Robot{}

	in := make(chan int)
	out := intcode.Run(data, in)
	in <- visited[robot.point]

loop:
	for {
		select {
		case c, ok := <-out:
			if !ok {
				break loop
			}
			visited[robot.point] = c
			rotate := <-out
			switch Rotate(rotate) {
			case left:
				robot.dirIdx = (robot.dirIdx - 1 + len(directions)) % len(directions)
			case right:
				robot.dirIdx = (robot.dirIdx + 1) % len(directions)
			}
			robot.point.x += directions[robot.dirIdx].dx
			robot.point.y += directions[robot.dirIdx].dy
		case in <- visited[robot.point]:
		}
	}

	fmt.Println("Part 1:", len(visited))

	visited = map[Point]int{}
	robot = Robot{}
	in = make(chan int)
	out = intcode.Run(data, in)
	in <- 1
loop2:
	for {
		select {
		case c, ok := <-out:
			if !ok {
				break loop2
			}
			visited[robot.point] = c
			rotate := <-out
			switch Rotate(rotate) {
			case left:
				robot.dirIdx = (robot.dirIdx - 1 + len(directions)) % len(directions)
			case right:
				robot.dirIdx = (robot.dirIdx + 1) % len(directions)
			}
			robot.point.x += directions[robot.dirIdx].dx
			robot.point.y += directions[robot.dirIdx].dy
		case in <- visited[robot.point]:
		}
	}

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for k := range visited {
		x := k.x
		y := k.y
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	for y := minY; y <= maxY; y++ {
		line := []rune{}
		for x := minX; x <= maxX; x++ {
			switch visited[Point{x, y}] {
			case 0:
				line = append(line, ' ')
			case 1:
				line = append(line, '#')
			}
		}
		fmt.Println(string(line))
	}

}
