package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2()
	part1()
}

func part2() {
	// file, _ := os.Open("example.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var schematic []string

	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	rows := len(schematic)
	cols := len(schematic[0])

	gearRatio := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if schematic[r][c] != '*' {
				continue
			}

			visited := make([][]bool, rows)
			for i := range visited {
				visited[i] = make([]bool, cols)
			}

			expand := func(r, c int) (int, bool) {
				visited[r][c] = true
				var lBuffer, rBuffer []string
				left, right := c-1, c+1
				for left >= 0 && isDigit(schematic[r][left]) {
					visited[r][left] = true
					lBuffer = append(lBuffer, string(schematic[r][left]))
					left--
				}
				for right < cols && isDigit(schematic[r][right]) {
					visited[r][right] = true
					rBuffer = append(rBuffer, string(schematic[r][right]))
					right++
				}

				slices.Reverse(lBuffer)

				buffer := append(lBuffer, string(schematic[r][c]))
				buffer = append(buffer, rBuffer...)
				value, err := strconv.Atoi(strings.Join(buffer, ""))
				if err != nil {
					panic(err)
				}
				return value, true
			}

			var gears []int

			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if !isDigit(schematic[nr][nc]) {
					continue
				}
				if visited[nr][nc] {
					continue
				}
				value, ok := expand(nr, nc)
				if ok {
					gears = append(gears, value)
				}

			}

			if len(gears) == 2 {
				gearRatio += gears[0] * gears[1]
			}

		}

	}

	fmt.Println("Gear Ratios:", gearRatio)

}

func part1() {
	// 	file, _ := os.Open("example.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var schematic []string

	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	rows := len(schematic)
	cols := len(schematic[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	sum := 0

	expand := func(r, c int) {
		isValid := false
		builder := strings.Builder{}

		for c < cols && isDigit(schematic[r][c]) {
			visited[r][c] = true
			builder.WriteByte(schematic[r][c])

			if !isValid {
				for _, dir := range directions {
					nr := r + dir[0]
					nc := c + dir[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}
					if isSymbol(schematic[nr][nc]) {
						isValid = true
						break
					}
				}
			}
			c++
		}

		if isValid {
			value, err := strconv.Atoi(builder.String())
			if err != nil {
				panic(err)
			}
			sum += value
		}

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if visited[r][c] {
				continue
			}
			expand(r, c)
		}
	}

	fmt.Println("Sum:", sum)
}

func isSymbol(s byte) bool {
	return !isDigit(s) && s != '.'
}

func isDigit(s byte) bool {
	return s >= '0' && s <= '9'
}

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}
