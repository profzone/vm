package vm

type Interpreter interface {
	Run(code []byte, input []byte) ([]byte, error)
}
