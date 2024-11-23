package sum

func Plus(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}
