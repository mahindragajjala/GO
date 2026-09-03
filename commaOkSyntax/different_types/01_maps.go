package different_types

import "fmt"

func Using_Maps() {
	data := make(map[string]int)
	data["one"] = 1
	data["two"] = 2
	data["three"] = 3
	not, ok := data["one"]
	if ok {
		fmt.Println(not)
	} else {
		fmt.Println("Not found")
	}
}
