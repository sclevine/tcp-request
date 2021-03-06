package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	defer l.Close()
	fmt.Println("Listening.")
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println(string(buf))
	conn.Write([]byte("HTTP/1.1 200 OK\n\n"))
	conn.Write(buf)
	conn.Close()
}
