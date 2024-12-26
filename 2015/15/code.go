package main

import (
	"fmt"
	"strings"

	"github.com/luckuveryx/aoc/utils"
)

type Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	input := utils.GetInput(15)
	ingredients := []Ingredient{}
	for _, line := range strings.Split(input, "\n") {
		var ingredient string
		var capacity, durability, flavor, texture, calories int
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ingredient, &capacity, &durability, &flavor, &texture, &calories)
		ingredient = strings.TrimRight(ingredient, ":")
		ingredients = append(ingredients, Ingredient{
			capacity,
			durability,
			flavor,
			texture,
			calories,
		})
	}

	maxScore := 0
	var backtrack func(idx, teaspoons int, quantities []int)
	backtrack = func(idx, teaspoons int, quantities []int) {
		if idx == len(ingredients) {
			if teaspoons != 0 {
				return
			}

			score, calories := calculateScore(ingredients, quantities)
			if calories == 500 && score > maxScore {
				maxScore = score
			}
			return
		}

		for i := 0; i <= teaspoons; i++ {
			nQuantities := append([]int{}, quantities...)
			nQuantities = append(nQuantities, i)
			backtrack(idx+1, teaspoons-i, nQuantities)
		}
	}

	backtrack(0, 100, []int{})

	fmt.Println("MaxScore:", maxScore)
}

func calculateScore(ingredients []Ingredient, quantities []int) (int, int) {
	var capacity, durability, flavor, texture, calories int

	// Calculate the sum of properties
	for i := 0; i < len(ingredients); i++ {
		quantity := quantities[i]
		capacity += ingredients[i].capacity * quantity
		durability += ingredients[i].durability * quantity
		flavor += ingredients[i].flavor * quantity
		texture += ingredients[i].texture * quantity
		calories += ingredients[i].calories * quantity
	}

	if capacity < 0 {
		capacity = 0
	}
	if durability < 0 {
		durability = 0
	}
	if flavor < 0 {
		flavor = 0
	}
	if texture < 0 {
		texture = 0
	}

	score := capacity * durability * flavor * texture
	return score, calories
}
