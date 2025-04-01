package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	udp, err := net.ResolveUDPAddr("udp", ":42069")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, udp)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(text))

		if err != nil {
			log.Fatal(err)
		}
	}
}
