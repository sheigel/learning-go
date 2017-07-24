package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type person struct {
	FirstName, LastName string
	Age                 int
	Address             string
}

func main() {
	me := person{FirstName: "silviu", LastName: "eigel", Age: 34, Address: "Romania"}
	fmt.Printf("Me: %v", me)

	fmt.Printf("Me encoded: ")
	json.NewEncoder(os.Stdout).Encode(me)

	jsonBytes, _ := json.Marshal(me)
	fmt.Printf("Me marshalled: %v", string(jsonBytes))
}
