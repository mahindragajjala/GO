package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    file, err := os.Open("largefile.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Print(line)
    }
}
/*
When you need to read large files, Reader is your muscle. 
It provides a buffered interface for reading from an io.Reader—which can be anything from a file to a network connection.

Let’s say you want to read a large file line by line, rather than loading the entire 
file into memory (which is a rookie mistake for large files).

How it works:

os.Open("largefile.txt") opens the file for reading.
bufio.NewReader(file) creates a buffered reader that reads data from the file in chunks.
reader.ReadString('\n') reads a line at a time from the file until it hits a newline.
*/
