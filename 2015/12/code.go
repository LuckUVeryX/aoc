package main

import (
	"encoding/json"
	"fmt"

	"github.com/luckuveryx/aoc/utils"
)

func main() {
	input := utils.GetInput(12)

	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}

	total := 0.0
	var count func(data interface{})
	count = func(data interface{}) {
		switch v := data.(type) {
		case []interface{}:
			for _, item := range v {
				count(item)
			}
		case map[string]interface{}:
			for _, value := range v {
				if value == "red" {
					return
				}
			}

			for _, item := range v {
				count(item)
			}
		case float64:
			total += v
		default:
			fmt.Println("Unknown:", v)
		}
	}

	count(data)

	fmt.Println("Total:", total)

}
