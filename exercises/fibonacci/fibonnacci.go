package main

import "fmt"

// fibonacci returns a closure function that generates Fibonacci numbers.
// Each time the closure function is called, it returns the next Fibonacci number in the sequence.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a = b
		b = result + b // I don't use a here because it's already updated and breaks the logic
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
