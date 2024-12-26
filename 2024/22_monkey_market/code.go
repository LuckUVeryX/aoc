package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(22)

	sum := 0
	priceMap := map[[4]int][]int{}
	for _, line := range strings.Split(input, "\n") {
		sequences := map[[4]int]int{}
		prices := []int{}
		differences := []int{}
		secret, _ := strconv.Atoi(line)
		for i := 0; i < 2000; i++ {
			prices = append(prices, secret%10)
			secret = evolve(secret)
		}
		sum += secret
		for i := 1; i < len(prices); i++ {
			differences = append(differences, prices[i]-prices[i-1])
		}
		for i := 0; i < len(differences)-3; i++ {
			key := [4]int{}
			for j := 0; j < 4; j++ {
				key[j] = differences[i+j]
			}

			if _, ok := sequences[key]; !ok {
				sequences[key] = prices[i+4]
			}

		}

		for key, price := range sequences {
			priceMap[key] = append(priceMap[key], price)
		}

	}

	fmt.Println("Sum:", sum)

	highest := 0
	highsetKey := [4]int{}

	for k, v := range priceMap {
		total := 0
		for _, price := range v {
			total += price
		}
		if total > highest {
			highest = total
			highsetKey = k
		}
	}

	fmt.Println("Highest:", highest)
	fmt.Println("Highest Key:", highsetKey)
}

func evolve(secret int) int {
	secret = mix(secret*64, secret)
	secret = prune(secret)

	secret = mix(secret/32, secret)
	secret = prune(secret)

	secret = mix(secret*2048, secret)
	secret = prune(secret)

	return secret
}

func mix(value, secret int) int {
	return value ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}
