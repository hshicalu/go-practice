package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getLuckyNum() {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky num is %d\n", num)

}

func main() {
	fmt.Println("what is today's lucky num?")
	// go getLuckyNum()
	// time.Sleep(time.Second * 5)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		getLuckyNum()
	}()
	wg.Wait()
}
