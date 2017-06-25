package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go f1(&wg)

	wg.Add(1)
	go f2(&wg)

	wg.Wait()
}

func f1(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		println("f1:", i)
	}
	println("Done f1")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		println("f2:", i)
	}
	println("Done f2")
	wg.Done()
}
