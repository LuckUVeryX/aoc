package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/luckuveryx/aoc/2019/intcode"
	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(7)

	phases := []int{0, 1, 2, 3, 4}
	sequences := [][]int{}
	var backtrack func(phases []int, path []int)
	backtrack = func(phases []int, path []int) {
		if len(phases) == 0 {
			sequences = append(sequences, path)
			return
		}

		for i, phase := range phases {
			npath := append([]int{}, path...)
			npath = append(npath, phase)
			nPhases := append([]int{}, phases[:i]...)
			nPhases = append(nPhases, phases[i+1:]...)
			backtrack(nPhases, npath)
		}
	}
	backtrack(phases, []int{})

	data := intcode.ReadProgram(input)
	signal := -math.MaxInt
	for _, sequence := range sequences {
		out := 0
		for _, phase := range sequence {
			in := make(chan int, 2)
			in <- phase
			in <- out
			close(in)
			out = <-intcode.Run(data, in)
		}
		if out > signal {
			signal = out
		}
	}
	fmt.Println("Part 1:", signal)

	signal = -math.MaxInt
	phases = []int{5, 6, 7, 8, 9}
	sequences = [][]int{}
	backtrack(phases, []int{})

	for _, sequence := range sequences {
		channels := make([]chan int, len(sequence))
		for i := range channels {
			channels[i] = make(chan int, 2)
			channels[i] <- sequence[i]
		}

		channels[0] <- 0

		var wg sync.WaitGroup
		for i := 0; i < len(sequence); i++ {
			wg.Add(1)
			go func(i int) {
				nIdx := (i + 1) % len(sequence)
				for output := range intcode.Run(data, channels[i]) {
					channels[nIdx] <- output
					if i == len(sequence)-1 && output > signal {
						signal = output
					}
				}
				wg.Done()
			}(i)
		}

		wg.Wait()

		for _, ch := range channels {
			close(ch)
		}

	}
	fmt.Println("Part 2:", signal)
}
