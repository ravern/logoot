package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Connecting to server on 8081...")

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")

	conn.Write([]byte("Hello worlb"))
	for {
	}

	fmt.Println("Disconnected!")
}
