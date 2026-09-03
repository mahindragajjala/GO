package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 2. Send message
	fmt.Fprintf(conn, "Hello Server!\n")

	// 3. Read response
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("📨 Server replied:", response)

	fmt.Println("Press Enter to exit")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

/*
conn, err := net.Dial("tcp", "localhost:8080")

- "tcp" - Tells OS: use TCP protocol
- localhost → 127.0.0.1 (same machine) / 8080 → server’s listening port
- OS actions:
			- Creates a client socket
			- Assigns a "random client port" (e.g., 54231)
			- Performs TCP 3-way handshake
			- Connects to server socket
*/

/*
defer conn.Close()

defer schedules Close() at end of main()
Properly releases:
OS socket
File descriptor
Sends TCP FIN packet
*/

/*
fmt.Fprintf(conn, "Hello Server!\n")

fmt.Fprintf
Writes formatted text to any writer
conn
→ Implements io.Writer
→ Writing here = sending data to socket

"Hello Server!\n"
→ Data being sent
→ \n is important (server reads line-by-line)

Behind the scenes
		String → bytes
		Bytes → TCP send buffer
		OS sends packets over network
*/

/*
Read response from server
response, _ := bufio.NewReader(conn).ReadString('\n')

1.bufio.NewReader(conn)
	Wraps socket with a buffer
	Improves performance
	Allows line-based reading
2.ReadString('\n')
	Blocks execution
	Waits until:
	server sends data
	newline \n is received
3.Returns:
	response → string from server
	error (ignored here)


conn → socket connection
bufio.NewReader → makes reading easier & faster
ReadString('\n') → read until Enter (newline)
*/
