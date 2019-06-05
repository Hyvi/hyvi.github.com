package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	go func() {

		go func() {
			time.Sleep(time.Second * 4)
			fmt.Println("go 2 over")
		}()
		time.Sleep(time.Second * 2)

		done <- 1
		fmt.Println("go 1 over")
	}()

	<-done

	time.Sleep(time.Second * 10)

}
