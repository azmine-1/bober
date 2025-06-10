package main 

import (
	"fmt"
	"string"
)

type CPU struct {
	A uint8 
	X,Y uint8
	SP uint8 
	PC uint16
	SR uint8
    Memory [65536]uint8
}

const (
    FlagCarry     = 1 << 0
    FlagZero      = 1 << 1
    FlagInterrupt = 1 << 2
    FlagDecimal   = 1 << 3
    FlagBreak     = 1 << 4
    FlagUnused    = 1 << 5
    FlagOverflow  = 1 << 6
    FlagNegative  = 1 << 7
)

type Status struct {
    N,V,B,D,I,Z,C bool
}

func (cpu *CPU) step(){
    opcode := cpu.Memory[cpu.PC]
    switch opcdoe{
    case 0xA9:
        value:= cpu.Memory[cpu.PC+1]
        cpu.A = value 
        cpu.updateZeroAndStatusFlags(cpu.A)
        cpu.PC += 2
    default:
        panic("unkown code")
    }
    
}