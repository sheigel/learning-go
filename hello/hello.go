package main

import "time"

var a string

func f(asdf string) {
	print(a)
	print(asdf)
}

func hello() {
	a = "hello, world"
	go f(a)
	a="haaaa"
}


func main() {

	hello()

	time.Sleep(3*time.Second)
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
