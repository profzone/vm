package vm

//go:generate tools gen enum OpCode
// swagger:enum
type OpCode byte

// 操作运算符
const (
	OP_CODE_UNKNOWN OpCode = iota
	OP_CODE__STOP    // 停止
	OP_CODE__ADD     // 加
	OP_CODE__SUB     // 减
	OP_CODE__MUL     // 乘
	OP_CODE__DIV     // 除
	OP_CODE__MOD     // 取模
	OP_CODE__EXP     // 幂
)

// 逻辑运算符
const (
	OP_CODE__LT  OpCode = iota + 0x10 // 小于
	OP_CODE__GT                       // 大于
	OP_CODE__EQ                       // 等于
	OP_CODE__AND                      // 按位与运算
	OP_CODE__OR                       // 按位或运算
	OP_CODE__XOR                      // 按位异或运算
	OP_CODE__NOT                      // 按位取反运算
)

// 流程控制运算符
const (
	OP_CODE__POP    OpCode = iota + 0x50 // 出栈
	OP_CODE__MLOAD                       // 从内存加载
	OP_CODE__MSTORE                      // 存储到内存
	OP_CODE__SLOAD                       // 从存储器加载
	OP_CODE__SSTORE                      // 存储到存储器
	OP_CODE__JUMP                        // 跳转
	OP_CODE__JUMPI                       // 有条件跳转
	OP_CODE__PC                          // 获取程序计数器
	OP_CODE__MSIZE                       // 获取活动内存的大小
)

const (
	OP_CODE__PUSH1  OpCode = iota + 0x60 // 入栈1字节
	OP_CODE__PUSH2                       // 入栈2字节
	OP_CODE__PUSH3                       // 入栈3字节
	OP_CODE__PUSH4                       // 入栈4字节
	OP_CODE__PUSH5                       // 入栈5字节
	OP_CODE__PUSH6                       // 入栈6字节
	OP_CODE__PUSH7                       // 入栈7字节
	OP_CODE__PUSH8                       // 入栈8字节
	OP_CODE__PUSH9                       // 入栈9字节
	OP_CODE__PUSH10                      // 入栈10字节
	OP_CODE__PUSH11                      // 入栈11字节
	OP_CODE__PUSH12                      // 入栈12字节
	OP_CODE__PUSH13                      // 入栈13字节
	OP_CODE__PUSH14                      // 入栈14字节
	OP_CODE__PUSH15                      // 入栈15字节
	OP_CODE__PUSH16                      // 入栈16字节
	OP_CODE__PUSH17                      // 入栈17字节
	OP_CODE__PUSH18                      // 入栈18字节
	OP_CODE__PUSH19                      // 入栈19字节
	OP_CODE__PUSH20                      // 入栈20字节
	OP_CODE__PUSH21                      // 入栈21字节
	OP_CODE__PUSH22                      // 入栈22字节
	OP_CODE__PUSH23                      // 入栈23字节
	OP_CODE__PUSH24                      // 入栈24字节
	OP_CODE__PUSH25                      // 入栈25字节
	OP_CODE__PUSH26                      // 入栈26字节
	OP_CODE__PUSH27                      // 入栈27字节
	OP_CODE__PUSH28                      // 入栈28字节
	OP_CODE__PUSH29                      // 入栈29字节
	OP_CODE__PUSH30                      // 入栈30字节
	OP_CODE__PUSH31                      // 入栈31字节
	OP_CODE__PUSH32                      // 入栈32字节
	OP_CODE__DUP1                        // 复制栈顶第1个元素
	OP_CODE__DUP2                        // 复制栈顶第2个元素
	OP_CODE__DUP3                        // 复制栈顶第3个元素
	OP_CODE__DUP4                        // 复制栈顶第4个元素
	OP_CODE__DUP5                        // 复制栈顶第5个元素
	OP_CODE__DUP6                        // 复制栈顶第6个元素
	OP_CODE__DUP7                        // 复制栈顶第7个元素
	OP_CODE__DUP8                        // 复制栈顶第8个元素
	OP_CODE__DUP9                        // 复制栈顶第9个元素
	OP_CODE__DUP10                       // 复制栈顶第10个元素
	OP_CODE__DUP11                       // 复制栈顶第11个元素
	OP_CODE__DUP12                       // 复制栈顶第12个元素
	OP_CODE__DUP13                       // 复制栈顶第13个元素
	OP_CODE__DUP14                       // 复制栈顶第14个元素
	OP_CODE__DUP15                       // 复制栈顶第15个元素
	OP_CODE__DUP16                       // 复制栈顶第16个元素
	OP_CODE__SWAP2                       // 栈顶第2个元素与栈顶元素交换
	OP_CODE__SWAP3                       // 栈顶第3个元素与栈顶元素交换
	OP_CODE__SWAP4                       // 栈顶第4个元素与栈顶元素交换
	OP_CODE__SWAP5                       // 栈顶第5个元素与栈顶元素交换
	OP_CODE__SWAP6                       // 栈顶第6个元素与栈顶元素交换
	OP_CODE__SWAP7                       // 栈顶第7个元素与栈顶元素交换
	OP_CODE__SWAP8                       // 栈顶第8个元素与栈顶元素交换
	OP_CODE__SWAP9                       // 栈顶第9个元素与栈顶元素交换
	OP_CODE__SWAP10                      // 栈顶第10个元素与栈顶元素交换
	OP_CODE__SWAP11                      // 栈顶第11个元素与栈顶元素交换
	OP_CODE__SWAP12                      // 栈顶第12个元素与栈顶元素交换
	OP_CODE__SWAP13                      // 栈顶第13个元素与栈顶元素交换
	OP_CODE__SWAP14                      // 栈顶第14个元素与栈顶元素交换
	OP_CODE__SWAP15                      // 栈顶第15个元素与栈顶元素交换
	OP_CODE__SWAP16                      // 栈顶第16个元素与栈顶元素交换
)
