In Go, the `bufio` package provides **buffered I/O** functionality. It wraps around `io.Reader` and `io.Writer` interfaces to improve **I/O efficiency** by minimizing direct system calls. Let’s break it down step by step with detailed explanations and examples.

---

## 1. Why `bufio.Reader` and `bufio.Writer`?

### Without buffering:

* Reading/writing data directly (e.g., from a file or network) invokes **system calls** frequently.
* Each call is **expensive** in terms of performance.

### With buffering:

* Data is read/written in **chunks** (buffered).
* Reduces **number of system calls**, making I/O much faster.

---

## 2. `bufio.Reader`

### **Definition**

* `bufio.Reader` wraps an `io.Reader` (e.g., file, string reader, network connection).
* Provides **methods** like:

  * `Read()`
  * `ReadByte()`
  * `ReadLine()`
  * `ReadString()`
  * `Peek(n int)`

---

### **Constructor**

```go
func NewReader(rd io.Reader) *Reader
func NewReaderSize(rd io.Reader, size int) *Reader
```

* `NewReader` uses a **default buffer size (4096 bytes)**.
* `NewReaderSize` lets you define custom buffer size.

---

### **Example: bufio.Reader**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Example 1: Read line until '\n'
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err)
	} else {
		fmt.Println("Line read:", line)
	}

	// Example 2: Peek 5 bytes (without advancing reader)
	peekBytes, _ := reader.Peek(5)
	fmt.Println("Peeked bytes:", string(peekBytes))

	// Example 3: Read 10 bytes
	buffer := make([]byte, 10)
	n, _ := reader.Read(buffer)
	fmt.Printf("Read %d bytes: %s\n", n, buffer)
}
```

---

### **Key Points:**

1. `ReadString(delim byte)` reads until a delimiter (e.g., `\n`).
2. `Peek(n int)` shows data without moving the read pointer.
3. `Read()` reads into a byte slice.
4. Efficient for line-based or chunk-based reading.

---

## 3. `bufio.Writer`

### **Definition**

* `bufio.Writer` wraps an `io.Writer`.
* Buffers data in memory before writing to the underlying writer (file, socket).
* Provides **methods** like:

  * `Write()`
  * `WriteByte()`
  * `WriteString()`
  * `Flush()`

---

### **Constructor**

```go
func NewWriter(w io.Writer) *Writer
func NewWriterSize(w io.Writer, size int) *Writer
```

---

### **Important Method**

* `Flush()` — forces any buffered data to be written to the underlying writer.

---

### **Example: bufio.Writer**

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a file
	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)

	// Example 1: Write string
	writer.WriteString("Hello, this is buffered writing!\n")

	// Example 2: Write bytes
	writer.Write([]byte("Buffered I/O is efficient.\n"))

	// Flush the buffer to ensure data is written
	writer.Flush()

	fmt.Println("Data written to file using bufio.Writer")
}
```

---

### **Key Points:**

1. Data stays in buffer until:

   * Buffer is full
   * `Flush()` is called
   * Writer is closed (if wrapped properly)
2. Increases performance for multiple small writes.

---

## 4. Combined Example (Reader + Writer)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Input: Read from console
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := consoleReader.ReadString('\n')

	// Output: Write to file
	file, _ := os.Create("greeting.txt")
	writer := bufio.NewWriter(file)
	writer.WriteString("Hello, " + name)
	writer.Flush() // Ensure data is written

	fmt.Println("Greeting saved to greeting.txt")
}
```

---

## 5. Buffer Internals (How it Works)

### **Without bufio**

* Each `Write()` or `Read()` goes directly to disk or network.
* Multiple syscalls for small chunks.

### **With bufio**

* Data is **stored temporarily in RAM (buffer)**.
* `Flush()` writes data in **one go** → fewer syscalls.

---

## 6. Use Cases

* **Reading large files line by line** (log processing).
* **Writing large logs or data chunks efficiently**.
* **Reducing system call overhead in network programming**.

---

## 7. Important Notes

* Always call `Flush()` after writing.
* Buffer size can be tuned using `NewReaderSize` or `NewWriterSize`.
* Buffered reading **does not discard data** until it’s read.
* Combine with `io.Copy()` for advanced I/O.

---
