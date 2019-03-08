package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 在默认不带缓冲的channel中，每一个发送者与接受者都会阻塞当前线程，
	// 只有当接受者接收者与发送者都准备就绪了，channel才能正常使用
	// all goroutines are asleep - deadlock
	// 发送者发送到channel中，但是此时阻塞了当前线程，而没有接收者，处于一直等待的状态
	jobs := make(chan int)
	// jobs := make(chan int, 10)
	results := make(chan int)
	// results := make(chan int, 10)

	defer close(jobs)
	defer close(results)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 如果使用goroutine，chan是不需要缓存
	go func() {
		for j := 1; j <= 9; j++ {
			jobs <- j
			fmt.Println("write jobs", j)
		}
	}()

	for a := 1; a <= 9; a++ {
		<-results
		fmt.Println("read results", a)
	}
}

// 参考
// https://www.cnblogs.com/mstmdev/p/5454945.html
// https://wetest.qq.com/lab/view/346.html?from=content_testerhome
