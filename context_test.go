package context

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

type Response struct {
	Value int
	Err   error
}

func FetchUserData(ctx context.Context, userID int) (int, error) {
	value := ctx.Value("you")
	fmt.Println(value)

	context, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	respch := make(chan Response)

	go func() {
		val, err := FetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			Value: val,
			Err:   err,
		}
	}()
	for {
		select {
		case <-context.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-respch:
			return resp.Value, resp.Err
		}
	}
}

func FetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100)

	return 666, nil
}

func TestFetchingThirdPartyWithContext(t *testing.T) {
	start := time.Now()

	ctx := context.WithValue(context.Background(), "you", "who")
	userID := 10
	val, err := FetchUserData(ctx, userID)
	if err != nil {
		log.Fatal("fething from third party took to long")
	}
	fmt.Println("value:", val)
	fmt.Println("took:", time.Since(start))
}
