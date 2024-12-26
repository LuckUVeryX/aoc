package main

import (
	"log"
	"strconv"
	"strings"
)

const (
	POS = 0
	IMM = 1
)

type IntcodeComputer struct {
	mem    []int
	pc     int
	input  int
	output int
}

func NewIntcodeComputer(s string, input int) *IntcodeComputer {
	mem := parseProgram(s)
	return &IntcodeComputer{
		mem:   mem,
		pc:    0,
		input: input,
	}
}

func parseProgram(s string) []int {
	parts := strings.Split(s, ",")
	mem := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		mem[i] = num
	}
	return mem
}

func (ic *IntcodeComputer) opcode() int {
	return ic.mem[ic.pc] % 100
}

func (ic *IntcodeComputer) mode(pIdx int) int {
	divisor := 10
	for i := 0; i < pIdx; i++ {
		divisor *= 10
	}
	return (ic.mem[ic.pc] / divisor) % 10
}

func (ic *IntcodeComputer) pVal(pIdx int) int {
	mode := ic.mode(pIdx)
	switch mode {
	case POS:
		return ic.mem[ic.mem[ic.pc+pIdx]]
	case IMM:
		return ic.mem[ic.pc+pIdx]
	}
	log.Fatalln("Invalid parameter mode", mode)
	return -1
}

func (ic *IntcodeComputer) Run() {
	for {
		switch ic.opcode() {
		case 1:
			ic.op1()
		case 2:
			ic.op2()
		case 3:
			ic.op3()
		case 4:
			ic.op4()
		case 5:
			ic.op5()
		case 6:
			ic.op6()
		case 7:
			ic.op7()
		case 8:
			ic.op8()
		case 99:
			return
		default:
			log.Fatalln("Unknown opcode:", ic.opcode())
		}
	}
}

func (ic *IntcodeComputer) op1() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	out := ic.mem[ic.pc+3]
	ic.mem[out] = p1 + p2
	ic.pc += 4
}

func (ic *IntcodeComputer) op2() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	out := ic.mem[ic.pc+3]
	ic.mem[out] = p1 * p2
	ic.pc += 4
}

func (ic *IntcodeComputer) op3() {
	out := ic.mem[ic.pc+1]
	ic.mem[out] = ic.input
	ic.pc += 2
}

func (ic *IntcodeComputer) op4() {
	p1 := ic.pVal(1)
	ic.output = p1
	ic.pc += 2
}

func (ic *IntcodeComputer) op5() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	if p1 != 0 {
		ic.pc = p2
	} else {
		ic.pc += 3
	}
}

func (ic *IntcodeComputer) op6() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	if p1 == 0 {
		ic.pc = p2
	} else {
		ic.pc += 3
	}
}

func (ic *IntcodeComputer) op7() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	out := ic.mem[ic.pc+3]
	if p1 < p2 {
		ic.mem[out] = 1
	} else {
		ic.mem[out] = 0
	}
	ic.pc += 4
}

func (ic *IntcodeComputer) op8() {
	p1 := ic.pVal(1)
	p2 := ic.pVal(2)
	out := ic.mem[ic.pc+3]
	if p1 == p2 {
		ic.mem[out] = 1
	} else {
		ic.mem[out] = 0
	}
	ic.pc += 4
}
