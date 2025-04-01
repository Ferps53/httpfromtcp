package main

import (
	"fmt"
	"os"
)

func main() {

  file, err := os.Open("./messages.txt")

  if err != nil {
    fmt.Printf("failed to read file: %v", err);
    return;
  }

  for {

    data := make([]byte, 8)
    count, err := file.Read(data)

    if err != nil {
      os.Exit(0);
    }

    fmt.Printf("read: %s\n", string(data[:count]))
  }

}
