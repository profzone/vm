package vm

import "math/big"

type executionFunc func(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error)
type stackValidationFunc func(stack *Stack) error
type memorySizeFunc func(stack *Stack) *big.Int

type step struct {
	exec            executionFunc
	stackValidation stackValidationFunc
	memorySize      memorySizeFunc

	halt    bool
	valid   bool
	returns bool
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
		OP_CODE__PUSH1: {
			exec:            makeOpPush(1),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH2: {
			exec:            makeOpPush(2),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH3: {
			exec:            makeOpPush(3),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH4: {
			exec:            makeOpPush(4),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH5: {
			exec:            makeOpPush(5),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH6: {
			exec:            makeOpPush(6),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH7: {
			exec:            makeOpPush(7),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH8: {
			exec:            makeOpPush(8),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH9: {
			exec:            makeOpPush(9),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH10: {
			exec:            makeOpPush(10),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH11: {
			exec:            makeOpPush(11),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH12: {
			exec:            makeOpPush(12),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH13: {
			exec:            makeOpPush(13),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH14: {
			exec:            makeOpPush(14),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH15: {
			exec:            makeOpPush(15),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH16: {
			exec:            makeOpPush(16),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH17: {
			exec:            makeOpPush(17),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH18: {
			exec:            makeOpPush(18),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH19: {
			exec:            makeOpPush(19),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH20: {
			exec:            makeOpPush(20),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH21: {
			exec:            makeOpPush(21),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH22: {
			exec:            makeOpPush(22),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH23: {
			exec:            makeOpPush(23),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH24: {
			exec:            makeOpPush(24),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH25: {
			exec:            makeOpPush(25),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH26: {
			exec:            makeOpPush(26),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH27: {
			exec:            makeOpPush(27),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH28: {
			exec:            makeOpPush(28),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH29: {
			exec:            makeOpPush(29),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH30: {
			exec:            makeOpPush(30),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH31: {
			exec:            makeOpPush(31),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__PUSH32: {
			exec:            makeOpPush(32),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},

		OP_CODE__DUP1: {
			exec:            makeOpDup(0),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP2: {
			exec:            makeOpDup(1),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP3: {
			exec:            makeOpDup(2),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP4: {
			exec:            makeOpDup(3),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP5: {
			exec:            makeOpDup(4),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP6: {
			exec:            makeOpDup(5),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP7: {
			exec:            makeOpDup(6),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP8: {
			exec:            makeOpDup(7),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP9: {
			exec:            makeOpDup(8),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP10: {
			exec:            makeOpDup(9),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP11: {
			exec:            makeOpDup(10),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP12: {
			exec:            makeOpDup(11),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP13: {
			exec:            makeOpDup(12),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP14: {
			exec:            makeOpDup(13),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP15: {
			exec:            makeOpDup(14),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__DUP16: {
			exec:            makeOpDup(15),
			stackValidation: makeStackValidationFunc(0, 1),
			valid:           true,
		},
		OP_CODE__SWAP2: {
			exec:            makeOpSwap(1),
			stackValidation: makeStackValidationFunc(2, 2),
			valid:           true,
		},
		OP_CODE__SWAP3: {
			exec:            makeOpSwap(2),
			stackValidation: makeStackValidationFunc(3, 3),
			valid:           true,
		},
		OP_CODE__SWAP4: {
			exec:            makeOpSwap(3),
			stackValidation: makeStackValidationFunc(4, 4),
			valid:           true,
		},
		OP_CODE__SWAP5: {
			exec:            makeOpSwap(4),
			stackValidation: makeStackValidationFunc(5, 5),
			valid:           true,
		},
		OP_CODE__SWAP6: {
			exec:            makeOpSwap(5),
			stackValidation: makeStackValidationFunc(6, 6),
			valid:           true,
		},
		OP_CODE__SWAP7: {
			exec:            makeOpSwap(6),
			stackValidation: makeStackValidationFunc(7, 7),
			valid:           true,
		},
		OP_CODE__SWAP8: {
			exec:            makeOpSwap(7),
			stackValidation: makeStackValidationFunc(8, 8),
			valid:           true,
		},
		OP_CODE__SWAP9: {
			exec:            makeOpSwap(8),
			stackValidation: makeStackValidationFunc(9, 9),
			valid:           true,
		},
		OP_CODE__SWAP10: {
			exec:            makeOpSwap(9),
			stackValidation: makeStackValidationFunc(10, 10),
			valid:           true,
		},
		OP_CODE__SWAP11: {
			exec:            makeOpSwap(10),
			stackValidation: makeStackValidationFunc(11, 11),
			valid:           true,
		},
		OP_CODE__SWAP12: {
			exec:            makeOpSwap(11),
			stackValidation: makeStackValidationFunc(12, 12),
			valid:           true,
		},
		OP_CODE__SWAP13: {
			exec:            makeOpSwap(12),
			stackValidation: makeStackValidationFunc(13, 13),
			valid:           true,
		},
		OP_CODE__SWAP14: {
			exec:            makeOpSwap(13),
			stackValidation: makeStackValidationFunc(14, 14),
			valid:           true,
		},
		OP_CODE__SWAP15: {
			exec:            makeOpSwap(14),
			stackValidation: makeStackValidationFunc(15, 15),
			valid:           true,
		},
		OP_CODE__SWAP16: {
			exec:            makeOpSwap(15),
			stackValidation: makeStackValidationFunc(16, 16),
			valid:           true,
		},
	}
}
