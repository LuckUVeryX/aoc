package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/luckuveryx/aoc/utils"
)

func main() {

	hash := func(s string) string {
		hash := md5.Sum([]byte(s))
		return hex.EncodeToString(hash[:])

	}

	input := utils.GetInput(4)
	value := 0
	for {
		s := input + strconv.Itoa(value)
		if hash(s)[:5] == "00000" {
			break
		}
		value++
	}
	fmt.Println("Value:", value)

	value = 0
	for {
		s := input + strconv.Itoa(value)
		if hash(s)[:6] == "000000" {
			break
		}
		value++
	}

	fmt.Println("Value:", value)
}
