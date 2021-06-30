// Slice are a key data type in Go, giving a more powerful interface to sequences the arrays.

package main

import "fmt"

func main() {

	// Unlike arrays, slices are typed only by the elements they contain
	// (not the number of elements). To create an empty slice with non-zero lenght, use the builtin make,
	// Here we make a slice of strings of lenght 3 (initially zero-valued)
	s := make([]string, 3)
	fmt.Println("emp", s)

	// We can get and set just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// In additional to these logic operations, slice support several more
	// that make them richer than arrays, One is the
	// builtin append, which rerturns a slice containing one or more new values.
	// Note that we need to accept a return
	// value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slices can aksi ve copy'd. Here we create an emprty slice c of the same lenght as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice support a "slice" operator with the syntax
	// slice [low:high]. for example, this get a slice of the elemtns s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// we can declare and initialize a variable for slice on a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slice can be composed into multi-dimensional data structure
	// The lenght of the inner slices can very, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
