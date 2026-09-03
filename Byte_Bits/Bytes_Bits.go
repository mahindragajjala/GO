package main

import "fmt"

func main() {
	//BYTES
	data := "xius"
	bytes := []byte(data)
	fmt.Println(bytes)

	//BITS
	for _, value := range bytes {
		data := fmt.Sprintf("%08b", value)
		fmt.Print(data)
	}
}
