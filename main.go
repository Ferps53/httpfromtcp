package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./messages.txt")

	if err != nil {
		fmt.Printf("failed to read file: %v", err)
		return
	}

	channel := getLinesChannel(file)

	defer file.Close()

	for line := range channel {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {

	channel := make(chan string)
	go readLines(file, channel)

	return channel
}

func readLines(file io.ReadCloser, channel chan string) {

	defer close(channel)
	var line string
	for {

		data := make([]byte, 8)
		count, err := file.Read(data)
		if err != nil {
			return
		}

		parts := strings.Split(string(data[:count]), "\n")

		line += parts[0]

		if len(parts) == 2 {
			channel <- line
			line = parts[1]
		}
	}
}
