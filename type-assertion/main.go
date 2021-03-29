package main

import (
	"errors"
	"fmt"
)

type BarError struct {
	e error
}

func (ba BarError) Error() string {
	return ba.e.Error()
}

type FooError struct {
	e error
}

func (fo FooError) Error() string {
	return fo.e.Error()
}

func main() {
	errorTypeA := errors.New("this is Foo errors")
	var err interface{}
	err = FooError{errorTypeA}

	switch err.(type) {
	case BarError:
		fmt.Printf("this is BarError: %+v\n", err)
	case FooError:
		fmt.Printf("this is FooError: %+v\n", err)
	default:
		fmt.Printf("%+v\n", err)
	}
}
