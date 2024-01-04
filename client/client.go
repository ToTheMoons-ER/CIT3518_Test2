package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// getUserInput ใช้สำหรับรับข้อมูลจากผู้ใช้
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(userInput)
}

// connectToServer เชื่อมต่อกับเซิร์ฟเวอร์และส่งข้อมูล
func connectToServer(address, data string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("Error sending data to server: %v", err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("Error receiving data from server: %v", err)
	}

	fmt.Println("Server response:", string(buffer[:n]))
	return nil
}

func main() {
	fmt.Println("Simple Chat Client")

	// Get user credentials
	username := getUserInput("Enter username: ")
	password := getUserInput("Enter password: ")

	// Format user credentials for sending
	data := fmt.Sprintf("%s:%s", username, password)

	// Connect to the server
	serverAddress := "localhost:12345"
	err := connectToServer(serverAddress, data)
	if err != nil {
		fmt.Println(err)
	}
}
