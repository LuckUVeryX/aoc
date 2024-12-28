package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

const (
	empty = iota
	wall
	block
	paddle
	ball
)

type Point struct {
	x, y int
}

type Game struct {
	screen    map[Point]int
	score     int
	ballPos   Point
	paddlePos Point
	in        chan<- int
	out       <-chan int
}

func main() {
	part1()
	part2()
}

func part1() {
	input := utils.GetInput(13)
	data := intcode.ReadProgram(input)
	out := intcode.Run(data, nil)

	tiles := map[Point]int{}

loop:
	for {
		select {
		case o, ok := <-out:
			if !ok {
				break loop
			}
			x := o
			y := <-out
			tile := <-out
			tiles[Point{x, y}] = tile
		}
	}
	blockCount := 0
	for _, tile := range tiles {
		if tile == block {
			blockCount++
		}
	}

	fmt.Println("Part 1:", blockCount)
}

func part2() {
	game := newGame()
	game.run()

	fmt.Println("Part 2:", game.score)

}

func newGame() *Game {
	data := intcode.ReadProgram(utils.GetInput(13))
	data[0] = 2
	in := make(chan int)
	out := intcode.Run(data, in)

	return &Game{
		screen: make(map[Point]int),
		in:     in,
		out:    out,
	}
}

func (g *Game) run() {
	for {
		joystick := 0
		if g.paddlePos.x < g.ballPos.x {
			joystick = 1
		}
		if g.paddlePos.x > g.ballPos.x {
			joystick = -1
		}

		select {
		case o, ok := <-g.out:
			if !ok {
				return
			}

			x := o
			y := <-g.out
			tile := <-g.out

			isScore := x == -1 && y == 0 && tile >= 5
			if isScore {
				g.score = tile
				continue
			}

			g.screen[Point{x, y}] = tile
			switch tile {
			case ball:
				g.ballPos = Point{x, y}
			case paddle:
				g.paddlePos = Point{x, y}
			}

		case g.in <- joystick:

		default:
		}
	}
}
