package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"

	termbox "github.com/nsf/termbox-go"
	"github.com/ravernkoh/logoot/doc"
)

var (
	site uint8
	ldoc *doc.Doc // Logoot document
	p    []doc.Pos
)

func main() {
	rand.Seed(10)

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

	go read(conn)
	ui(conn)

	fmt.Println("Disconnected!")
}

func render() {
	w, h := termbox.Size()
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
}

func ui(conn net.Conn) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	p = doc.End
	w, _ := termbox.Size()

	for {
		i, _ := ldoc.Index(p)
		i--
		r := []rune(ldoc.Content())
		render()

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
				pp, _ := ldoc.Left(p)
				ldoc.DeleteLeft(p)

				b := []byte{2}
				b = append(b, doc.PosBytes(pp)...)
				b = append(b, 233, 244, 255, 0)
				conn.Write(b)
			}
		default:
			p, _ = ldoc.InsertLeft(p, string(e.Ch))

			b := []byte{1}
			b = append(b, doc.PosBytes(p)...)
			b = append(b, byte(e.Ch))
			b = append(b, 233, 244, 255, 0)
			conn.Write(b)

			p, _ = ldoc.Right(p)
		}
	}
}

func read(conn net.Conn) {
	for {
		b, err := readTill(conn, []byte{233, 244, 255, 0})
		if err != nil {
			panic(err)
		}

		if b[0] == 1 {
			p := doc.NewPos(b[1:])
			ldoc.Insert(p, string(b[len(b)-1:]))
		} else if b[0] == 2 {
			p := doc.NewPos(b[1:])
			ldoc.Delete(p)
		}
		render()
	}
}

func readTill(conn net.Conn, delim []byte) (line []byte, err error) {
	r := bufio.NewReader(conn)
	for {
		s := ""
		s, err = r.ReadString(delim[len(delim)-1])
		if err != nil {
			return
		}

		line = append(line, []byte(s)...)
		if bytes.HasSuffix(line, delim) {
			return line[:len(line)-len(delim)], nil
		}
	}
}
