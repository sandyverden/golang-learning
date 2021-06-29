package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "contants"

func main() {
	fmt.Println(s)

	// A conts statement can appear anyware a var statement can
	const n = 50000000

	// Comstant espressions perform arithmatic with arbitary percision
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type untul it's given one, such as by an explicit conversion
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignement or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}
