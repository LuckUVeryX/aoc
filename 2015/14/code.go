package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

const (
	totalSeconds = 2503
)

type Reindeer struct {
	speed    int
	duration int
	rest     int
}

func main() {
	input := utils.GetInput(14)
	lines := strings.Split(input, "\n")
	reindeers := []Reindeer{}
	for _, line := range lines {
		var reindeer string
		var speed, duration, rest int
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &reindeer, &speed, &duration, &rest)
		reindeers = append(reindeers, Reindeer{speed, duration, rest})
	}

	maxDistance := 0
	for _, r := range reindeers {
		cycleTime := r.duration + r.rest
		cycles := totalSeconds / cycleTime
		remaining := totalSeconds % cycleTime
		distance := cycles * r.speed * r.duration

		if remaining <= r.duration {
			distance += remaining * r.speed
		} else {
			distance += r.duration * r.speed
		}

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println("Max Distance:", maxDistance)

	distances := []int{}
	points := []int{}
	for i := 0; i < len(reindeers); i++ {
		distances = append(distances, 0)
		points = append(points, 0)
	}

	for t := 0; t < totalSeconds; t++ {
		for i, r := range reindeers {
			cycleTime := r.duration + r.rest
			timeInCycle := t % cycleTime
			if timeInCycle == 0 || timeInCycle < r.duration {
				distances[i] += r.speed
			}
		}

		maxDistance := 0
		for _, d := range distances {
			if d > maxDistance {
				maxDistance = d
			}
		}

		for i, d := range distances {
			if d == maxDistance {
				points[i]++
			}
		}
	}

	maxPoints := 0
	for _, p := range points {
		if p > maxPoints {
			maxPoints = p
		}
	}
	fmt.Println("Max Points:", maxPoints)
}
