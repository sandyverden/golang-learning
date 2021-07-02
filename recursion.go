// Go support recursive functions, Here's a classic factorial example

package main

import "fmt"

// This fact function call itself untul it reaches the base case of fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(4))
}
