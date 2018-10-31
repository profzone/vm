package vm

import (
	"bytes"
	"encoding"
	"errors"

	git_chinawayltd_com_golib_tools_courier_enumeration "git.chinawayltd.com/golib/tools/courier/enumeration"
)

var InvalidOpCode = errors.New("invalid OpCode")

func init() {
	git_chinawayltd_com_golib_tools_courier_enumeration.RegisterEnums("OpCode", map[string]string{
		"ADD":    "加",
		"AND":    "按位与运算",
		"DIV":    "除",
		"DUP1":   "复制栈顶第1个元素",
		"DUP10":  "复制栈顶第10个元素",
		"DUP11":  "复制栈顶第11个元素",
		"DUP12":  "复制栈顶第12个元素",
		"DUP13":  "复制栈顶第13个元素",
		"DUP14":  "复制栈顶第14个元素",
		"DUP15":  "复制栈顶第15个元素",
		"DUP16":  "复制栈顶第16个元素",
		"DUP2":   "复制栈顶第2个元素",
		"DUP3":   "复制栈顶第3个元素",
		"DUP4":   "复制栈顶第4个元素",
		"DUP5":   "复制栈顶第5个元素",
		"DUP6":   "复制栈顶第6个元素",
		"DUP7":   "复制栈顶第7个元素",
		"DUP8":   "复制栈顶第8个元素",
		"DUP9":   "复制栈顶第9个元素",
		"EQ":     "等于",
		"EXP":    "幂",
		"GT":     "大于",
		"JUMP":   "跳转",
		"JUMPI":  "有条件跳转",
		"LT":     "小于",
		"MLOAD":  "从内存加载",
		"MOD":    "取模",
		"MSIZE":  "获取活动内存的大小",
		"MSTORE": "存储到内存",
		"MUL":    "乘",
		"NOT":    "按位取反运算",
		"OR":     "按位或运算",
		"PC":     "获取程序计数器",
		"POP":    "出栈",
		"PUSH1":  "入栈1字节",
		"PUSH10": "入栈10字节",
		"PUSH11": "入栈11字节",
		"PUSH12": "入栈12字节",
		"PUSH13": "入栈13字节",
		"PUSH14": "入栈14字节",
		"PUSH15": "入栈15字节",
		"PUSH16": "入栈16字节",
		"PUSH17": "入栈17字节",
		"PUSH18": "入栈18字节",
		"PUSH19": "入栈19字节",
		"PUSH2":  "入栈2字节",
		"PUSH20": "入栈20字节",
		"PUSH21": "入栈21字节",
		"PUSH22": "入栈22字节",
		"PUSH23": "入栈23字节",
		"PUSH24": "入栈24字节",
		"PUSH25": "入栈25字节",
		"PUSH26": "入栈26字节",
		"PUSH27": "入栈27字节",
		"PUSH28": "入栈28字节",
		"PUSH29": "入栈29字节",
		"PUSH3":  "入栈3字节",
		"PUSH30": "入栈30字节",
		"PUSH31": "入栈31字节",
		"PUSH32": "入栈32字节",
		"PUSH4":  "入栈4字节",
		"PUSH5":  "入栈5字节",
		"PUSH6":  "入栈6字节",
		"PUSH7":  "入栈7字节",
		"PUSH8":  "入栈8字节",
		"PUSH9":  "入栈9字节",
		"SLOAD":  "从存储器加载",
		"SSTORE": "存储到存储器",
		"STOP":   "停止",
		"SUB":    "减",
		"SWAP10": "栈顶第10个元素与栈顶元素交换",
		"SWAP11": "栈顶第11个元素与栈顶元素交换",
		"SWAP12": "栈顶第12个元素与栈顶元素交换",
		"SWAP13": "栈顶第13个元素与栈顶元素交换",
		"SWAP14": "栈顶第14个元素与栈顶元素交换",
		"SWAP15": "栈顶第15个元素与栈顶元素交换",
		"SWAP16": "栈顶第16个元素与栈顶元素交换",
		"SWAP2":  "栈顶第2个元素与栈顶元素交换",
		"SWAP3":  "栈顶第3个元素与栈顶元素交换",
		"SWAP4":  "栈顶第4个元素与栈顶元素交换",
		"SWAP5":  "栈顶第5个元素与栈顶元素交换",
		"SWAP6":  "栈顶第6个元素与栈顶元素交换",
		"SWAP7":  "栈顶第7个元素与栈顶元素交换",
		"SWAP8":  "栈顶第8个元素与栈顶元素交换",
		"SWAP9":  "栈顶第9个元素与栈顶元素交换",
		"XOR":    "按位异或运算",
	})
}

