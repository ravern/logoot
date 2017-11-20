package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	termbox "github.com/nsf/termbox-go"
	"github.com/ravernkoh/logoot/doc"
)

var (
	site uint8
	ldoc *doc.Doc // Logoot document
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

	c := strings.Split("Demo of Logoot collaborative editing.", "")
	ldoc = doc.New(c, site)

	ui(conn)
	go read(conn)

	fmt.Println("Disconnected!")
}

func ui(conn net.Conn) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	w, h := termbox.Size()
	p := doc.End

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		i, _ := ldoc.Index(p)
		i--
		termbox.SetCursor(i%w, i/w)
		r := []rune(ldoc.Content())
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				j := x + y*w
				if j < len(r) {
					termbox.SetCell(x, y, r[j], termbox.ColorWhite, termbox.ColorDefault)
				}
			}
		}
		termbox.Flush()

		e := termbox.PollEvent()
		switch e.Key {
		// Quit
		case termbox.KeyCtrlC:
			os.Exit(0)

		// Navigating
		case termbox.KeyArrowUp:
			if i-w > 0 {
				for j := 0; j < w; j++ {
					p, _ = ldoc.Left(p)
				}
			}
		case termbox.KeyArrowDown:
			if i+w < len(r) {
				for j := 0; j < w; j++ {
					p, _ = ldoc.Right(p)
				}
			}
		case termbox.KeyArrowLeft:
			if i > 0 {
				p, _ = ldoc.Left(p)
			}
		case termbox.KeyArrowRight:
			if i < len(r) {
				p, _ = ldoc.Right(p)
			}

		// Typing
		case termbox.KeyBackspace:
			fallthrough
		case termbox.KeyBackspace2:
			if i != 0 {
				ldoc.DeleteLeft(p)
			}
		default:
			p, _ = ldoc.InsertLeft(p, string(e.Ch))
			p, _ = ldoc.Right(p)
		}
	}
}

func read(conn net.Conn) {
	// Create main reader
	r := bufio.NewReader(conn)

	for {
		s, err := r.ReadString(0)
		if err != nil {
			panic(err)
		}
		fmt.Println([]byte(s)[0])
	}
}
