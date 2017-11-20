package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	conns []net.Conn
)

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server listening on 8081...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("New connection from %s!\n", conn.RemoteAddr().String())

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	conns = append(conns, conn)

	// Assign the site id
	n := uint8(len(conns))
	_, err := conn.Write([]byte{n, 0})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Assigned site %d to %s.\n", n, conn.RemoteAddr().String())

	// Begin reading the data
	_ = bufio.NewReader(conn)
	for {
	}

	fmt.Printf("Dropped connection from %s.\n", conn.RemoteAddr().String())
}
