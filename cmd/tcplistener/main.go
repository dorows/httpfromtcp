/*
read lines from a TCP connection.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)

		var current_line_contents string = ""

		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)

			if n > 0 {
				chunk := string(buffer[:n])
				parts := strings.Split(chunk, "\n")

				for i := 0; i < len(parts)-1; i++ {
					frase := current_line_contents + parts[i]
					ch <- frase
					current_line_contents = ""
				}
				current_line_contents = current_line_contents + parts[len(parts)-1]
			}

			if err != nil {
				break
			}
		}

		if current_line_contents != "" {
			ch <- current_line_contents
		}
	}()

	return ch
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("connection has been accepted")

		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Println(line)
		}

		fmt.Println("connection has been closed")
	}
}
