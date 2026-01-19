package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var current_line_contents string = ""

	for {
		buffer := make([]byte, 8)
		n, err := file.Read(buffer)

		chunk := string(buffer[:n])
		parts := strings.Split(chunk, "\n")

		if err != nil {
			if current_line_contents != "" {
				fmt.Printf("read: %s\n", current_line_contents)
				current_line_contents = ""
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("error: %s\n", err.Error())
			break
		}
		for i := 0; i < len(parts)-1; i++ {
			frase := current_line_contents + parts[i]
			fmt.Printf("read: %s", frase)
			current_line_contents = ""
		}
		current_line_contents = current_line_contents + parts[len(parts)-1]

		if err == io.EOF {
			break
		}
	}
	if current_line_contents != "" {
		fmt.Printf("read: %s\n", current_line_contents)
	}

}
