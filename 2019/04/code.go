package main

import (
	"fmt"
	"strconv"
)

const (
	start = 234208
	end   = 765869
)

func main() {
	count := 0

	for i := start; i <= end; i++ {
		if isValid(i) {
			count++
		}
	}

	fmt.Println("Part 1:", count)

	count2 := 0
	for i := start; i <= end; i++ {
		if isValid2(i) {
			count2++
		}
	}
	fmt.Println("Part 2:", count2)
}

func isValid(num int) bool {
	str := strconv.Itoa(num)
	hasDouble := false

	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
		if str[i] == str[i+1] {
			hasDouble = true
		}
	}

	return hasDouble
}

func isValid2(num int) bool {
	str := strconv.Itoa(num)
	hasExactDouble := false

	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
	}

	for i := 0; i < len(str)-1; i++ {
		if str[i] == str[i+1] &&
			(i-1 < 0 || str[i-1] != str[i+1]) &&
			(i+2 >= len(str) || str[i+2] != str[i]) {
			hasExactDouble = true
			break
		}
	}

	return hasExactDouble

}
