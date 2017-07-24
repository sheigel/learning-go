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
		fmt.Printf("building and sending took: %v\n", time.Now().Sub(startTime))
	}(time.Now())

	pipelinesCount, _ := strconv.Atoi(os.Args[1])

	println("creating ", pipelinesCount, " go routines and channels")

	for ; pipelinesCount > 0; pipelinesCount-- {

		outPipe = make(chan string)
		go filter(inPipe, outPipe)
		inPipe = outPipe
	}
	println("finished building")
	defer func(startTime time.Time) {
		fmt.Printf("value received: %v after: %v\n", <-outPipe, time.Now().Sub(startTime))
	}(time.Now())

	in <- "ana are mere"

}

func filter(in <-chan string, out chan string) {
	val := <-in
	out <- val
	return
}
