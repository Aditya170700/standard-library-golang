package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// bufio reader
	input := strings.NewReader("aditya ricki julianto\nselasa rabu kamis\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// bufio writer
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("aditya ricki julianto\n")
	writer.WriteString("senin selasa rabu\n")
	writer.Flush()
}
