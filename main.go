package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("./messages.txt")

	if err != nil {
		fmt.Printf("failed to read file: %v", err)
		return
	}

	var line string
	for {

		data := make([]byte, 8)
		count, err := file.Read(data)

		if err != nil {
			os.Exit(0)
		}

		parts := strings.Split(string(data[:count]), "\n")

		line += parts[0]

		if len(parts) == 2 {
			fmt.Printf("read: %s\n", line)
			line = parts[1]
		}
	}
}
