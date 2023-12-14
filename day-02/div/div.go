package main

import (
	"fmt"
)

func main() {
	// fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(7, 2))
}

func safeDiv(a, b int) (q int, err error) {

	defer func() {
		// e's type is any (or interface{}) *not* error
		if e := recover(); e != nil {
			// fmt.Println("panic occurred:", err)
			err = fmt.Errorf("panic occurred: %v", e)
		}
	}()

	return a / b, nil
}

func div(a, b int) int {
	return a / b
}
