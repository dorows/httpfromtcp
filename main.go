/*
Creating a channel of strings to use insted of a state machine.
The goroutine (go func() {}()) contains the logic now, but it sends one at a time to the channel.
*/
package main

import (
	"fmt"
	"io"
	"os"
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
	file, _ := os.Open("messages.txt")

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
