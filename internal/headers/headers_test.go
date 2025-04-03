package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:42069" + crlf + crlf)
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       " + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	//Test: Valid single header with extra space
	headers = NewHeaders()
	data = []byte(" Host: localhost:42069" + crlf + crlf + " ")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	//Test: Valid done
	headers = NewHeaders()
	data = []byte(crlf)
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.True(t, done)

	//Test: Valid 2 headers with existing headers
	headers = NewHeaders()
	data = []byte("Host: localhost:42069" + crlf + crlf)
	n, done, err = headers.Parse(data)

	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	data = []byte("Test: test" + crlf + crlf)
	n, done, err = headers.Parse(data)

	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "test", headers["test"])
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 12, n)
	assert.False(t, done)

	data = []byte(crlf)
	n, done, err = headers.Parse(data)

	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "test", headers["test"])
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 0, n)
	assert.True(t, done)

	//Test: Invalid chars in key
	headers = NewHeaders()
	data = []byte("H©st: localhost:42069" + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

  headers = NewHeaders()
  data = []byte("Host+: localhost:42069" + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host+"])
	assert.Equal(t, 24, n)
	assert.False(t, done)

  //Test: Multiple values
	headers = NewHeaders()
  data = []byte("Set-Test: test1" + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "test1", headers["set-test"])
	assert.Equal(t, 17, n)
	assert.False(t, done)

  data = []byte("Set-Test: test2" + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "test1, test2", headers["set-test"])
	assert.Equal(t, 17, n)
	assert.False(t, done)

  data = []byte("Set-Test: test3" + crlf + crlf)
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "test1, test2, test3", headers["set-test"])
	assert.Equal(t, 17, n)
	assert.False(t, done)
}
