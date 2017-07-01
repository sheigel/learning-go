package main

import (
	"fmt"
	errs "github.com/pkg/errors"
	"errors"
)

func main() {
	fmt.Printf("direct error: %+v\n", createSimpleError())
	println()
	fmt.Printf("returned error: %+v\n", returnSimpleError())
	println()
	fmt.Printf("wrapped error: %+v\n", wrapperOfCreateError())
	println()
	fmt.Printf("direct stack: %+v\n", stack())
	println()
	fmt.Printf("wrapped stack: %+v\n", stackWrapper())
}

func stackWrapper() error {
	return errs.WithStack(stack())
}
func stack() error {
	return errs.Errorf("should have stack")
}
func wrapperOfCreateError() error {
	return errs.Wrap(createSimpleError(), "wrapper")
}
func returnSimpleError() error {
	return createSimpleError()
}

func createSimpleError() error {
	return errors.New("the simple error")
}
