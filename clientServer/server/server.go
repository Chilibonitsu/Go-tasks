package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue

		}
		go handleConn(conn)
	}

}
func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected", conn.LocalAddr().Network(), conn.LocalAddr())
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("read error:", err)
			break
		}
		ans := "Server: " + string(input[0:n])
		conn.Write([]byte(ans))
	}

}
