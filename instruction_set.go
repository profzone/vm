package vm

type executionFunc func(pc *int64, memory *Memory, stack *Stack) ([]byte, error)
type stackValidationFunc func(stack *Stack) error
type memorySizeFunc func(stack *Stack) error

type step struct {
	exec            executionFunc
	stackValidation stackValidationFunc
	memorySize      memorySizeFunc

	halt  bool
	valid bool
}

func newBaseInstructionSet() [256]step {
	return [256]step{
		OP_CODE__STOP: {
			exec:            opStop,
			stackValidation: makeStackValidationFunc(0, 0),
			halt:            true,
			valid:           true,
		},
		OP_CODE__ADD: {
			exec:            opAdd,
			stackValidation: makeStackValidationFunc(2, 1),
			valid:           true,
		},
	}
}
