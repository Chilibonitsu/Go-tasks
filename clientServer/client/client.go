package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":4545")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	for {
		var input string
		fmt.Println("Введите сообщение: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Некорректный ввод")
			continue
		}
		if n, err := conn.Write([]byte(input)); n == 0 || err != nil {
			log.Fatal(err)
			return
		} else {
			fmt.Printf("write to server = %d bytes\n", n)
		}
		reply := make([]byte, 1024)
		n, err := conn.Read(reply)

		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("received: ", string(reply[0:n]))

	}

}
