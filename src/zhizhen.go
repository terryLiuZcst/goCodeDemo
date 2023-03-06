package src

func AddOne(x int) int {
	x++
	return x
}
func AddOneDemo(x *int) {
	*x++
	//return *x
}

type Demo struct {
	Tag string
}

func SetDemo(demo *Demo, s string) {
	(*demo).Tag = s
}
