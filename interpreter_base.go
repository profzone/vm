package vm

import (
	"fmt"
)

var _ interface {
	Interpreter
} = (*BaseInterpreter)(nil)

type BaseInterpreter struct {
	vm               *VM
	stackElementPool *stackElementPool
	returnData       []byte

	instructionSet [256]step
}

func NewBaseInterpreter(vm *VM, instructionSet [256]step) *BaseInterpreter {
	// 检查指令集是否初始化，否则用默认代替
	if !instructionSet[OP_CODE__STOP].valid {
		instructionSet = newBaseInstructionSet()
	}

	return &BaseInterpreter{
		vm:             vm,
		instructionSet: instructionSet,
	}
}

func (i *BaseInterpreter) Run(code []byte, input []byte) ([]byte, error) {
	if i.stackElementPool == nil {
		i.stackElementPool = stackElementPools.get()
		defer func() {
			stackElementPools.put(i.stackElementPool)
			i.stackElementPool = nil
		}()
	}

	if len(code) == 0 {
		return nil, nil
	}

	var (
		opCode OpCode
		memory = NewMemory()
		stack = newStack()
		pc uint64 = 0
	)

	defer func() {
		i.stackElementPool.put(stack.data...)
	}()

	for {
		if pc < uint64(len(code)) {
			opCode = OpCode(code[pc])
		} else {
			opCode = OP_CODE__STOP
		}

		step := i.instructionSet[opCode]
		if !step.valid {
			return nil, fmt.Errorf("invalid opcode 0x%x", int(opCode))
		}
		if err := step.stackValidation(stack); err != nil {
			return nil, err
		}

		//var memorySize uint64
		//if step.memorySize != nil {
		//	memorySize, overflow := algorithm.BigToUint64(step.memorySize(stack))
		//	if overflow {
		//		return nil, errors.New("memory overflow")
		//	}
		//}
		//
		//if memorySize > 0 {
		//	memory.Resize(memorySize)
		//}

		result, err := step.exec(&pc, code, memory, stack)
		if step.returns {
			i.returnData = result
		}

		switch {
		case err != nil:
			return nil, err
		case step.halt:
			return result, nil
		}
		pc++
	}

	return nil, nil
}
