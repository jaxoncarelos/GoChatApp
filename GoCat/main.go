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
	for {
		input := make([]byte, 1024)
		len, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		toSend := append([]byte(username+": "), input[:len-1]...)
		listener.Write(toSend)
	}
}
