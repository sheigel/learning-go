package main

import (
	"math"
	"fmt"
)

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	c.radius+=1
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	a:=[]circle{c}
	//info(c)
	println(a[0].area())
	println(a[0].area())
	println(a[0].area())
}
