package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// file, err := os.Open("example.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	calibrationValue := 0
	for scanner.Scan() {
		text := replaceLine(scanner.Text())
		fmt.Println(text)
		first := getFirst(text)
		second := getLast(text)
		fmt.Println(first, second)
		calibrationValue += first*10 + second
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("CalibrationValue:", calibrationValue)
}

func getFirst(s string) int {
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("No Digit found in " + s)
}

func getLast(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return int(s[i] - '0')
		}
	}
	panic("No digit found in " + s)
}

var replacementMap = map[string]string{
	"one":   "o1ne",
	"two":   "t2wo",
	"three": "t3hree",
	"four":  "f4our",
	"five":  "f5ive",
	"six":   "s6ix",
	"seven": "s7even",
	"eight": "e8ight",
	"nine":  "n9ine",
}

func replaceLine(s string) string {
	for word, value := range replacementMap {
		s = strings.ReplaceAll(s, word, value)
	}
	return s
}
