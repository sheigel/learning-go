package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"sort"
)
type T struct{
	str string
}
type T1 struct{
	a int
	t T
}
type T2 struct{
	a int
	t T
}
func main() {
sort.Reverse()
	t1:=T1{a:10, t:T{"text"}}
	var t2 T2 = T2{a: 10, t: T{str:"text1"}}

	fmt.Printf("asdf%v",t2)
	fmt.Printf("%t", (interface{})(t1)==(interface{})(T1(t2)))
	return
	if len(os.Args) < 1 {
		os.Exit(1)
	}

	fileName := os.Args[1]
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	fileTextLines := strings.Split(string(fileBytes), "\n")
	fileTextLines = fileTextLines[:len(fileTextLines)-1]
	for idx, line := range fileTextLines {
		println("idx:", idx, "->", line)
	}
	var fl float32=10
	var in int32 = (int32)(fl)

	println(in)
}
