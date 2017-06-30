package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"log"
	"os"
)

func main() {
	workList := make(chan []string)
	go func() {
		workList <- os.Args[1:]
	}()

	availableSlots := make(chan struct{}, 20)
	seen := make(map[string]bool)
	for list := range workList {
		for _, url := range list {
			if seen[url] {
				continue
			}
			seen[url] = true

			go func(url string) {
				availableSlots <- struct{}{}
				workList <- crawl(url)
				<-availableSlots
			}(url)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
