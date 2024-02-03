package vm

import (
	"interpreter/code"
	"interpreter/object"
)

type Frame struct {
	fn                 *object.CompiledFunction
	instructionPointer int
	basePointer        int
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	return &Frame{fn: fn, instructionPointer: -1, basePointer: basePointer}
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
