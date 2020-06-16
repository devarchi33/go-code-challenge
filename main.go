package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	elapsedTime := longFunc()
	fmt.Println(elapsedTime)

}
func longFunc() time.Duration {
	start := time.Now()
	duration := time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5))
	time.Sleep(time.Second * duration)
	return time.Now().Sub(start)
}
