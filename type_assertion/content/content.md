Type Assertion

- I have a value stored inside an interface, and I want to check/retrive its actual concrete type.
- It is mainly used with interface values.
- x.(string) - is a type assertion
- I believe the actual value inside x is a string. Give me that string.
- An interface{} can hold values of different types
            var x interface{}

            x = 100
            x = "Hello"
            x = true
- At runtime, x can contain different concrete types.
- You first assert its type.
Syntax:
    package main

    import "fmt"

    func main() {
        var x interface{} = "Hello"

        s, ok := x.(string)

        fmt.Println(s)
        fmt.Println(ok)
    }