package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Point struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type Asteroid struct {
	point    Point
	angle    float64
	distance int
}

func main() {
	input := utils.GetInput(10)
	lines := strings.Split(input, "\n")

	asteroids := []Point{}
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				asteroids = append(asteroids, Point{x, y})
			}
		}
	}

	bestLocation := Point{-1, -1}
	maxVisible := -1
	for _, a := range asteroids {
		directions := map[Direction]bool{}
		for _, b := range asteroids {
			if a == b {
				continue
			}

			dx := a.x - b.x
			dy := a.y - b.y
			divisor := gcd(dx, dy)
			dir := Direction{dx / divisor, dy / divisor}
			directions[dir] = true
		}

		if len(directions) > maxVisible {
			maxVisible = len(directions)
			bestLocation = a
		}
	}

	fmt.Println("Part 1:", maxVisible)

	angleMap := map[float64][]Asteroid{}
	for _, a := range asteroids {
		if a == bestLocation {
			continue
		}
		angle := getAngle(bestLocation, a)
		distance := getDistance(bestLocation, a)
		angleMap[angle] = append(angleMap[angle], Asteroid{a, angle, distance})
	}

	for a := range angleMap {
		sort.Slice(angleMap[a], func(i, j int) bool {
			return angleMap[a][i].distance < angleMap[a][j].distance
		})
	}

	angles := make([]float64, 0, len(angleMap))
	for k := range angleMap {
		angles = append(angles, k)
	}

	sort.Float64s(angles)

	vaporized := []Point{}
	for len(vaporized) < len(asteroids)-1 {
		for _, a := range angles {
			if len(angleMap[a]) > 0 {
				vaporized = append(vaporized, angleMap[a][0].point)
				angleMap[a] = angleMap[a][1:]
			}
		}
	}

	ans := vaporized[200-1]

	fmt.Println("Part 2:", ans.x*100+ans.y)

}

func getAngle(from, to Point) float64 {
	dx := float64(to.x - from.x)
	dy := float64(to.y - from.y)
	angle := math.Atan2(dx, -dy) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}
	return angle
}

func getDistance(from, to Point) int {
	dx := to.x - from.x
	dy := to.y - from.y
	return abs(dx) + abs(dy)
}

func gcd(a, b int) int {
	a = abs(a)
	b = abs(b)

	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
