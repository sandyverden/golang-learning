// In Go it's idiomatic to communicate errors via an explicit,
// separate return value. This contrasts with the exceptions
// used in languages like Java and Ruby and the overloaded
// single result/ error value sometimes used in C. Go's
// approach makes it easy to see which functions return
// errors and tohandle them using the same language
// construct employed for any other, non-error tasks.

package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have
// type error, a built-in interfaces.
func f1(arg int) (int, error) {
	if arg == 42 {

		// error.New constructs a basic error value with the given error message
		return -1, errors.New("Can't work at 42")
	}

	// A nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// It's possible to use costum type are errors by
// implementing the Error() method on them. Here's a
// variant on the example above that use a custom type to
// explicitly represent an argument error.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// in this case we use &argError syntax to build a new struct,
		// supplying values for the two fields arg and prob
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// The two loops below test out each of our error-returning functions.
	// Note that the use of an inline error check on the
	// if line is a common idiom in Go Code
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
