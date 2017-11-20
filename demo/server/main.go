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
	// Initialize server
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server listening on 8081...")

	// Listen for connections
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
	r := bufio.NewReader(conn)
	for {
		s, err := r.ReadString(0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Broadcast
		for i, c := range conns {
			if i+1 != int(n) {
				_, err := c.Write([]byte(s))
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}
	}

	fmt.Printf("Dropped connection from %s.\n", conn.RemoteAddr().String())
	os.Exit(1)
}
