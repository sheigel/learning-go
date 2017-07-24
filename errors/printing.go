package main

import (
	"fmt"
	errs "github.com/pkg/errors"
)

func main() {
	fmt.Printf("direct error: %+v\n", createSimpleError())
	println()
	fmt.Printf("returned error: %+v\n", returnSimpleError())
	println()
	fmt.Printf("wrapped error: %+v\n", wrappers())
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
func wrappers() error {
	return errs.Wrap(wrapper3(), "wrapper")
}
func wrapper3() error {
	return errs.Wrap(wrapper2(), "wrapper")
}
func wrapper2() error {
	return errs.Wrap(wrapper1(), "wrapper")
}
func wrapper1() error {
	return errs.Wrap(createSimpleError(), "wrapper")
}
func returnSimpleError() error {
	return createSimpleError()
}

func createSimpleError() error {
	return errs.New("the simple error")
}
