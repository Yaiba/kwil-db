package tree

import (
	"fmt"
)

// SafeToSQL is a helper function that calls ToSQL on the given node and
// recovers from any panics.
func SafeToSQL(node Ast) (str string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err2, ok := r.(error)
			if !ok {
				err2 = fmt.Errorf("%v", r)
			}

			err = err2
		}
	}()

	return node.ToSQL(), nil
}
