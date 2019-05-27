package main

import "fmt"

// type handle int
type handle = int

func main() {
	fmt.Println("vim-go")
	var var1 = 1
	var var2 handle = 2
	types(var1)
	types(var2)

}

func types(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println(fmt.Sprintf("I am an int: %d", v))
	}
	switch v := val.(type) {
	case handle:
		fmt.Println(fmt.Sprintf("I am a handle: %d", v))
	}
}
