package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("What's your name?")
    
    if scanner.Scan() {
        fmt.Printf("Hello, %s!\n", scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
    }
}



/*
How it works:

bufio.NewScanner(os.Stdin) creates a new scanner that reads from standard input (the terminal).
scanner.Scan() reads the next line.
scanner.Text() gives you the actual line of input as a string.
The if err := scanner.Err() check ensures there are no issues during the scanning process.
In this example, the Scanner buffers input for efficiency, but you don’t need to think about the buffer itself — it just works.
*/
