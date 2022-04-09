package radian

import (
	"math"
	"testing"
)

var cases = []struct {
	degree float64
	radian float64
}{
	{
		degree: 0,
		radian: 0,
	},
	{
		degree: 45,
		radian: math.Pi / 4,
	},
	{
		degree: 90,
		radian: math.Pi / 2,
	},
	{
		degree: 180,
		radian: math.Pi,
	},
	{
		degree: 270,
		radian: (3 * math.Pi) / 2,
	},
	{
		degree: 360,
		radian: 2 * math.Pi,
	},
}

func TestDegreeToRadian(t *testing.T) {
	for _, input := range cases {
		result := degreeToRadian(input.degree)

		if input.radian != result {
			t.Errorf("FAIL: degree=%v expected=%v, result=%v", input.degree, input.radian, result)
		}
	}
}

func TestRadianToDegree(t *testing.T) {
	for _, input := range cases {
		result := radianToDegree(input.radian)

		if input.degree != result {
			t.Errorf("FAIL: radian=%v expected=%v, result=%v", input.radian, input.degree, result)
		}
	}
}
