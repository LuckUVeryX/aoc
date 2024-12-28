package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Moon struct {
	x, y, z Axis
}

type Axis struct {
	pos int
	vel int
}

func main() {
	part1()
	part2()

}

func part1() {
	input := utils.GetInput(12)
	lines := strings.Split(input, "\n")

	moons := []Moon{}
	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, Moon{Axis{x, 0}, Axis{y, 0}, Axis{z, 0}})
	}

	for i := 0; i < 1000; i++ {
		for i := range moons {
			for j := range moons {
				if i == j {
					continue
				}

				if moons[i].x.pos < moons[j].x.pos {
					moons[i].x.vel++
				}
				if moons[i].x.pos > moons[j].x.pos {
					moons[i].x.vel--
				}

				if moons[i].y.pos < moons[j].y.pos {
					moons[i].y.vel++
				}
				if moons[i].y.pos > moons[j].y.pos {
					moons[i].y.vel--
				}

				if moons[i].z.pos < moons[j].z.pos {
					moons[i].z.vel++
				}
				if moons[i].z.pos > moons[j].z.pos {
					moons[i].z.vel--
				}

			}
		}

		for i := range moons {
			moons[i].x.pos += moons[i].x.vel
			moons[i].y.pos += moons[i].y.vel
			moons[i].z.pos += moons[i].z.vel
		}

	}

	energy := 0
	for _, moon := range moons {
		pE := abs(moon.x.pos) + abs(moon.y.pos) + abs(moon.z.pos)
		kE := abs(moon.x.vel) + abs(moon.y.vel) + abs(moon.z.vel)
		energy += pE * kE
	}

	fmt.Println("Part 1:", energy)
}

func part2() {
	input := utils.GetInput(12)
	lines := strings.Split(input, "\n")

	moons := []Moon{}
	for _, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)
		moons = append(moons, Moon{Axis{x, 0}, Axis{y, 0}, Axis{z, 0}})
	}

	xAxes := []Axis{}
	yAxes := []Axis{}
	zAxes := []Axis{}
	for _, moon := range moons {
		xAxes = append(xAxes, moon.x)
		yAxes = append(yAxes, moon.y)
		zAxes = append(zAxes, moon.z)
	}

	xPeriod := getPeriod(xAxes)
	yPeriod := getPeriod(yAxes)
	zPeriod := getPeriod(zAxes)

	result := lcmMultiple([]int{xPeriod, yPeriod, zPeriod})
	fmt.Println("Part 2:", result)

}

func getPeriod(axes []Axis) int {
	seen := map[string]bool{}
	steps := 0

	for {
		keys := keys(axes)
		if seen[keys] {
			return steps
		}
		seen[keys] = true

		for i := range axes {
			for j := range axes {
				if i == j {
					continue
				}

				if axes[i].pos < axes[j].pos {
					axes[i].vel++
				}
				if axes[i].pos > axes[j].pos {
					axes[i].vel--
				}

			}
		}

		for i := range axes {
			axes[i].pos += axes[i].vel
		}
		steps++
	}

}

func keys(axes []Axis) string {
	keys := []string{}
	for _, axis := range axes {
		keys = append(keys, key(axis))
	}
	return strings.Join(keys, ",")
}

func key(axis Axis) string {
	return fmt.Sprintf("%d-%d", axis.pos, axis.vel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmMultiple(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	return lcm(numbers[0], lcmMultiple(numbers[1:]))
}
