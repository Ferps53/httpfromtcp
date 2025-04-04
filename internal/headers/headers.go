package headers

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Headers map[string]string

const crlf = "\r\n"

func NewHeaders() Headers {
	return make(Headers)
}

func (header Headers) Parse(data []byte) (n int, done bool, err error) {

	dataString := string(data)
	dataString = strings.Trim(dataString, " ")
	idx := strings.Index(dataString, crlf)

	if idx == -1 {
		return 0, false, nil
	}

	if idx == 0 {
		return 0, true, nil
	}

	headerLine := strings.TrimSpace(dataString[:idx])
	parts := strings.Split(headerLine, " ")

	if !strings.Contains(parts[0], ":") {
		return 0, false, fmt.Errorf("Extra whitespace in header")
	}

	key := strings.ReplaceAll(parts[0], ":", "")

	if len(key) < 1 {
		return 0, false, fmt.Errorf("Len of key is less than one")
	}

	regex, err := regexp.Compile("[^A-Za-z1-9!#$%^&*+-_|~.`']")

	if err != nil {
		log.Fatal(err)
	}

	key = strings.ToLower(key)

	if regex.MatchString(key) {
		return 0, false, fmt.Errorf("Key has invalid chars")
	}

	curr, ok := header[key]

	if !ok {
		header[key] = parts[1]
	} else {
		fmt.Printf("Curr: %s\n", curr)
		partsValue := []string{curr, parts[1]}

		fmt.Println(partsValue)

		header[key] = strings.Join(partsValue, ", ")

		fmt.Println(header[key])

	}

	return idx + 2, false, nil
}