func ParseOpCodeFromString(s string) (OpCode, error) {
	switch s {
	case "":
		return OP_CODE_UNKNOWN, nil
	case "ADD":
		return OP_CODE__ADD, nil
	case "AND":
		return OP_CODE__AND, nil
	case "DIV":
		return OP_CODE__DIV, nil
	case "DUP1":
		return OP_CODE__DUP1, nil
	case "DUP10":
		return OP_CODE__DUP10, nil
	case "DUP11":
		return OP_CODE__DUP11, nil
	case "DUP12":
		return OP_CODE__DUP12, nil
	case "DUP13":
		return OP_CODE__DUP13, nil
	case "DUP14":
		return OP_CODE__DUP14, nil
	case "DUP15":
		return OP_CODE__DUP15, nil
	case "DUP16":
		return OP_CODE__DUP16, nil
	case "DUP2":
		return OP_CODE__DUP2, nil
	case "DUP3":
		return OP_CODE__DUP3, nil
	case "DUP4":
		return OP_CODE__DUP4, nil
	case "DUP5":
		return OP_CODE__DUP5, nil
	case "DUP6":
		return OP_CODE__DUP6, nil
	case "DUP7":
		return OP_CODE__DUP7, nil
	case "DUP8":
		return OP_CODE__DUP8, nil
	case "DUP9":
		return OP_CODE__DUP9, nil
	case "EQ":
		return OP_CODE__EQ, nil
	case "EXP":
		return OP_CODE__EXP, nil
	case "GT":
		return OP_CODE__GT, nil
	case "JUMP":
		return OP_CODE__JUMP, nil
	case "JUMPI":
		return OP_CODE__JUMPI, nil
	case "LT":
		return OP_CODE__LT, nil
	case "MLOAD":
		return OP_CODE__MLOAD, nil
	case "MOD":
		return OP_CODE__MOD, nil
	case "MSIZE":
		return OP_CODE__MSIZE, nil
	case "MSTORE":
		return OP_CODE__MSTORE, nil
	case "MUL":
		return OP_CODE__MUL, nil
	case "NOT":
		return OP_CODE__NOT, nil
	case "OR":
		return OP_CODE__OR, nil
	case "PC":
		return OP_CODE__PC, nil
	case "POP":
		return OP_CODE__POP, nil
	case "PUSH1":
		return OP_CODE__PUSH1, nil
	case "PUSH10":
		return OP_CODE__PUSH10, nil
	case "PUSH11":
		return OP_CODE__PUSH11, nil
	case "PUSH12":
		return OP_CODE__PUSH12, nil
	case "PUSH13":
		return OP_CODE__PUSH13, nil
	case "PUSH14":
		return OP_CODE__PUSH14, nil
	case "PUSH15":
		return OP_CODE__PUSH15, nil
	case "PUSH16":
		return OP_CODE__PUSH16, nil
	case "PUSH17":
		return OP_CODE__PUSH17, nil
	case "PUSH18":
		return OP_CODE__PUSH18, nil
	case "PUSH19":
		return OP_CODE__PUSH19, nil
	case "PUSH2":
		return OP_CODE__PUSH2, nil
	case "PUSH20":
		return OP_CODE__PUSH20, nil
	case "PUSH21":
		return OP_CODE__PUSH21, nil
	case "PUSH22":
		return OP_CODE__PUSH22, nil
	case "PUSH23":
		return OP_CODE__PUSH23, nil
	case "PUSH24":
		return OP_CODE__PUSH24, nil
	case "PUSH25":
		return OP_CODE__PUSH25, nil
	case "PUSH26":
		return OP_CODE__PUSH26, nil
	case "PUSH27":
		return OP_CODE__PUSH27, nil
	case "PUSH28":
		return OP_CODE__PUSH28, nil
	case "PUSH29":
		return OP_CODE__PUSH29, nil
	case "PUSH3":
		return OP_CODE__PUSH3, nil
	case "PUSH30":
		return OP_CODE__PUSH30, nil
	case "PUSH31":
		return OP_CODE__PUSH31, nil
	case "PUSH32":
		return OP_CODE__PUSH32, nil
	case "PUSH4":
		return OP_CODE__PUSH4, nil
	case "PUSH5":
		return OP_CODE__PUSH5, nil
	case "PUSH6":
		return OP_CODE__PUSH6, nil
	case "PUSH7":
		return OP_CODE__PUSH7, nil
	case "PUSH8":
		return OP_CODE__PUSH8, nil
	case "PUSH9":
		return OP_CODE__PUSH9, nil
	case "SLOAD":
		return OP_CODE__SLOAD, nil
	case "SSTORE":
		return OP_CODE__SSTORE, nil
	case "STOP":
		return OP_CODE__STOP, nil
	case "SUB":
		return OP_CODE__SUB, nil
	case "SWAP10":
		return OP_CODE__SWAP10, nil
	case "SWAP11":
		return OP_CODE__SWAP11, nil
	case "SWAP12":
		return OP_CODE__SWAP12, nil
	case "SWAP13":
		return OP_CODE__SWAP13, nil
	case "SWAP14":
		return OP_CODE__SWAP14, nil
	case "SWAP15":
		return OP_CODE__SWAP15, nil
	case "SWAP16":
		return OP_CODE__SWAP16, nil
	case "SWAP2":
		return OP_CODE__SWAP2, nil
	case "SWAP3":
		return OP_CODE__SWAP3, nil
	case "SWAP4":
		return OP_CODE__SWAP4, nil
	case "SWAP5":
		return OP_CODE__SWAP5, nil
	case "SWAP6":
		return OP_CODE__SWAP6, nil
	case "SWAP7":
		return OP_CODE__SWAP7, nil
	case "SWAP8":
		return OP_CODE__SWAP8, nil
	case "SWAP9":
		return OP_CODE__SWAP9, nil
	case "XOR":
		return OP_CODE__XOR, nil
	}
	return OP_CODE_UNKNOWN, InvalidOpCode
}

