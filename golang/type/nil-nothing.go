package main

import "fmt"

// MyError 实现了error接口
type MyError string

func (m MyError) Error() string {
	return string(m)
}

func main() {
	fmt.Println(nil == test(true))
	fmt.Println(nil == test(false))

}

func test(v bool) error {
	var e *MyError = nil
	if e == nil {
		fmt.Println("e == nil")
	}
	// var e MyError
	if v {
		return nil
	}
	return e
}
