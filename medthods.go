// Go support methods defined on struct type

package main

import "fmt"

type rect struct {
	width, height int
}

// This are method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Method can be defined for either pointer or value receiver type.
// Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// here we call the 2 methods defined for our struct
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handle conversion between value and pointers
	// for method call. You may want to use a pointer receiver type to avoid copying
	// on method call or to allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Print("perim:", rp.perim())

}
