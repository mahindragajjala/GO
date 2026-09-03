package different_types

import "fmt"

func Using_Type_Assertion() {
	var x interface{} = "Hello"
	str, ok := x.(string)
	if ok {
		fmt.Println("String value:", str)
	} else {
		fmt.Println("Not a string!")
	}
}
