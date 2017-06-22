package init_example

var Called bool

func init() {
	println("init_example was Called")
	Called =true
}
