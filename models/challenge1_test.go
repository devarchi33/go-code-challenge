package models_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func longFunc() time.Duration {
	start := time.Now()
	duration := time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5))
	fmt.Printf("Execute Time: %#v\n", time.Second*duration)
	time.Sleep(time.Second * duration)
	return time.Now().Sub(start)
}

func TestChallenge1(t *testing.T) {
	t.Run("challenge1", func(t *testing.T) {
		channel1 := make(chan time.Duration, 1)
		go func() {
			channel1 <- longFunc()
		}()

		select {
		case res := <-channel1:
			fmt.Printf("Execute time is less than 2 second, time = %#v\n", res)
		case <-time.After(time.Second * 2):
			fmt.Println("Execute time is more than 2 second")
		}
	})
}
