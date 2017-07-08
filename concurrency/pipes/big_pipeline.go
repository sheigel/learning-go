package main

import (
	"os"
	"strconv"
	"time"
	"fmt"
)

func main() {
	in := make(chan string)
	inPipe := in

	var outPipe chan string

	defer func(startTime time.Time) {
		fmt.Printf("value received: %v after: %v", <-outPipe, time.Now().Sub(startTime))
	}(time.Now())

	pipelinesCount, _ := strconv.Atoi(os.Args[1])

	println("creating ", pipelinesCount, " go routines and channels")

	for ; pipelinesCount > 0; pipelinesCount-- {
		println("building #", pipelinesCount)

		outPipe = make(chan string)
		go filter(inPipe, outPipe)
		inPipe = outPipe
	}

	in <- "ana are mere"

}

func filter(in <-chan string, out chan string) {
	val := <-in
	out <- val
	return
}
