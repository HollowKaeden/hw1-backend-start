package main

import (
	"fmt"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	_, err := fmt.Fprint(conn, "OK\n")
	if err != nil {
		log.Println("write:", err)
	}
}

func main() {
	const addr = ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept:", err)
			continue
		}
		go handle(conn)
	}
}
