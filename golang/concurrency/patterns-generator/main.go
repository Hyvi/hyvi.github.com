package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := boring("boring!") // Function returning a channel
	for i := 0; i < 5; i++ {

		fmt.Printf("You say: %q\n", <-c) // 读管道
	}
	fmt.Println("You're boring; I'm leaving.")

}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // Web launch the gorountine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller

}

// main线程等待c Channel数据，并阻塞。
// boring func 起了新goroutine，无限循环并向c Channel中发送信息
