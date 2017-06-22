package main

import (
	"encoding/json"
	"strings"
	"github.com/kelseyhightower/confd/log"
	"fmt"
)

type Animal int

func (a Animal) MarshalJSON() ([]byte, error) {
	switch a {
	default:
		return json.Marshal("unknown")
	case Gopher:
		return json.Marshal("gopher")
	case Zebra:
		return json.Marshal("zebra")
	}
}

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}
	return nil
}

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func main() {
	jsonStr := `["gopher","armadillo","zebra", "unknown", "zebra","gopher", "zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(jsonStr), &zoo); err != nil {
		log.Fatal("bad stuff: %v", err)
	}

	census := make(map[Animal]int)

	for _, animal := range zoo {
		census[animal]++
	}
	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}
