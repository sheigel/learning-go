package main

import (
	"fmt"
)

var a string

func f(asdf string) {
	print(a)
	print(asdf)
}

func hello() {
	a = "hello, world"
	go f(a)
	a = "haaaa"
}

type str string

func (s *str) Write(p []byte) (n int, err error) {
	return 10, nil
}

func (s str) Read(p []byte) (n int, err error) {
	return 11, nil
}

func (s str) String() string {
	return fmt.Sprintf("test:%v", string(s))
}

func main() {
	//test(str("direct string"))
	//test(nil)
	//test(new(str))
	//var s str
	//test(s)
	//s := str("asdfk")
	rw := str("asdf")

	var val int
	val, _ = (&rw).Write([]byte("qqqq"))
	println(val)
	val, _ = (&rw).Read([]byte("qqqq"))
	println(val)
}

func test(s fmt.Stringer) {
	println("is nil", s == nil)
	if s == nil {
		return
	}
	println("string is", s.String())
}

func receiver(source chan int, completed chan interface{}) {
	for {
		select {
		case val, ok := <-source:
			if ok {
				println("the ok:", ok, "the value:", val)
			} else {
				completed <- nil

			}
		}

	}
}
