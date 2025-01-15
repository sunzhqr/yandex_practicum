package sum

func Sum(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func Mult(values ...int) int {
	var sum = 1
	for _, v := range values {
		sum *= v
	}
	return sum
}
