package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const redMax = 12
const greenMax = 13
const blueMax = 14

func main() {
	// file, _ := os.Open("example.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	power := 0
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ": ")
		// * Part 1
		if isValidGame(parts[1]) {
			num, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])
			result += num
		}

		// * Part 2
		cubes := getCubeCount(parts[1])
		power += cubes.red * cubes.blue * cubes.green
	}

	fmt.Println("Sum:", result)
	fmt.Println("Power:", power)
}

type cubes struct {
	red   int
	blue  int
	green int
}

func getCubeCount(game string) cubes {
	c := cubes{}

	sets := strings.Split(game, "; ")

	for _, v := range sets {
		cubes := strings.Split(v, ", ")
		for _, cube := range cubes {
			parts := strings.Split(cube, " ")
			count, _ := strconv.Atoi(parts[0])
			switch parts[1] {
			case "red":
				if count > c.red {
					c.red = count
				}
			case "blue":
				if count > c.blue {
					c.blue = count
				}
			case "green":
				if count > c.green {
					c.green = count
				}
			default:
				panic("Invalid color " + parts[1])
			}
		}
	}
	return c

}

func isValidGame(game string) bool {
	sets := strings.Split(game, "; ")

	for _, v := range sets {
		if !isValidSet(v) {
			return false
		}
	}
	return true
}

func isValidSet(set string) bool {
	cubes := strings.Split(set, ", ")
	for _, cube := range cubes {
		parts := strings.Split(cube, " ")
		count, _ := strconv.Atoi(parts[0])
		switch parts[1] {
		case "red":
			if count > redMax {
				return false
			}
		case "blue":
			if count > blueMax {
				return false
			}
		case "green":
			if count > greenMax {
				return false
			}
		default:
			panic("Invalid color " + parts[1])
		}
	}
	return true
}
