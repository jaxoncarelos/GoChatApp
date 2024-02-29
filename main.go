package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	out := CreateTCP()

	for {
		select {
		case msg := <-out:
      fmt.Println("Received message: ", msg)
		}
	}
}

func CreateTCP() (chan string) {
	out := make(chan string)
  conns := make([]net.Conn, 0)
	go func() {
		listener, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatal("Error!")
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
      conns = append(conns, conn)
			if err != nil {
				log.Fatal("Error!")
			}
			go func(conn net.Conn){
				defer conn.Close()
				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Error!")
						return
					}
          for _, c := range filter(conns, func(conn1 net.Conn) bool { return conn1 != conn }){
            c.Write(buf[:n])
          }
					out <- string(buf[:n])
				}
			}(conn)
		}
	}()
	return out
}
func filter[T any](array []T, f func(T) bool) ([]T) {
  new := make([]T, 0)
  for _, v := range array {
    if f(v) {
      new = append(new, v)
    }
  }
  return new
}
