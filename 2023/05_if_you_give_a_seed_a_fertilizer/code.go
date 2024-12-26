package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Range struct {
	from    int
	to      int
	transfo int
}

type Map struct {
	ranges []Range
}

func main() {
	input := utils.GetInput(5)
	parts := strings.Split(input, "\n\n")

	seeds := strings.Split(strings.Split(parts[0], ": ")[1], " ")
	maps := []Map{}
	for _, part := range parts {
		maps = append(maps, parseMap(strings.Split(part, ":")))
	}
	lowest := math.MaxInt

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		for _, map_ := range maps {
			if result, ok := map_.get(seedInt); ok {
				if result < lowest {
					seedInt = result
				}

			}
			lowest = min(lowest, seedInt)
		}
	}
	fmt.Println("Lowest:", lowest)
}

func (m Map) get(seed int) (int, bool) {
	l, r := 0, len(m.ranges)-1
	for l <= r {
		mid := (l + r) / 2
		rng := m.ranges[mid]
		if seed > rng.to {
			l = mid + 1
		} else if seed < rng.from {
			r = mid - 1
		} else {
			return seed + rng.transfo, true
		}
	}

	return 0, false
}

func parseMap(lines []string) Map {
	ranges := []Range{}

	for i := 1; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		ints := make([]int, len(parts))
		for j := 0; j < len(parts); j++ {
			ints[j], _ = strconv.Atoi(parts[j])
		}
		dstRange := ints[0]
		srcRange := ints[1]
		rangeLength := ints[2]

		ranges = append(ranges, Range{
			from:    srcRange,
			to:      srcRange + rangeLength - 1,
			transfo: dstRange,
		})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})

	return Map{ranges}
}
