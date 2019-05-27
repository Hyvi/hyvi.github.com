package main

import "fmt"

// Callable new TypeSpec for func
type Callable func()

func main() {
	callable1 := func() {
		fmt.Println("callable 1")
	}
	var callable2 Callable
	callable2 = func() {
		fmt.Println("callable 2")
	}

	test(callable1)
	test(callable2)
	// callable2 被显式声明为Callable. Callable 是一个独立的TypeSpec， 因此与func()的类型不一样。
}

func test(val interface{}) {
	switch v := val.(type) {
	case func():
		v()
	default:
		fmt.Println("wrong type")
	}
}
