package fibonacci

func fibonacci(n int) []int {
	fibSequence := []int{1}

	currentValue := 1
	previousValue := 0

	if n == 1 {
		return fibSequence
	}

	iterationsCounter := n - 1

	for iterationsCounter > 0 {
		currentValue += previousValue
		previousValue = currentValue - previousValue

		fibSequence = append(fibSequence, currentValue)

		iterationsCounter -= 1
	}

	return fibSequence
}
