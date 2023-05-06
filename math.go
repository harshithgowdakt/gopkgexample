package gopkgexample

// Add function adds all the arguments and returns sum
func Add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// Multiply function multiplies all the arguments and returns product
func Multiply(args ...int) int {
	product := 1
	for _, arg := range args {
		product = product * arg
	}
	return product
}

