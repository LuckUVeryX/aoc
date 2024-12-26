package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Point struct {
	r, c int
}

func main() {
	input := utils.GetInput(3)
	lines := strings.Split(input, "\n")

	paths := []map[Point]int{}

	for _, line := range lines {
		r, c := 0, 0
		steps := 0
		parts := strings.Split(line, ",")
		coord := map[Point]int{}
		for _, part := range parts {
			dist, _ := strconv.Atoi(part[1:])
			for i := 0; i < dist; i++ {
				switch part[0] {
				case 'U':
					r--
				case 'D':
					r++
				case 'L':
					c--
				case 'R':
					c++
				}
				steps++
				if _, ok := coord[Point{r, c}]; !ok {
					coord[Point{r, c}] = steps
				}
			}
		}
		paths = append(paths, coord)
	}

	intersections := []Point{}
	for p := range paths[0] {
		if _, ok := paths[1][p]; ok {
			intersections = append(intersections, p)
		}
	}

	distance := math.MaxInt
	for _, p := range intersections {
		dist := intAbs(p.r) + intAbs(p.c)
		if dist < distance {
			distance = dist
		}
	}
	fmt.Println("Part 1:", distance)

	minSteps := math.MaxInt
	for _, p := range intersections {
		steps := paths[0][p] + paths[1][p]
		if steps < minSteps {
			minSteps = steps
		}
	}
	fmt.Println("Part 2:", minSteps)

}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
