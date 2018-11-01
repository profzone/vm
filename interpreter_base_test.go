package vm

import "testing"

var interpreter = NewBaseInterpreter(&VM{}, newBaseInstructionSet())
var code = []byte{
	byte(OP_CODE__PUSH1),
	0x02,
	byte(OP_CODE__PUSH1),
	0x04,
	byte(OP_CODE__ADD),
	byte(OP_CODE__STOP),
}

func TestRun(t *testing.T) {
	_, err := interpreter.Run(code, []byte{})
	if err != nil {
		t.Error(err)
	}
}
