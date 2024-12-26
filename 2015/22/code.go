package main

import (
	"fmt"
	"math"
)

type State struct {
	PlayerHP, PlayerMana, BossHP, BossDamage, ManaSpent, ShieldTimer, PoisonTimer, RechargeTimer, PlayerArmor int
}

const (
	MagicMissileCost = 53
	DrainCost        = 73
	ShieldCost       = 113
	PoisonCost       = 173
	RechargeCost     = 229
)

var minManaSpent = math.MaxInt

func simulate(state State, isPlayerTurn bool) {
	if state.ManaSpent >= minManaSpent {
		return
	}

	if isPlayerTurn {
		// Hard mode: Player loses 1 HP at the start of their turn
		state.PlayerHP--
		if state.PlayerHP <= 0 {
			return
		}
	}

	// Apply effects
	if state.ShieldTimer > 0 {
		state.PlayerArmor = 7
		state.ShieldTimer--
	} else {
		state.PlayerArmor = 0
	}

	if state.PoisonTimer > 0 {
		state.BossHP -= 3
		state.PoisonTimer--
	}

	if state.RechargeTimer > 0 {
		state.PlayerMana += 101
		state.RechargeTimer--
	}

	// Check if the boss is dead
	if state.BossHP <= 0 {
		if state.ManaSpent < minManaSpent {
			minManaSpent = state.ManaSpent
		}
		return
	}

	if isPlayerTurn {
		// Player's turn: try each spell
		if state.PlayerMana >= MagicMissileCost {
			nextState := state
			nextState.PlayerMana -= MagicMissileCost
			nextState.ManaSpent += MagicMissileCost
			nextState.BossHP -= 4
			simulate(nextState, false)
		}

		if state.PlayerMana >= DrainCost {
			nextState := state
			nextState.PlayerMana -= DrainCost
			nextState.ManaSpent += DrainCost
			nextState.BossHP -= 2
			nextState.PlayerHP += 2
			simulate(nextState, false)
		}

		if state.PlayerMana >= ShieldCost && state.ShieldTimer == 0 {
			nextState := state
			nextState.PlayerMana -= ShieldCost
			nextState.ManaSpent += ShieldCost
			nextState.ShieldTimer = 6
			simulate(nextState, false)
		}

		if state.PlayerMana >= PoisonCost && state.PoisonTimer == 0 {
			nextState := state
			nextState.PlayerMana -= PoisonCost
			nextState.ManaSpent += PoisonCost
			nextState.PoisonTimer = 6
			simulate(nextState, false)
		}

		if state.PlayerMana >= RechargeCost && state.RechargeTimer == 0 {
			nextState := state
			nextState.PlayerMana -= RechargeCost
			nextState.ManaSpent += RechargeCost
			nextState.RechargeTimer = 5
			simulate(nextState, false)
		}
	} else {
		// Boss's turn: attack
		damage := state.BossDamage - state.PlayerArmor
		if damage < 1 {
			damage = 1
		}
		state.PlayerHP -= damage

		// Check if the player is dead
		if state.PlayerHP > 0 {
			simulate(state, true)
		}
	}
}

func main() {
	initialState := State{
		PlayerHP:      50,
		PlayerMana:    500,
		BossHP:        55,
		BossDamage:    8,
		ManaSpent:     0,
		ShieldTimer:   0,
		PoisonTimer:   0,
		RechargeTimer: 0,
		PlayerArmor:   0,
	}

	simulate(initialState, true)

	fmt.Println("Minimum mana spent to win:", minManaSpent)
}
