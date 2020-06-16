package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const maxDuration = time.Second * 2

func main() {
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT)

	ctx, cancel := context.WithTimeout(context.Background(), maxDuration)

	go func() {
		fmt.Println("Signal:", <-exit)
		cancel()
	}()

	result, err := longFuncWithCtx(ctx)
	fmt.Println(result)
	if err != nil {
		log.Fatal(err)
	}

}

func longFuncWithCtx(ctx context.Context) (string, error) {
	done := make(chan string)

	go func() {
		done <- fmt.Sprint(longFunc())
	}()

	select {
	case <-ctx.Done():
		return "Fail", ctx.Err()
	case result := <-done:
		return result, nil
	}
}
func longFunc() time.Duration {
	start := time.Now()
	duration := time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5))
	time.Sleep(time.Second * duration)
	return time.Now().Sub(start)
}
