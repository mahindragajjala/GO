package different_types

import "fmt"

func Using_Function_Return() {
	value, ok := someFunction()
	if ok {
		fmt.Println("executed", value)
	} else {
		fmt.Println("not executed")
	}
}

func someFunction() (int, bool) {
	return 42, true
}
