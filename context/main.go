package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type key string

func main() {
	start := time.Now()

	ctx := context.WithValue(context.Background(), key("you"), "who") // Use the defined context key
	userID := 10
	v, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal("fetching from third took to long")
	}
	fmt.Println("value:", v)
	fmt.Println("took:", time.Since(start))

}

type response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {

	contextWithTimeout, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	respch := make(chan response)

	go func() {
		val, err := fetchThirdPartyWhichCanBeSlow()
		respch <- response{
			value: val,
			err:   err,
		}
	}()
	for {
		select {
		case <-contextWithTimeout.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyWhichCanBeSlow() (int, error) {
	time.Sleep(100 * time.Millisecond)
	return 666, nil
}
