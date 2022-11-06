package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num
}

func main() {
	fmt.Println("what is today's lucky num?")
	c := make(chan int)
	go getLuckyNum(c)
	// time.Sleep(time.Second * 5)

	num := <-c

	fmt.Printf("Today's your lucky num is %d\n", num)

	close(c)
}
