package main

// import (
// 	"crypto/sha1"
// 	"encoding/hex"
// 	"fmt"
// 	"io"
// )

// func hashAndBroadcastBefore(r io.Reader) error {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	hash := sha1.Sum(b)
// 	fmt.Println(hex.EncodeToString(hash[:]))

// 	return broadcastBefore(r)
// }

// func broadcastBefore(r io.Reader) error {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("string of the bytes:", string(b))

// 	return nil
// }
