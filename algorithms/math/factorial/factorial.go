package factorial

func factorial(number int) int {
	result := 1

	for i := 2; i <= number; i++ {
		result *= i
	}

	return result
}

func recursiveFactorial(number int) int {
	if number > 1 {
		return number * recursiveFactorial(number-1)
	} else {
		return 1
	}
}
