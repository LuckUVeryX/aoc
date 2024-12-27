package intcode

import (
	"strconv"
	"strings"
)

func Run(data []int, in <-chan int) <-chan int {
	nData := make([]int, len(data))
	copy(nData, data)
	out := make(chan int)
	m := newComputer(nData, in, out)
	go m.run()
	return out
}

func ReadProgram(s string) []int {
	parts := strings.Split(s, ",")
	data := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		data[i] = num
	}
	return data
}
