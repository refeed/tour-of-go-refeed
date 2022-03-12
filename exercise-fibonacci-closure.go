package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib := []int{0, 1}
	return func() int {
		ret := fib[0]
		fib[0], fib[1] = fib[1], fib[0] + fib[1]
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
