package main

func main() {
	result :=deferrer_sets_parameter(10)
	println("parameter:", result)
	result =deferrer_sets_named_returns()
	println("named return::", result)
}
func deferrer_sets_parameter(i int)  int {
	defer func() {
		i++
	}()
	println("value is ", i)
	return  i
}

func deferrer_sets_named_returns()  (i int) {
	defer func() {
		i=12
	}()
	println("value is ", i)
	return 1
}