func ParseOpCodeFromLabelString(s string) (OpCode, error) {
	switch s {
	case "":
		return OP_CODE_UNKNOWN, nil
	case "加":
		return OP_CODE__ADD, nil
	case "按位与运算":
		return OP_CODE__AND, nil
	case "除":
		return OP_CODE__DIV, nil
	case "复制栈顶第1个元素":
		return OP_CODE__DUP1, nil
	case "复制栈顶第10个元素":
		return OP_CODE__DUP10, nil
	case "复制栈顶第11个元素":
		return OP_CODE__DUP11, nil
	case "复制栈顶第12个元素":
		return OP_CODE__DUP12, nil
	case "复制栈顶第13个元素":
		return OP_CODE__DUP13, nil
	case "复制栈顶第14个元素":
		return OP_CODE__DUP14, nil
	case "复制栈顶第15个元素":
		return OP_CODE__DUP15, nil
	case "复制栈顶第16个元素":
		return OP_CODE__DUP16, nil
	case "复制栈顶第2个元素":
		return OP_CODE__DUP2, nil
	case "复制栈顶第3个元素":
		return OP_CODE__DUP3, nil
	case "复制栈顶第4个元素":
		return OP_CODE__DUP4, nil
	case "复制栈顶第5个元素":
		return OP_CODE__DUP5, nil
	case "复制栈顶第6个元素":
		return OP_CODE__DUP6, nil
	case "复制栈顶第7个元素":
		return OP_CODE__DUP7, nil
	case "复制栈顶第8个元素":
		return OP_CODE__DUP8, nil
	case "复制栈顶第9个元素":
		return OP_CODE__DUP9, nil
	case "等于":
		return OP_CODE__EQ, nil
	case "幂":
		return OP_CODE__EXP, nil
	case "大于":
		return OP_CODE__GT, nil
	case "跳转":
		return OP_CODE__JUMP, nil
	case "有条件跳转":
		return OP_CODE__JUMPI, nil
	case "小于":
		return OP_CODE__LT, nil
	case "从内存加载":
		return OP_CODE__MLOAD, nil
	case "取模":
		return OP_CODE__MOD, nil
	case "获取活动内存的大小":
		return OP_CODE__MSIZE, nil
	case "存储到内存":
		return OP_CODE__MSTORE, nil
	case "乘":
		return OP_CODE__MUL, nil
	case "按位取反运算":
		return OP_CODE__NOT, nil
	case "按位或运算":
		return OP_CODE__OR, nil
	case "获取程序计数器":
		return OP_CODE__PC, nil
	case "出栈":
		return OP_CODE__POP, nil
	case "入栈1字节":
		return OP_CODE__PUSH1, nil
	case "入栈10字节":
		return OP_CODE__PUSH10, nil
	case "入栈11字节":
		return OP_CODE__PUSH11, nil
	case "入栈12字节":
		return OP_CODE__PUSH12, nil
	case "入栈13字节":
		return OP_CODE__PUSH13, nil
	case "入栈14字节":
		return OP_CODE__PUSH14, nil
	case "入栈15字节":
		return OP_CODE__PUSH15, nil
	case "入栈16字节":
		return OP_CODE__PUSH16, nil
	case "入栈17字节":
		return OP_CODE__PUSH17, nil
	case "入栈18字节":
		return OP_CODE__PUSH18, nil
	case "入栈19字节":
		return OP_CODE__PUSH19, nil
	case "入栈2字节":
		return OP_CODE__PUSH2, nil
	case "入栈20字节":
		return OP_CODE__PUSH20, nil
	case "入栈21字节":
		return OP_CODE__PUSH21, nil
	case "入栈22字节":
		return OP_CODE__PUSH22, nil
	case "入栈23字节":
		return OP_CODE__PUSH23, nil
	case "入栈24字节":
		return OP_CODE__PUSH24, nil
	case "入栈25字节":
		return OP_CODE__PUSH25, nil
	case "入栈26字节":
		return OP_CODE__PUSH26, nil
	case "入栈27字节":
		return OP_CODE__PUSH27, nil
	case "入栈28字节":
		return OP_CODE__PUSH28, nil
	case "入栈29字节":
		return OP_CODE__PUSH29, nil
	case "入栈3字节":
		return OP_CODE__PUSH3, nil
	case "入栈30字节":
		return OP_CODE__PUSH30, nil
	case "入栈31字节":
		return OP_CODE__PUSH31, nil
	case "入栈32字节":
		return OP_CODE__PUSH32, nil
	case "入栈4字节":
		return OP_CODE__PUSH4, nil
	case "入栈5字节":
		return OP_CODE__PUSH5, nil
	case "入栈6字节":
		return OP_CODE__PUSH6, nil
	case "入栈7字节":
		return OP_CODE__PUSH7, nil
	case "入栈8字节":
		return OP_CODE__PUSH8, nil
	case "入栈9字节":
		return OP_CODE__PUSH9, nil
	case "从存储器加载":
		return OP_CODE__SLOAD, nil
	case "存储到存储器":
		return OP_CODE__SSTORE, nil
	case "停止":
		return OP_CODE__STOP, nil
	case "减":
		return OP_CODE__SUB, nil
	case "栈顶第10个元素与栈顶元素交换":
		return OP_CODE__SWAP10, nil
	case "栈顶第11个元素与栈顶元素交换":
		return OP_CODE__SWAP11, nil
	case "栈顶第12个元素与栈顶元素交换":
		return OP_CODE__SWAP12, nil
	case "栈顶第13个元素与栈顶元素交换":
		return OP_CODE__SWAP13, nil
	case "栈顶第14个元素与栈顶元素交换":
		return OP_CODE__SWAP14, nil
	case "栈顶第15个元素与栈顶元素交换":
		return OP_CODE__SWAP15, nil
	case "栈顶第16个元素与栈顶元素交换":
		return OP_CODE__SWAP16, nil
	case "栈顶第2个元素与栈顶元素交换":
		return OP_CODE__SWAP2, nil
	case "栈顶第3个元素与栈顶元素交换":
		return OP_CODE__SWAP3, nil
	case "栈顶第4个元素与栈顶元素交换":
		return OP_CODE__SWAP4, nil
	case "栈顶第5个元素与栈顶元素交换":
		return OP_CODE__SWAP5, nil
	case "栈顶第6个元素与栈顶元素交换":
		return OP_CODE__SWAP6, nil
	case "栈顶第7个元素与栈顶元素交换":
		return OP_CODE__SWAP7, nil
	case "栈顶第8个元素与栈顶元素交换":
		return OP_CODE__SWAP8, nil
	case "栈顶第9个元素与栈顶元素交换":
		return OP_CODE__SWAP9, nil
	case "按位异或运算":
		return OP_CODE__XOR, nil
	}
	return OP_CODE_UNKNOWN, InvalidOpCode
}

