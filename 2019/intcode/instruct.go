package intcode

type mode int

const (
	POS mode = iota
	IMM
)

type opcode int

type instruction struct {
	op    opcode
	modes []mode
	arity int
}

func parseInstr(i int) instruction {
	var instr instruction
	instr.op = opcode(i % 100)
	instr.arity = opcodeArities[instr.op]
	for i /= 100; len(instr.modes) < instr.arity; i /= 10 {
		instr.modes = append(instr.modes, mode(i%10))
	}

	return instr
}
