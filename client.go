package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	const addr = "localhost:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	data, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	if string(data) != "OK\n" {
		fmt.Fprintf(os.Stderr, "unexpected response: %q\n", data)
		os.Exit(1)
	}
	fmt.Println("OK received")
}
