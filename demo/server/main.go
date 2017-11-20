package main

import (
	"bufio"
	"fmt"
	"net"
)

var (
	conns []net.Conn
)

func main() {
	// Initialize server
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server listening on 8081...")

	// Listen for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
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
		panic(err)
	}
	fmt.Printf("Assigned site %d to %s.\n", n, conn.RemoteAddr().String())

	// Begin reading the data
	r := bufio.NewReader(conn)
	for {
		s, err := r.ReadString(0)
		if err != nil {
			panic(err)
		}

		// Broadcast
		for i, c := range conns {
			if i+1 != int(n) {
				_, err := c.Write([]byte(s))
				if err != nil {
					panic(err)
				}
			}
		}
	}

	fmt.Printf("Dropped connection from %s.\n", conn.RemoteAddr().String())
	panic("Connection dropped!")
}
