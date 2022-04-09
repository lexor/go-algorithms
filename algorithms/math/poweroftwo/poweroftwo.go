package poweroftwo

func isPowerOfTwo(number int) bool {
	// 1 (2^0) is the smallest power of two.
	if number < 1 {
		return false
	}

	dividedNumber := number
	for dividedNumber != 1 {
		if dividedNumber%2 != 0 {
			// For every case when remainder isn't zero we can say that this number
			// couldn't be a result of power of two.
			return false
		}

		dividedNumber /= 2
	}

	return true
}
