package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(9)

	isFile := true
	fileId := 0
	const empty = -1

	blocks := []int{}

	for _, v := range strings.Split(input, "") {
		num, _ := strconv.Atoi(v)
		if isFile {
			for i := 0; i < num; i++ {
				blocks = append(blocks, fileId)
			}
			fileId++
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, empty)
			}
		}
		isFile = !isFile
	}

	// l, r := 0, len(blocks)-1

	// for l < r {
	// 	if blocks[r] == empty {
	// 		r--
	// 		continue
	// 	}

	// 	if blocks[l] != empty {
	// 		l++
	// 		continue
	// 	}

	// 	blocks[l], blocks[r] = blocks[r], empty
	// 	l++
	// 	r--
	// }

	// checksum := 0
	// for i, v := range blocks {
	// 	if v == empty {
	// 		continue
	// 	}
	// 	checksum += i * v
	// }
	// fmt.Println("Checksum:", checksum)

	emptyBlocks := [][2]int{}
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		if block != empty {
			continue
		}
		j := i + 1
		for j < len(blocks) && blocks[j] == empty {
			j++
		}
		emptyBlocks = append(emptyBlocks, [2]int{i, j - i})
		i = j
	}

	r := len(blocks) - 1
	for r >= 0 {
		if blocks[r] == empty {
			r--
			continue
		}
		l := r - 1
		for l >= 0 && blocks[l] == blocks[r] {
			l--
		}
		length := r - l

		for i := 0; i < len(emptyBlocks); i++ {
			emptyBlock := emptyBlocks[i]
			emptyIdx, emptyLen := emptyBlock[0], emptyBlock[1]
			if emptyIdx > r {
				break
			}
			if length > emptyLen {
				continue
			}

			for j := 0; j < length; j++ {
				blocks[emptyIdx+j], blocks[r-j] = blocks[r-j], empty
			}

			if length == emptyLen {
				emptyBlocks = append(emptyBlocks[:i], emptyBlocks[i+1:]...)
			} else {
				emptyBlocks[i][0] += length
				emptyBlocks[i][1] = emptyBlocks[i][1] - length
			}
			break

		}
		r -= length
	}

	checksum := 0
	for i, v := range blocks {
		if v == empty {
			continue
		}
		checksum += i * v
	}
	fmt.Println("Checksum:", checksum)
}