func (OpCode) EnumType() string {
	return "OpCode"
}

func (OpCode) Enums() map[int][]string {
	return map[int][]string{
		int(OP_CODE__ADD):    {"ADD", "加"},
		int(OP_CODE__AND):    {"AND", "按位与运算"},
		int(OP_CODE__DIV):    {"DIV", "除"},
		int(OP_CODE__DUP1):   {"DUP1", "复制栈顶第1个元素"},
		int(OP_CODE__DUP10):  {"DUP10", "复制栈顶第10个元素"},
		int(OP_CODE__DUP11):  {"DUP11", "复制栈顶第11个元素"},
		int(OP_CODE__DUP12):  {"DUP12", "复制栈顶第12个元素"},
		int(OP_CODE__DUP13):  {"DUP13", "复制栈顶第13个元素"},
		int(OP_CODE__DUP14):  {"DUP14", "复制栈顶第14个元素"},
		int(OP_CODE__DUP15):  {"DUP15", "复制栈顶第15个元素"},
		int(OP_CODE__DUP16):  {"DUP16", "复制栈顶第16个元素"},
		int(OP_CODE__DUP2):   {"DUP2", "复制栈顶第2个元素"},
		int(OP_CODE__DUP3):   {"DUP3", "复制栈顶第3个元素"},
		int(OP_CODE__DUP4):   {"DUP4", "复制栈顶第4个元素"},
		int(OP_CODE__DUP5):   {"DUP5", "复制栈顶第5个元素"},
		int(OP_CODE__DUP6):   {"DUP6", "复制栈顶第6个元素"},
		int(OP_CODE__DUP7):   {"DUP7", "复制栈顶第7个元素"},
		int(OP_CODE__DUP8):   {"DUP8", "复制栈顶第8个元素"},
		int(OP_CODE__DUP9):   {"DUP9", "复制栈顶第9个元素"},
		int(OP_CODE__EQ):     {"EQ", "等于"},
		int(OP_CODE__EXP):    {"EXP", "幂"},
		int(OP_CODE__GT):     {"GT", "大于"},
		int(OP_CODE__JUMP):   {"JUMP", "跳转"},
		int(OP_CODE__JUMPI):  {"JUMPI", "有条件跳转"},
		int(OP_CODE__LT):     {"LT", "小于"},
		int(OP_CODE__MLOAD):  {"MLOAD", "从内存加载"},
		int(OP_CODE__MOD):    {"MOD", "取模"},
		int(OP_CODE__MSIZE):  {"MSIZE", "获取活动内存的大小"},
		int(OP_CODE__MSTORE): {"MSTORE", "存储到内存"},
		int(OP_CODE__MUL):    {"MUL", "乘"},
		int(OP_CODE__NOT):    {"NOT", "按位取反运算"},
		int(OP_CODE__OR):     {"OR", "按位或运算"},
		int(OP_CODE__PC):     {"PC", "获取程序计数器"},
		int(OP_CODE__POP):    {"POP", "出栈"},
		int(OP_CODE__PUSH1):  {"PUSH1", "入栈1字节"},
		int(OP_CODE__PUSH10): {"PUSH10", "入栈10字节"},
		int(OP_CODE__PUSH11): {"PUSH11", "入栈11字节"},
		int(OP_CODE__PUSH12): {"PUSH12", "入栈12字节"},
		int(OP_CODE__PUSH13): {"PUSH13", "入栈13字节"},
		int(OP_CODE__PUSH14): {"PUSH14", "入栈14字节"},
		int(OP_CODE__PUSH15): {"PUSH15", "入栈15字节"},
		int(OP_CODE__PUSH16): {"PUSH16", "入栈16字节"},
		int(OP_CODE__PUSH17): {"PUSH17", "入栈17字节"},
		int(OP_CODE__PUSH18): {"PUSH18", "入栈18字节"},
		int(OP_CODE__PUSH19): {"PUSH19", "入栈19字节"},
		int(OP_CODE__PUSH2):  {"PUSH2", "入栈2字节"},
		int(OP_CODE__PUSH20): {"PUSH20", "入栈20字节"},
		int(OP_CODE__PUSH21): {"PUSH21", "入栈21字节"},
		int(OP_CODE__PUSH22): {"PUSH22", "入栈22字节"},
		int(OP_CODE__PUSH23): {"PUSH23", "入栈23字节"},
		int(OP_CODE__PUSH24): {"PUSH24", "入栈24字节"},
		int(OP_CODE__PUSH25): {"PUSH25", "入栈25字节"},
		int(OP_CODE__PUSH26): {"PUSH26", "入栈26字节"},
		int(OP_CODE__PUSH27): {"PUSH27", "入栈27字节"},
		int(OP_CODE__PUSH28): {"PUSH28", "入栈28字节"},
		int(OP_CODE__PUSH29): {"PUSH29", "入栈29字节"},
		int(OP_CODE__PUSH3):  {"PUSH3", "入栈3字节"},
		int(OP_CODE__PUSH30): {"PUSH30", "入栈30字节"},
		int(OP_CODE__PUSH31): {"PUSH31", "入栈31字节"},
		int(OP_CODE__PUSH32): {"PUSH32", "入栈32字节"},
		int(OP_CODE__PUSH4):  {"PUSH4", "入栈4字节"},
		int(OP_CODE__PUSH5):  {"PUSH5", "入栈5字节"},
		int(OP_CODE__PUSH6):  {"PUSH6", "入栈6字节"},
		int(OP_CODE__PUSH7):  {"PUSH7", "入栈7字节"},
		int(OP_CODE__PUSH8):  {"PUSH8", "入栈8字节"},
		int(OP_CODE__PUSH9):  {"PUSH9", "入栈9字节"},
		int(OP_CODE__SLOAD):  {"SLOAD", "从存储器加载"},
		int(OP_CODE__SSTORE): {"SSTORE", "存储到存储器"},
		int(OP_CODE__STOP):   {"STOP", "停止"},
		int(OP_CODE__SUB):    {"SUB", "减"},
		int(OP_CODE__SWAP10): {"SWAP10", "栈顶第10个元素与栈顶元素交换"},
		int(OP_CODE__SWAP11): {"SWAP11", "栈顶第11个元素与栈顶元素交换"},
		int(OP_CODE__SWAP12): {"SWAP12", "栈顶第12个元素与栈顶元素交换"},
		int(OP_CODE__SWAP13): {"SWAP13", "栈顶第13个元素与栈顶元素交换"},
		int(OP_CODE__SWAP14): {"SWAP14", "栈顶第14个元素与栈顶元素交换"},
		int(OP_CODE__SWAP15): {"SWAP15", "栈顶第15个元素与栈顶元素交换"},
		int(OP_CODE__SWAP16): {"SWAP16", "栈顶第16个元素与栈顶元素交换"},
		int(OP_CODE__SWAP2):  {"SWAP2", "栈顶第2个元素与栈顶元素交换"},
		int(OP_CODE__SWAP3):  {"SWAP3", "栈顶第3个元素与栈顶元素交换"},
		int(OP_CODE__SWAP4):  {"SWAP4", "栈顶第4个元素与栈顶元素交换"},
		int(OP_CODE__SWAP5):  {"SWAP5", "栈顶第5个元素与栈顶元素交换"},
		int(OP_CODE__SWAP6):  {"SWAP6", "栈顶第6个元素与栈顶元素交换"},
		int(OP_CODE__SWAP7):  {"SWAP7", "栈顶第7个元素与栈顶元素交换"},
		int(OP_CODE__SWAP8):  {"SWAP8", "栈顶第8个元素与栈顶元素交换"},
		int(OP_CODE__SWAP9):  {"SWAP9", "栈顶第9个元素与栈顶元素交换"},
		int(OP_CODE__XOR):    {"XOR", "按位异或运算"},
	}
}
func (v OpCode) String() string {
	switch v {
	case OP_CODE_UNKNOWN:
		return ""
	case OP_CODE__ADD:
		return "ADD"
	case OP_CODE__AND:
		return "AND"
	case OP_CODE__DIV:
		return "DIV"
	case OP_CODE__DUP1:
		return "DUP1"
	case OP_CODE__DUP10:
		return "DUP10"
	case OP_CODE__DUP11:
		return "DUP11"
	case OP_CODE__DUP12:
		return "DUP12"
	case OP_CODE__DUP13:
		return "DUP13"
	case OP_CODE__DUP14:
		return "DUP14"
	case OP_CODE__DUP15:
		return "DUP15"
	case OP_CODE__DUP16:
		return "DUP16"
	case OP_CODE__DUP2:
		return "DUP2"
	case OP_CODE__DUP3:
		return "DUP3"
	case OP_CODE__DUP4:
		return "DUP4"
	case OP_CODE__DUP5:
		return "DUP5"
	case OP_CODE__DUP6:
		return "DUP6"
	case OP_CODE__DUP7:
		return "DUP7"
	case OP_CODE__DUP8:
		return "DUP8"
	case OP_CODE__DUP9:
		return "DUP9"
	case OP_CODE__EQ:
		return "EQ"
	case OP_CODE__EXP:
		return "EXP"
	case OP_CODE__GT:
		return "GT"
	case OP_CODE__JUMP:
		return "JUMP"
	case OP_CODE__JUMPI:
		return "JUMPI"
	case OP_CODE__LT:
		return "LT"
	case OP_CODE__MLOAD:
		return "MLOAD"
	case OP_CODE__MOD:
		return "MOD"
	case OP_CODE__MSIZE:
		return "MSIZE"
	case OP_CODE__MSTORE:
		return "MSTORE"
	case OP_CODE__MUL:
		return "MUL"
	case OP_CODE__NOT:
		return "NOT"
	case OP_CODE__OR:
		return "OR"
	case OP_CODE__PC:
		return "PC"
	case OP_CODE__POP:
		return "POP"
	case OP_CODE__PUSH1:
		return "PUSH1"
	case OP_CODE__PUSH10:
		return "PUSH10"
	case OP_CODE__PUSH11:
		return "PUSH11"
	case OP_CODE__PUSH12:
		return "PUSH12"
	case OP_CODE__PUSH13:
		return "PUSH13"
	case OP_CODE__PUSH14:
		return "PUSH14"
	case OP_CODE__PUSH15:
		return "PUSH15"
	case OP_CODE__PUSH16:
		return "PUSH16"
	case OP_CODE__PUSH17:
		return "PUSH17"
	case OP_CODE__PUSH18:
		return "PUSH18"
	case OP_CODE__PUSH19:
		return "PUSH19"
	case OP_CODE__PUSH2:
		return "PUSH2"
	case OP_CODE__PUSH20:
		return "PUSH20"
	case OP_CODE__PUSH21:
		return "PUSH21"
	case OP_CODE__PUSH22:
		return "PUSH22"
	case OP_CODE__PUSH23:
		return "PUSH23"
	case OP_CODE__PUSH24:
		return "PUSH24"
	case OP_CODE__PUSH25:
		return "PUSH25"
	case OP_CODE__PUSH26:
		return "PUSH26"
	case OP_CODE__PUSH27:
		return "PUSH27"
	case OP_CODE__PUSH28:
		return "PUSH28"
	case OP_CODE__PUSH29:
		return "PUSH29"
	case OP_CODE__PUSH3:
		return "PUSH3"
	case OP_CODE__PUSH30:
		return "PUSH30"
	case OP_CODE__PUSH31:
		return "PUSH31"
	case OP_CODE__PUSH32:
		return "PUSH32"
	case OP_CODE__PUSH4:
		return "PUSH4"
	case OP_CODE__PUSH5:
		return "PUSH5"
	case OP_CODE__PUSH6:
		return "PUSH6"
	case OP_CODE__PUSH7:
		return "PUSH7"
	case OP_CODE__PUSH8:
		return "PUSH8"
	case OP_CODE__PUSH9:
		return "PUSH9"
	case OP_CODE__SLOAD:
		return "SLOAD"
	case OP_CODE__SSTORE:
		return "SSTORE"
	case OP_CODE__STOP:
		return "STOP"
	case OP_CODE__SUB:
		return "SUB"
	case OP_CODE__SWAP10:
		return "SWAP10"
	case OP_CODE__SWAP11:
		return "SWAP11"
	case OP_CODE__SWAP12:
		return "SWAP12"
	case OP_CODE__SWAP13:
		return "SWAP13"
	case OP_CODE__SWAP14:
		return "SWAP14"
	case OP_CODE__SWAP15:
		return "SWAP15"
	case OP_CODE__SWAP16:
		return "SWAP16"
	case OP_CODE__SWAP2:
		return "SWAP2"
	case OP_CODE__SWAP3:
		return "SWAP3"
	case OP_CODE__SWAP4:
		return "SWAP4"
	case OP_CODE__SWAP5:
		return "SWAP5"
	case OP_CODE__SWAP6:
		return "SWAP6"
	case OP_CODE__SWAP7:
		return "SWAP7"
	case OP_CODE__SWAP8:
		return "SWAP8"
	case OP_CODE__SWAP9:
		return "SWAP9"
	case OP_CODE__XOR:
		return "XOR"
	}
	return "UNKNOWN"
}

