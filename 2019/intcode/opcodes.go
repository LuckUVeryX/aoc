package intcode

const (
	add opcode = iota + 1
	mul
	in
	out
	jmpIf
	jmpNot
	lt
	eq
	adjRel
	halt opcode = 99
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
	adjRel: 1,
	halt:   0,
}

type handler func(c *computer, instr instruction) bool

var handlers = map[opcode]handler{
	add: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		c.set(c.pc+3, p1+p2, instr.modes[2])
		c.pc += instr.arity + 1
		return true
	},
	mul: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		c.set(c.pc+3, p1*p2, instr.modes[2])
		c.pc += instr.arity + 1
		return true
	},
	in: func(c *computer, instr instruction) bool {
		c.set(c.pc+1, c.read(), instr.modes[0])
		c.pc += instr.arity + 1
		return true
	},
	out: func(c *computer, instr instruction) bool {
		p1 := c.get(c.pc+1, instr.modes[0])
		c.write(p1)
		c.pc += instr.arity + 1
		return true
	},
	jmpIf: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		if p1 != 0 {
			c.pc = p2
		} else {
			c.pc += instr.arity + 1
		}
		return true
	},
	jmpNot: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		if p1 == 0 {
			c.pc = p2
		} else {
			c.pc += instr.arity + 1
		}
		return true
	},
	lt: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		if p1 < p2 {
			c.set(c.pc+3, 1, instr.modes[2])
		} else {
			c.set(c.pc+3, 0, instr.modes[2])
		}
		c.pc += instr.arity + 1
		return true
	},
	eq: func(c *computer, instr instruction) bool {
		p1, p2 := c.get(c.pc+1, instr.modes[0]), c.get(c.pc+2, instr.modes[1])
		if p1 == p2 {
			c.set(c.pc+3, 1, instr.modes[2])
		} else {
			c.set(c.pc+3, 0, instr.modes[2])
		}
		c.pc += instr.arity + 1
		return true
	},
	adjRel: func(c *computer, instr instruction) bool {
		p1 := c.get(c.pc+1, instr.modes[0])
		c.relBase += p1
		c.pc += instr.arity + 1
		return true
	},
	halt: func(c *computer, instr instruction) bool {
		return false
	},
}
