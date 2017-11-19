package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server listening on 8081...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("New connection from %s!\n", conn.RemoteAddr().String())

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	r := bufio.NewReader(conn)

	s, err := r.ReadString('b')
	for err != io.EOF {
		fmt.Println(s)
		s, err = r.ReadString('b')
	}

	fmt.Printf("Dropped connection from %s.\n", conn.RemoteAddr().String())
}
