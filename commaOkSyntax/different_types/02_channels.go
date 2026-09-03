package different_types

import "fmt"

func Using_Channels() {
	channel := make(chan int, 2)
	channel <- 1
	value, ok := <-channel
	if ok {
		fmt.Println("executed", value)
	} else {
		fmt.Println("not executed")
	}
	channel <- 1
	value2, ok := <-channel
	if ok {
		fmt.Println("executed", value2)
	} else {
		fmt.Println("not executed")
	}

}
