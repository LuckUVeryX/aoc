package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Robot struct {
	r, c   int
	dr, dc int
}

func main() {
	input := utils.GetInput(14)
	lines := strings.Split(input, "\n")

	rows, cols := 103, 101

	robots := []Robot{}

	for _, line := range lines {
		var r, c, dr, dc int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &c, &r, &dc, &dr)

		robot := Robot{r, c, dr, dc}

		robots = append(robots, robot)
	}

	var simulate = func(robot *Robot) {
		robot.r += robot.dr
		robot.c += robot.dc
		if robot.r < 0 {
			robot.r += rows
		}
		if robot.c < 0 {
			robot.c += cols
		}
		if robot.r >= rows {
			robot.r -= rows
		}
		if robot.c >= cols {
			robot.c -= cols
		}
	}

	for i := 0; i < 10000; i++ {
		for j := range robots {
			simulate(&robots[j])
		}
		grid := make([][]int, rows)
		for i := 0; i < rows; i++ {
			grid[i] = make([]int, cols)
		}
		for _, robot := range robots {
			grid[robot.r][robot.c] += 1
		}

		shouldPrint := true
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] > 1 {
					shouldPrint = false
					break
				}
			}
		}
		if !shouldPrint {
			continue
		}

		img := image.NewGray(image.Rect(0, 0, cols, rows))
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == 0 {
					img.SetGray(c, r, color.Gray{Y: 0})
				} else {
					img.SetGray(c, r, color.Gray{Y: 255})
				}
			}
		}
		file, err := os.Create(fmt.Sprintf("output/%d.png", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Encode the image to PNG and save it
		err = png.Encode(file, img)
		if err != nil {
			panic(err)
		}

	}

	quadrants := [4]int{}
	for _, robot := range robots {
		if robot.r == rows/2 || robot.c == cols/2 {
			continue
		}

		if robot.r < rows/2 && robot.c < cols/2 {
			quadrants[0]++
		} else if robot.r < rows/2 && robot.c >= cols/2 {
			quadrants[1]++
		} else if robot.r >= rows/2 && robot.c < cols/2 {
			quadrants[2]++
		} else {
			quadrants[3]++
		}

	}

	fmt.Println("SafetyFactor:", quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3])

}
