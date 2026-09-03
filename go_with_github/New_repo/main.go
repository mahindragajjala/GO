package main

import (
	"fmt"

	"github.com/mahindragajjala/Go_With_Github/nested_package"
)

func main() {
	fmt.Println("Hello world")
	result := nested_package.Addition_function(5, 10)
	fmt.Println(result)
}
