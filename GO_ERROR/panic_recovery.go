What is panic?
panic stops the normal flow of execution and begins panicking. When a function panics:
It starts unwinding the stack (like throwing an exception).
If the program doesn't recover from it, the application crashes.






2. What is recover?
recover is used to regain control of a panicking goroutine. 
It only works inside a defer block. 
If called during a panic, it stops the panic and returns the error passed to panic.



package main

import "fmt"

func divide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Result:", a/b) // panic if b is 0
}

func main() {
	divide(10, 0) // triggers panic
	fmt.Println("Continuing execution...")
}

OUTPUT:
Recovered from panic: runtime error: integer divide by zero
Continuing execution...




