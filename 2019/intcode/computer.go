package intcode

import (
	"log"
)

type computer struct {
	pc  int
	mem []int
	in  <-chan int
	out chan<- int
}

func newComputer(mem []int, in <-chan int, out chan<- int) *computer {
	return &computer{
		pc:  0,
		mem: mem,
		in:  in,
		out: out,
	}
}

func (c *computer) get(i int, instr instruction) int {
	v := c.mem[c.pc+i]
	mode := instr.modes[i-1]
	switch mode {
	case POS:
		return c.mem[v]
	case IMM:
		return v
	}
	log.Fatalln("Invalid parameter mode", mode)
	return -1
}

func (c *computer) read() int {
	v := <-c.in
	return v
}

func (c *computer) write(v int) {
	c.out <- v
}

func (c *computer) run() {
	for ok := true; ok; {
		instr := parseInstr(c.mem[c.pc])
		h := handlers[instr.op]
		ok = h(c, instr)
	}
	close(c.out)
}
