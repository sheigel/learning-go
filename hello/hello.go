package main

func main() {
	source := make(chan int)
	completed := make(chan interface{})
	go receiver(source, completed)
	source <- 10
	source <- 11
	source <- 12
	source <- 13
	close(source)

	 <- completed
	println("done waiting")
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
