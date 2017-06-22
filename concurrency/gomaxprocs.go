package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		println("f1")
	}()

	go func() {
		println("f2")
	}()

	result := 0
	for i := 0; i < 10000; i++ {
		result = +i
	}
	println("result", result, runtime.GOMAXPROCS(1))
}
