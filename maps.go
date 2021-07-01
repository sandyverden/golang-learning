// maps are Go's builtin associative data type (Sometime called hashes or dicts on other languages)

package main

import "fmt"

func main() {

	// To carete an empty map, use builtin make:
	// make (map[key-type] val-type)

	m := make(map[string]int)

	// set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs
	fmt.Println("map:", m)

	// Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len: ", len(m))

	// the builtin delete removes key/value pair from a map
	delete(m, "k2")
	fmt.Println("map: ", m)

	// The optional second return value when getting a value
	// from a map indicate if the key was present in the map.
	// This can be use to disambiguate between missing keys
	// and keys with zero value like 0 or "". Here we didn't need
	// value itself, so we ignore it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Note that maps apprear in the form map[k:v k:v] when printed with fmt.Println
}
