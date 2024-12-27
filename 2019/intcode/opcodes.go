package intcode

const (
	add    opcode = 1
	mul    opcode = 2
	in     opcode = 3
	out    opcode = 4
	jmpIf  opcode = 5
	jmpNot opcode = 6
	lt     opcode = 7
	eq     opcode = 8
	halt   opcode = 99
)

var opcodeArities = map[opcode]int{
	add:    3,
	mul:    3,
	in:     1,
	out:    1,
	jmpIf:  2,
	jmpNot: 2,
	lt:     3,
	eq:     3,
	halt:   0,
}

type handler func(c *computer, instr instruction) bool

var handlers = map[opcode]handler{
	add: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		out := c.mem[c.pc+3]
		c.mem[out] = p1 + p2
		c.pc += instr.arity + 1
		return true
	},
	mul: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		out := c.mem[c.pc+3]
		c.mem[out] = p1 * p2
		c.pc += instr.arity + 1
		return true
	},
	in: func(c *computer, instr instruction) bool {
		out := c.mem[c.pc+1]
		c.mem[out] = c.read()
		c.pc += instr.arity + 1
		return true
	},
	out: func(c *computer, instr instruction) bool {
		p1 := c.get(1, instr)
		c.write(p1)
		c.pc += instr.arity + 1
		return true
	},
	jmpIf: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		if p1 != 0 {
			c.pc = p2
		} else {
			c.pc += instr.arity + 1
		}
		return true
	},
	jmpNot: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		if p1 == 0 {
			c.pc = p2
		} else {
			c.pc += instr.arity + 1
		}
		return true
	},
	lt: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		out := c.mem[c.pc+3]
		if p1 < p2 {
			c.mem[out] = 1
		} else {
			c.mem[out] = 0
		}
		c.pc += instr.arity + 1
		return true
	},
	eq: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(1, instr), c.get(2, instr)
		out := c.mem[c.pc+3]
		if p1 == p2 {
			c.mem[out] = 1
		} else {
			c.mem[out] = 0
		}
		c.pc += instr.arity + 1
		return true
	},
	halt: func(c *computer, instr instruction) bool {
		return false
	},
}
