package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <address> <port> <username>")
		return
	}
	address := args[1]
	port := args[2]
	username := args[3]
	listener, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
  fmt.Print("Enter chatroom number: ")
	var chatroom int
  fmt.Scanln(&chatroom)
  fmt.Println("\033[H\033[2J\nWelcome")
  listener.Write([]byte(fmt.Sprintf(`{"chatroom": %d, "text": "A person has joined the chatroom", "username": "%s"}`, chatroom, username)))
	defer listener.Close()
	go func() {
		for {
			read := make([]byte, 1024)
			n, err := listener.Read(read)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(read[:n]))
		}
	}()
	// clear terminal
	for {
		input := make([]byte, 1024)
		len, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		toSend := fmt.Sprintf(`{"chatroom": %d, "text": "%s", "username": "%s"}`, chatroom, string(input[:len-1]), username)
		listener.Write([]byte(toSend))
	}
}
