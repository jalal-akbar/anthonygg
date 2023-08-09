package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type HasReader interface {
	io.Reader // embed Reader interface
	hash() string
}

func main() {
	payload := []byte("hello high value software engineer")
	hashAndBroadcast(NewHasReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHasReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

// // still works
// func hashAndBroadcast(r io.Reader) error {
// 	hash := r.(*hashReader).hash()
// 	fmt.Println(hash)

// 	return broadcast(r)
// }

func hashAndBroadcast(r HasReader) error {
	hash := r.hash()
	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes:", string(b))

	return nil
}
