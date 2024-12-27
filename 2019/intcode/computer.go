package intcode

import (
	"log"
)

type computer struct {
	pc      int
	relBase int
	mem     map[int]int
	in      <-chan int
	out     chan<- int
}

func newComputer(mem []int, in <-chan int, out chan<- int) *computer {
	c := &computer{
		pc:      0,
		relBase: 0,
		mem:     map[int]int{},
		in:      in,
		out:     out,
	}
	for i, v := range mem {
		c.mem[i] = v
	}
	return c

}

func (c *computer) get(addr int, md mode) int {
	v := c.mem[addr]
	switch md {
	case POS:
		return c.mem[v]
	case IMM:
		return v
	case REL:
		return c.mem[v+c.relBase]
	default:
		log.Fatalf("Unknown Mode: %d", md)
	}
	return -1
}

func (c *computer) set(addr, val int, md mode) {
	v := c.mem[addr]
	switch md {
	case POS:
		c.mem[v] = val
	case REL:
		c.mem[v+c.relBase] = val
	default:
		log.Fatalln("Bad mode for set", md)
	}
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
