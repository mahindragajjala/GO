/*
In addition to reading, bufio also gives you Writer, which buffers data you’re writing to an output stream. 
It’s useful when you’re writing to files or sending data over the network, and you want to avoid writing byte-by-byte. 
Writer collects data in a buffer and writes it in chunks, improving performance.
*/
package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
    writer := bufio.NewWriter(file)
    writer.WriteString("Hello, Go with Buffered Output!\n")
    writer.Flush() // Don't forget to flush!
}

/*
How it works:

os.Create("output.txt") creates a new file for writing.
bufio.NewWriter(file) wraps the file with a buffered writer.
writer.WriteString() writes to the buffer.
writer.Flush() ensures the buffer is emptied into the file. Always call Flush to guarantee that data is written out.
*/
