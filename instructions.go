package vm

func opStop(pc *int64, memory *Memory, stack *Stack) ([]byte, error) {
	return nil, nil
}

func opAdd(pc *int64, memory *Memory, stack *Stack) ([]byte, error) {
	a, b := stack.pop(), stack.peek()
	b.Add(a, b)

	return nil, nil
}
