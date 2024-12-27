package main

import (
	"fmt"
	"math"

	"github.com/luckuveryx/aoc/utils"
)

const (
	rows = 6
	cols = 25
)

func main() {
	input := utils.GetInput(8)
	pixels := []rune(input)

	size := rows * cols
	numLayers := len(pixels) / size

	layers := [][]rune{}
	for i := 0; i < numLayers; i++ {
		layers = append(layers, pixels[i*size:(i+1)*size])
	}

	counters := []map[rune]int{}
	for _, layer := range layers {
		counter := map[rune]int{}
		for _, pixel := range layer {
			counter[pixel]++
		}
		counters = append(counters, counter)
	}

	idx := -1
	zeros := math.MaxInt
	for i, c := range counters {
		if c['0'] < zeros {
			zeros = c['0']
			idx = i
		}
	}

	result := counters[idx]['1'] * counters[idx]['2']
	fmt.Println("Part 1:", result)

	image := [][][]rune{}
	for layer := 0; layer < numLayers; layer++ {
		grid := [][]rune{}
		for r := 0; r < rows; r++ {
			grid = append(grid, []rune{})
			for c := 0; c < cols; c++ {
				grid[r] = append(grid[r], pixels[layer*size+r*cols+c])
			}
		}
		image = append(image, grid)
	}

	grid := make([][]string, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]string, cols)
		for c := 0; c < cols; c++ {
			pixel := "2"
			for layer := 0; layer < numLayers; layer++ {
				if image[layer][r][c] == '0' {
					pixel = " "
					break
				}
				if image[layer][r][c] == '1' {
					pixel = "█"
					break
				}
			}
			grid[r][c] = pixel
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}

}
