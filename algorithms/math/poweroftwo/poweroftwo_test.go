package poweroftwo

import "testing"

var cases = []struct {
	number   int
	expected bool
}{
	{
		number:   -1,
		expected: false,
	},
	{
		number:   0,
		expected: false,
	},
	{
		number:   1,
		expected: true,
	},
	{
		number:   2,
		expected: true,
	},
	{
		number:   3,
		expected: false,
	},
	{
		number:   4,
		expected: true,
	},
	{
		number:   5,
		expected: false,
	},
	{
		number:   6,
		expected: false,
	},
	{
		number:   7,
		expected: false,
	},
	{
		number:   8,
		expected: true,
	},
	{
		number:   10,
		expected: false,
	},
	{
		number:   12,
		expected: false,
	},
	{
		number:   16,
		expected: true,
	},
	{
		number:   31,
		expected: false,
	},
	{
		number:   64,
		expected: true,
	},
	{
		number:   1024,
		expected: true,
	},
	{
		number:   1023,
		expected: false,
	},
}

func TestPowerOfTwo(t *testing.T) {
	for _, input := range cases {
		result := isPowerOfTwo(input.number)

		if result != input.expected {
			t.Errorf("FAIL: number=%v, expected=%v, result=%v", input.number, input.expected, result)
		}
	}
}
