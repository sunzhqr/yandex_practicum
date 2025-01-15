package sum

func Sum(values ...int) int {
	var sum int
	for _, value := range values {
		sum += value
	}
	return sum
}
