package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		result := sum1
		sum1, sum2 = sum2, sum1+sum2
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
