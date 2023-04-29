package gopkgexample

func Add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func Multiplay(args ...int) int {
	product := 1
	for _, arg := range args {
		product = product * arg
	}
	return product
}

