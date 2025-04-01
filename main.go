package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	tcpListener, err := net.Listen("tcp", ":42069")
	defer tcpListener.Close()

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := tcpListener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connection accepted")

		channel := getLinesChannel(conn)

		for line := range channel {
			fmt.Println(line)
		}
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {

	channel := make(chan string)
	go readLines(file, channel)

	return channel
}

func readLines(file io.ReadCloser, channel chan string) {

	defer close(channel)
	fmt.Println("Connection closed")
	var line string
	for {
		data := make([]byte, 8)
		count, err := file.Read(data)

		if err != nil {
			if err == io.EOF {
				channel <- line
				return
			}
			log.Fatal(err)
		}

		parts := strings.Split(string(data[:count]), "\n")
		line += parts[0]

		if len(parts) == 2 {
			channel <- line
			line = parts[1]
		}
	}
}