func (v OpCode) Label() string {
	switch v {
	case OP_CODE_UNKNOWN:
		return ""
	case OP_CODE__ADD:
		return "加"
	case OP_CODE__AND:
		return "按位与运算"
	case OP_CODE__DIV:
		return "除"
	case OP_CODE__DUP1:
		return "复制栈顶第1个元素"
	case OP_CODE__DUP10:
		return "复制栈顶第10个元素"
	case OP_CODE__DUP11:
		return "复制栈顶第11个元素"
	case OP_CODE__DUP12:
		return "复制栈顶第12个元素"
	case OP_CODE__DUP13:
		return "复制栈顶第13个元素"
	case OP_CODE__DUP14:
		return "复制栈顶第14个元素"
	case OP_CODE__DUP15:
		return "复制栈顶第15个元素"
	case OP_CODE__DUP16:
		return "复制栈顶第16个元素"
	case OP_CODE__DUP2:
		return "复制栈顶第2个元素"
	case OP_CODE__DUP3:
		return "复制栈顶第3个元素"
	case OP_CODE__DUP4:
		return "复制栈顶第4个元素"
	case OP_CODE__DUP5:
		return "复制栈顶第5个元素"
	case OP_CODE__DUP6:
		return "复制栈顶第6个元素"
	case OP_CODE__DUP7:
		return "复制栈顶第7个元素"
	case OP_CODE__DUP8:
		return "复制栈顶第8个元素"
	case OP_CODE__DUP9:
		return "复制栈顶第9个元素"
	case OP_CODE__EQ:
		return "等于"
	case OP_CODE__EXP:
		return "幂"
	case OP_CODE__GT:
		return "大于"
	case OP_CODE__JUMP:
		return "跳转"
	case OP_CODE__JUMPI:
		return "有条件跳转"
	case OP_CODE__LT:
		return "小于"
	case OP_CODE__MLOAD:
		return "从内存加载"
	case OP_CODE__MOD:
		return "取模"
	case OP_CODE__MSIZE:
		return "获取活动内存的大小"
	case OP_CODE__MSTORE:
		return "存储到内存"
	case OP_CODE__MUL:
		return "乘"
	case OP_CODE__NOT:
		return "按位取反运算"
	case OP_CODE__OR:
		return "按位或运算"
	case OP_CODE__PC:
		return "获取程序计数器"
	case OP_CODE__POP:
		return "出栈"
	case OP_CODE__PUSH1:
		return "入栈1字节"
	case OP_CODE__PUSH10:
		return "入栈10字节"
	case OP_CODE__PUSH11:
		return "入栈11字节"
	case OP_CODE__PUSH12:
		return "入栈12字节"
	case OP_CODE__PUSH13:
		return "入栈13字节"
	case OP_CODE__PUSH14:
		return "入栈14字节"
	case OP_CODE__PUSH15:
		return "入栈15字节"
	case OP_CODE__PUSH16:
		return "入栈16字节"
	case OP_CODE__PUSH17:
		return "入栈17字节"
	case OP_CODE__PUSH18:
		return "入栈18字节"
	case OP_CODE__PUSH19:
		return "入栈19字节"
	case OP_CODE__PUSH2:
		return "入栈2字节"
	case OP_CODE__PUSH20:
		return "入栈20字节"
	case OP_CODE__PUSH21:
		return "入栈21字节"
	case OP_CODE__PUSH22:
		return "入栈22字节"
	case OP_CODE__PUSH23:
		return "入栈23字节"
	case OP_CODE__PUSH24:
		return "入栈24字节"
	case OP_CODE__PUSH25:
		return "入栈25字节"
	case OP_CODE__PUSH26:
		return "入栈26字节"
	case OP_CODE__PUSH27:
		return "入栈27字节"
	case OP_CODE__PUSH28:
		return "入栈28字节"
	case OP_CODE__PUSH29:
		return "入栈29字节"
	case OP_CODE__PUSH3:
		return "入栈3字节"
	case OP_CODE__PUSH30:
		return "入栈30字节"
	case OP_CODE__PUSH31:
		return "入栈31字节"
	case OP_CODE__PUSH32:
		return "入栈32字节"
	case OP_CODE__PUSH4:
		return "入栈4字节"
	case OP_CODE__PUSH5:
		return "入栈5字节"
	case OP_CODE__PUSH6:
		return "入栈6字节"
	case OP_CODE__PUSH7:
		return "入栈7字节"
	case OP_CODE__PUSH8:
		return "入栈8字节"
	case OP_CODE__PUSH9:
		return "入栈9字节"
	case OP_CODE__SLOAD:
		return "从存储器加载"
	case OP_CODE__SSTORE:
		return "存储到存储器"
	case OP_CODE__STOP:
		return "停止"
	case OP_CODE__SUB:
		return "减"
	case OP_CODE__SWAP10:
		return "栈顶第10个元素与栈顶元素交换"
	case OP_CODE__SWAP11:
		return "栈顶第11个元素与栈顶元素交换"
	case OP_CODE__SWAP12:
		return "栈顶第12个元素与栈顶元素交换"
	case OP_CODE__SWAP13:
		return "栈顶第13个元素与栈顶元素交换"
	case OP_CODE__SWAP14:
		return "栈顶第14个元素与栈顶元素交换"
	case OP_CODE__SWAP15:
		return "栈顶第15个元素与栈顶元素交换"
	case OP_CODE__SWAP16:
		return "栈顶第16个元素与栈顶元素交换"
	case OP_CODE__SWAP2:
		return "栈顶第2个元素与栈顶元素交换"
	case OP_CODE__SWAP3:
		return "栈顶第3个元素与栈顶元素交换"
	case OP_CODE__SWAP4:
		return "栈顶第4个元素与栈顶元素交换"
	case OP_CODE__SWAP5:
		return "栈顶第5个元素与栈顶元素交换"
	case OP_CODE__SWAP6:
		return "栈顶第6个元素与栈顶元素交换"
	case OP_CODE__SWAP7:
		return "栈顶第7个元素与栈顶元素交换"
	case OP_CODE__SWAP8:
		return "栈顶第8个元素与栈顶元素交换"
	case OP_CODE__SWAP9:
		return "栈顶第9个元素与栈顶元素交换"
	case OP_CODE__XOR:
		return "按位异或运算"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*OpCode)(nil)

func (v OpCode) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidOpCode
	}
	return []byte(str), nil
}

func (v *OpCode) UnmarshalText(data []byte) (err error) {
	*v, err = ParseOpCodeFromString(string(bytes.ToUpper(data)))
	return
}
