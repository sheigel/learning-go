package fibonacio

import (
	"fmt"
)

func main() {
	var val int
	var err error

	for ; err == nil; {
		fmt.Scanf("%d", &val)

		fmt.Println(fmt.Sprintf("Fibonacci for %d is %d", val, fibonacci((val))))

	}
}
func fibonacci(idx int) int {
	if idx < 1 {
		return 1

	}

	return fibonacci(idx-1) + fibonacci(idx-2)
}

type Fibo struct {
	value int
	Text  string
}

func NewFibo() *Fibo {
	return &Fibo{
		value: 10,
		Text:  "asdf",
	}
}
