package primefactors

import "math"

func primeFactors(n int) []int {
	nn := n
	factors := []int{}

	for factor := 2; factor <= int(math.Sqrt(float64(nn))); factor++ {
		// Check that factor divides n without a reminder.
		for nn%factor == 0 {
			// Overriding the value of n.
			nn /= factor

			// Saving the factor.
			factors = append(factors, factor)
		}
	}

	// The ultimate reminder should be a last prime factor,
	// unless it is not 1 (since 1 is not a prime number).
	if nn != 1 {
		factors = append(factors, nn)
	}

	return factors
}
