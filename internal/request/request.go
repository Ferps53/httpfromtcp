package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

const crlf = "\r\n"
const bufferSize = 8

type Request struct {
	RequestLine
	state requestState
}

type requestState int

const (
	requestStateInitialized requestState = iota
	requestStateDone
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (rl RequestLine) String() string {
	return fmt.Sprintf(
		"Request line:\n- Method: %s\n- Target: %s\n- Version: %s",
		rl.Method,
		rl.RequestTarget,
		rl.HttpVersion,
	)
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	buf := make([]byte, bufferSize, bufferSize)
	readToIndex := 0

	request := &Request{
		state: requestStateInitialized,
	}

	for request.state != requestStateDone {
		if readToIndex >= len(buf) {
			newBuf := make([]byte, len(buf)*2)
			copy(newBuf, buf)
			buf = newBuf
		}

		numBytesRead, err := reader.Read(buf[readToIndex:])
		if err != nil {
			if errors.Is(io.EOF, err) {
				request.state = requestStateDone
				break
			}

			return nil, err
		}

		readToIndex += numBytesRead
		numBytesParsed, err := request.parse(buf[:readToIndex])
		if err != nil {
			return nil, err
		}

		copy(buf, buf[numBytesParsed:])
		readToIndex -= numBytesParsed
	}

	return request, nil
}

func parseRequestLine(data []byte) (*RequestLine, int, error) {

	idx := bytes.Index(data, []byte(crlf))

	if idx == -1 {
		return nil, 0, nil
	}

	line := string(data[:idx])

	requestLineSplit := strings.Split(line, " ")

	if err := validateRequestLine(requestLineSplit); err != nil {
		return nil, 0, err
	}

	return &RequestLine{
			Method:        requestLineSplit[0],
			RequestTarget: requestLineSplit[1],
			HttpVersion:   strings.Split(requestLineSplit[2], "/")[1],
		},
		idx + 2,
		nil
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

func (request *Request) parse(data []byte) (int, error) {

	switch request.state {
	case requestStateInitialized:
		requestLine, n, err := parseRequestLine(data)

		if err != nil {
			return 0, err
		}

		if n == 0 {
			return 0, nil
		}

		request.RequestLine = *requestLine
		request.state = requestStateDone

		return n, nil
	case requestStateDone:
		return 0, fmt.Errorf("error: Data already done")
	default:
		return 0, fmt.Errorf("unknown state")
	}
}
