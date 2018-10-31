package vm

import (
	"fmt"
	"github.com/profzone/vm/conf"
)

func makeStackValidationFunc(pop, push int) stackValidationFunc {
	return func(stack *Stack) error {
		if err := stack.require(pop); err != nil {
			return err
		}

		if stack.len()+push-pop > int(conf.Config.StackLimit) {
			return fmt.Errorf("stack limit reached %d (%d)", stack.len(), conf.Config.StackLimit)
		}
		return nil
	}
}
