package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Usage: go run main.go <address> <port>")
		return
	}
	out := CreateTCP(args[1:])
	for {
		select {
		case msg := <-out:
			fmt.Println("Received message: ", msg)
		}
	}
}

func CreateTCP(args []string) chan string {
	out := make(chan string)
	conns := make(map[int][]net.Conn)
	fmt.Println("Listening on " + args[0] + ":" + args[1])
	go func() {
		listener, err := net.Listen("tcp", args[0]+":"+args[1])
		if err != nil {
			log.Fatal("Error!", err)
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			// conns[-1] = append(conns[-1], conn)
			if err != nil {
				log.Fatal("Error!", err)
				continue
			}
			go func(conn net.Conn) {
				defer conn.Close()
				conn.Write([]byte("A person has joined the chatroom"))

				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Error!", err)
						return
					}

					message, err := tryDecodeJson(buf[:n])
					if err != nil {
						fmt.Println("Error decoding JSON:", err)
						continue
					}
					if !contains(conns[message.Chatroom], conn) {
						conns[message.Chatroom] = append(conns[message.Chatroom], conn)
					}

					for _, c := range filter(conns[message.Chatroom], func(conn1 net.Conn) bool { return conn1 != conn }) {
						c.Write([]byte(message.Username + ": " + message.Text))
					}

					out <- message.Username + ": " + message.Text
				}
			}(conn)
		}
	}()
	return out
}

type Message struct {
	Chatroom int    `json:"chatroom"`
	Text     string `json:"text"`
	Username string `json:"username"`
}

func tryDecodeJson(s []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(s, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}

func filter[T any](array []T, f func(T) bool) []T {
	new := make([]T, 0)
	for _, v := range array {
		if f(v) {
			new = append(new, v)
		}
	}
	return new
}

func contains(array []net.Conn, value net.Conn) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}
