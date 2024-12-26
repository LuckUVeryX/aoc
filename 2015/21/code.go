package main

import (
	"fmt"
	"slices"
	"sort"
)

// Weapons:    Cost  Damage  Armor
// Dagger        8     4       0
// Shortsword   10     5       0
// Warhammer    25     6       0
// Longsword    40     7       0
// Greataxe     74     8       0

// Armor:      Cost  Damage  Armor
// Leather      13     0       1
// Chainmail    31     0       2
// Splintmail   53     0       3
// Bandedmail   75     0       4
// Platemail   102     0       5

// Rings:      Cost  Damage  Armor
// Damage +1    25     1       0
// Damage +2    50     2       0
// Damage +3   100     3       0
// Defense +1   20     0       1
// Defense +2   40     0       2
// Defense +3   80     0       3

// Hit Points: 103
// Damage: 9
// Armor: 2

type Player struct {
	hp, damage, armor int
}

type Equipment struct {
	cost, damage, armor int
}

func main() {

	weapons := []Equipment{
		{cost: 0, damage: 0, armor: 0},
		{cost: 8, damage: 4, armor: 0},
		{cost: 10, damage: 5, armor: 0},
		{cost: 25, damage: 6, armor: 0},
		{cost: 40, damage: 7, armor: 0},
		{cost: 74, damage: 8, armor: 0},
	}

	armors := []Equipment{
		{cost: 0, damage: 0, armor: 0},
		{cost: 13, damage: 0, armor: 1},
		{cost: 31, damage: 0, armor: 2},
		{cost: 53, damage: 0, armor: 3},
		{cost: 75, damage: 0, armor: 4},
		{cost: 102, damage: 0, armor: 5},
	}

	rings := []Equipment{
		{cost: 0, damage: 0, armor: 0},
		{cost: 25, damage: 1, armor: 0},
		{cost: 50, damage: 2, armor: 0},
		{cost: 100, damage: 3, armor: 0},
		{cost: 20, damage: 0, armor: 1},
		{cost: 40, damage: 0, armor: 2},
		{cost: 80, damage: 0, armor: 3},
	}

	equipsMap := map[int][][]Equipment{}

	var backtrack func(equips []Equipment)
	backtrack = func(equips []Equipment) {
		if len(equips) == 4 {
			if equips[2] == equips[3] {
				return
			}
			cost := equips[0].cost + equips[1].cost + equips[2].cost + equips[3].cost
			equipsMap[cost] = append(equipsMap[cost], equips)
			return
		}

		var consider []Equipment
		switch len(equips) {
		case 0:
			consider = weapons
		case 1:
			consider = armors
		case 2:
			consider = rings
		case 3:
			consider = rings
		}

		for _, equip := range consider {
			nEquips := append([]Equipment{}, equips...)
			nEquips = append(nEquips, equip)
			backtrack(nEquips)
		}
	}

	backtrack([]Equipment{})

	keys := make([]int, 0, len(equipsMap))
	for k := range equipsMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	slices.Reverse(keys)

	for _, cost := range keys {
		for _, equips := range equipsMap[cost] {
			boss := Player{hp: 103, damage: 9, armor: 2}
			me := Player{hp: 100, damage: 0, armor: 0}
			for _, equip := range equips {
				me.damage += equip.damage
				me.armor += equip.armor
			}
			// Simulate the battle
			for me.hp > 0 && boss.hp > 0 {
				dmgToBoss := me.damage - boss.armor
				if dmgToBoss < 0 {
					dmgToBoss = 1
				}
				dmgToMe := boss.damage - me.armor
				if dmgToMe < 0 {
					dmgToMe = 1
				}
				boss.hp -= dmgToBoss
				me.hp -= dmgToMe
			}

			if boss.hp <= 0 {
				fmt.Println("Gold:", cost)
			}

		}

	}

}
