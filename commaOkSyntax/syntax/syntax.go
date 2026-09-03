package syntax

import "fmt"

func Syntax() {
	fmt.Println("Syntax")
	users := map[string]string{
		"john": "developer",
	}

	value, ok := users["john"]

	fmt.Println(value)
	fmt.Println(ok)
}
