package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	request := new(Request)
	data, err := io.ReadAll(reader)

	if err != nil {
		return request, err
	}

	response := string(data)

	return readRequestLine(request, response)
}

func readRequestLine(request *Request, response string) (*Request, error) {
	lines := strings.Split(response, "\r\n")
	requestLineSplit := strings.Split(lines[0], " ")

	if err := validateRequestLine(requestLineSplit); err != nil {
		return nil, err
	}

	request.Method = requestLineSplit[0]
	request.RequestTarget = requestLineSplit[1]
	request.HttpVersion = strings.Split(requestLineSplit[2], "/")[1]

	return request, nil
}

func validateRequestLine(requestLineSplit []string) error {

	if len(requestLineSplit) != 3 {
		return fmt.Errorf("Invalid Request Line")
	}

	if err := validateMethod(requestLineSplit[0]); err != nil {
		return err
	}

	return nil
}

func validateMethod(method string) error {

	switch method {
	case "POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD":
		return nil

	default:
		return fmt.Errorf("Method: %s does not exist!", method)
	}
}
