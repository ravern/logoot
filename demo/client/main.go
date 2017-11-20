package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	site uint8
)

func main() {
	fmt.Println("Connecting to server on 8081...")

	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Create main reader
	r := bufio.NewReader(conn)

	// Read site id
	s, err := r.ReadString(0)
	if err != nil {
		panic(err)
	}
	site = uint8([]byte(s)[0])
	fmt.Printf("Assigned site %d.\n", site)

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// TODO tmp
	go func() {
		for {
			time.Sleep(time.Second)
			_, err := conn.Write([]byte{site, 0})
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}()

	// Read
	for {
		s, err = r.ReadString(0)
		if err != nil {
			panic(err)
		}
		fmt.Println([]byte(s)[0])
	}

	fmt.Println("Disconnected!")
}
