package main

import (
	"fmt"
	"math/big"
)

const (
	startValue = 20151125
	multiplier = 252533
	modulus    = 33554393
)

func main() {
	targetRow, targetCol := 3010, 3019

	var n *big.Int
	for row := 1; ; row++ {
		r := row
		for col := 1; ; col++ {
			if r == 1 && col == 1 {
				n = big.NewInt(20151125)
			} else {
				n = code(n)
			}

			if r == targetRow && col == targetCol {
				fmt.Println("N:", n)
				return
			}

			if col == row {
				break
			}
			r--
		}
	}
}

func code(a *big.Int) *big.Int {
	v := new(big.Int)
	a = v.Mul(a, big.NewInt(252533))
	v = new(big.Int)
	return v.Mod(a, big.NewInt(33554393))
}
