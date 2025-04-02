package main

import (
	"fmt"
	r "github.com/Ferps53/httpfromtcp/internal/request"
	"log"
	"net"
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

		request, err := r.RequestFromReader(conn)

    if err != nil {
      log.Fatal(err)
    }

    fmt.Println(request.RequestLine)
	}
}
