package vm

import "github.com/profzone/data_structures/algorithm"

func opStop(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error) {
	return nil, nil
}

func opAdd(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error) {
	a, b := stack.pop(), stack.peek()
	b.Add(a, b)

	return nil, nil
}

func makeOpPush(size int) executionFunc {
	return func(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error) {

		codeLen := len(code)

		// 检查code不要越界
		start := algorithm.IntMin(int64(codeLen), int64(*pc+1))
		end := algorithm.IntMin(int64(codeLen), start+int64(size))

		integer := stackElementPools.get().get()
		stack.push(integer.SetBytes(algorithm.RightPadBytes(code[start:end], size)))

		return nil, nil
	}
}

func makeOpDup(offset int) executionFunc {
	return func(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error) {

		stack.dup(stackElementPools.get(), offset)
		return nil, nil
	}
}

func makeOpSwap(offset int) executionFunc {
	return func(pc *uint64, code []byte, memory *Memory, stack *Stack) ([]byte, error) {

		stack.swap(offset)
		return nil, nil
	}
}
