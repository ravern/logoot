package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	site uint8
)

func main() {
	fmt.Println("Connecting to server on 8081...")

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")

	r := bufio.NewReader(conn)

	// Read site id
	s, err := r.ReadString(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	site = uint8([]byte(s)[0])
	fmt.Printf("Assigned site %d.\n", site)

	fmt.Println("Disconnected!")
}
