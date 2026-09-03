package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// 1. Start listening on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("🚀 Server started on port 8080")

	for {
		// 2. Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}

		// 3. Handle client in goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("✅ Client connected:", conn.RemoteAddr())

	// Read data from client
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')

	fmt.Println("📩 Received:", message)

	// Send response back to client
	conn.Write([]byte("Hello Client, message received!\n"))
}
