package main

import (
	"fmt"
	"net"
)

const (
	validUsername = "std1"
	validPassword = "p@ssw0rd"
	port          = 12345
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read client data
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])

	// Validate credentials
	if clientData == fmt.Sprintf("%s:%s", validUsername, validPassword) {
		conn.Write([]byte("Hello\n"))
	} else {
		conn.Write([]byte("Invalid credentials\n"))
	}
}

func startServer() {
	fmt.Println("Server is starting...")

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func main() {
	startServer()
}
