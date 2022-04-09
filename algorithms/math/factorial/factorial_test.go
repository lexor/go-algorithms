package factorial

import (
	"testing"
)

var cases = []struct {
	number   int
	expected int
}{
	{
		number:   1,
		expected: 1,
	},
	{
		number:   2,
		expected: 2,
	},
	{
		number:   3,
		expected: 6,
	},
	{
		number:   4,
		expected: 24,
	},
}

func TestFactorial(t *testing.T) {
	for _, input := range cases {
		result := factorial(input.number)

		if input.expected != result {
			t.Errorf("FAIL: expected=%d, result=%d", input.expected, result)
		}
	}
}

func TestRecursiveFactorial(t *testing.T) {
	for _, input := range cases {
		result := recursiveFactorial(input.number)

		if input.expected != result {
			t.Errorf("FAIL: expected=%d, result=%d", input.expected, result)
		}
	}
}
