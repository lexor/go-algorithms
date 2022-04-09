package primefactors

import (
	"reflect"
	"testing"
)

var cases = []struct {
	number   int
	expected []int
}{
	{number: 1, expected: []int{}},
	{number: 2, expected: []int{2}},
	{number: 3, expected: []int{3}},
	{number: 4, expected: []int{2, 2}},
	{number: 14, expected: []int{2, 7}},
	{number: 40, expected: []int{2, 2, 2, 5}},
	{number: 54, expected: []int{2, 3, 3, 3}},
	{number: 100, expected: []int{2, 2, 5, 5}},
	{number: 156, expected: []int{2, 2, 3, 13}},
	{number: 273, expected: []int{3, 7, 13}},
	{number: 300, expected: []int{2, 2, 3, 5, 5}},
	{number: 980, expected: []int{2, 2, 5, 7, 7}},
	{number: 1000, expected: []int{2, 2, 2, 5, 5, 5}},
	{number: 52734, expected: []int{2, 3, 11, 17, 47}},
	{number: 343434, expected: []int{2, 3, 7, 13, 17, 37}},
	{number: 456745, expected: []int{5, 167, 547}},
	{number: 510510, expected: []int{2, 3, 5, 7, 11, 13, 17}},
	{number: 8735463, expected: []int{3, 3, 11, 88237}},
	{number: 873452453, expected: []int{149, 1637, 3581}},
}

func TestPrimeFactors(t *testing.T) {
	for _, input := range cases {
		result := primeFactors(input.number)

		if !reflect.DeepEqual(result, input.expected) {
			t.Errorf("FAIL: number=%v, expected=%v, result=%v", input.number, input.expected, result)
		}
	}
}
