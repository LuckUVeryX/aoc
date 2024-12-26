package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(17)

	containers := []int{}
	for _, line := range strings.Split(input, "\n") {
		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		containers = append(containers, v)
	}

	combinations := 0
	results := map[int]int{}
	var dfs func(idx, litres int, path []int)
	dfs = func(idx, litres int, path []int) {
		if litres <= 0 {
			if litres == 0 {
				combinations++
				results[len(path)]++
			}
			return
		}
		if idx == len(containers) {
			return
		}

		npath := append([]int{}, path...)
		npath = append(npath, containers[idx])
		dfs(idx+1, litres-containers[idx], npath)
		dfs(idx+1, litres, append([]int{}, path...))
	}
	dfs(0, 150, []int{})
	fmt.Println("Combinations:", combinations)

	fmt.Println(results)

}
