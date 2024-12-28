package main

import (
	"fmt"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

const (
	north Direction = iota + 1
	south
	west
	east
)

const (
	wall Status = iota
	ok
	oxy
)

type Status int
type Direction int

type Point struct {
	x, y int
}

var directions = map[Direction]Point{
	north: {0, -1},
	south: {0, 1},
	west:  {-1, 0},
	east:  {1, 0},
}

type Droid struct {
	in  chan<- int
	out <-chan int
}

func main() {
	m := map[Point]Status{}
	droid := NewDroid()
	droid.visit(Point{0, 0}, m)

	part1(m)
	part2(m)
}

func part1(m map[Point]Status) {
	var target Point
	for pos, status := range m {
		if status == oxy {
			target = pos
		}
	}

	visited := map[Point]bool{}
	distances := map[Point]int{}
	q := []Point{{0, 0}}
	dist := 0
	for len(q) > 0 {
		nq := []Point{}
		for _, node := range q {
			if visited[node] {
				continue
			}
			visited[node] = true
			distances[node] = dist
			for _, dir := range directions {
				next := Point{node.x + dir.x, node.y + dir.y}
				if m[next] == wall {
					continue
				}
				nq = append(nq, next)
			}
		}
		q = nq
		dist++
	}

	fmt.Println("Part 1:", distances[target])
}

func part2(m map[Point]Status) {
	var target Point
	for pos, status := range m {
		if status == oxy {
			target = pos
		}
	}

	visited := map[Point]bool{}
	distances := map[Point]int{}

	q := []Point{target}
	dist := 0
	for len(q) > 0 {
		nq := []Point{}
		for _, node := range q {
			if visited[node] {
				continue
			}
			visited[node] = true
			distances[node] = dist
			for _, dir := range directions {
				next := Point{node.x + dir.x, node.y + dir.y}
				if m[next] == wall {
					continue
				}
				nq = append(nq, next)
			}
		}
		q = nq
		dist++
	}

	longest := 0
	for _, dist := range distances {
		if dist > longest {
			longest = dist
		}
	}
	fmt.Println("Part 2:", longest)

}

func NewDroid() *Droid {
	data := intcode.ReadProgram(utils.GetInput(15))
	in := make(chan int)
	out := intcode.Run(data, in)

	return &Droid{
		in:  in,
		out: out,
	}
}

func (d *Droid) step(dir Direction) int {
	d.in <- int(dir)
	status := <-d.out
	return status
}

func (d *Droid) visit(p Point, m map[Point]Status) {
	for dir, delta := range directions {
		next := Point{p.x + delta.x, p.y + delta.y}
		if _, visited := m[next]; visited {
			continue
		}
		s := Status(d.step(dir))
		m[next] = s
		if s == wall {
			continue
		}
		d.visit(next, m)
		d.step(dir.opposite())
	}
}

func (d Direction) opposite() Direction {
	switch d {
	case north:
		return south
	case south:
		return north
	case west:
		return east
	case east:
		return west
	}
	return 0
}
