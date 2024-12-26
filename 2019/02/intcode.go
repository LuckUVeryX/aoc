package main

import (
	"strconv"
	"strings"
)

type IntCodeComputer struct {
	memory []int
	addr   int
}

func NewIntCodeComputer(input string, replacements map[int]int) *IntCodeComputer {
	memory := parseInput(input, replacements)
	return &IntCodeComputer{
		memory: memory,
		addr:   0,
	}
}

func parseInput(s string, rMap map[int]int) []int {
	input := strings.Split(s, ",")
	memory := make([]int, len(input))
	for i, v := range input {
		memory[i], _ = strconv.Atoi(v)
	}
	for k, v := range rMap {
		memory[k] = v
	}
	return memory
}

func (ic *IntCodeComputer) GetMemory() []int {
	return ic.memory
}

func (ic *IntCodeComputer) Execute() {
	for {
		opcode := ic.opcode()
		switch opcode {
		case 1:
			ic.op1()
		case 2:
			ic.op2()
		case 99:
			return
		}
	}
}

func (ic *IntCodeComputer) opcode() int {
	return ic.memory[ic.addr]
}

func (ic *IntCodeComputer) adv() {
	ic.addr += 4
}

func (ic *IntCodeComputer) op1() {
	arg1, arg2, out := ic.memory[ic.addr+1], ic.memory[ic.addr+2], ic.memory[ic.addr+3]
	ic.memory[out] = ic.memory[arg1] + ic.memory[arg2]
	ic.adv()
}

func (ic *IntCodeComputer) op2() {
	arg1, arg2, out := ic.memory[ic.addr+1], ic.memory[ic.addr+2], ic.memory[ic.addr+3]
	ic.memory[out] = ic.memory[arg1] * ic.memory[arg2]
	ic.adv()
}
