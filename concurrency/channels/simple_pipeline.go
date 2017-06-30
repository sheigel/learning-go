package main

import (
	"math"
	"./testp"
)

func main() {
	Printer(Squarer(Counter()))

	testp.Val(10).True()
	f := testp.Val.True
	f(10)
}

func Counter() chan int {
	generator := make(chan int)

	go func() {
		for x := 0; ; x++ {
			generator <- x
		}
	}()
	return generator
}

func Squarer(source chan int) chan int {
	generator := make(chan int)
	go func() {
		for val := range source {
			generator <- int(math.Pow(float64(val), 2))
		}
	}()
	return generator
}

func Printer(source chan int) {
	for val := range source {
		println(val)
	}
}
