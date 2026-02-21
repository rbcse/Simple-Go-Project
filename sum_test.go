package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_SumOf2Numbers(t *testing.T) {
	testCases := []struct {
		description	string
		num1 float64
		num2 float64
		expectedSum float64
	}{
		{
			description: "Should return 3 when num1 is 1 and num2 is 2",
			num1: 1,
			num2 : 2,
			expectedSum: 3,
		},
		{
			description: "Should return 2 when num1 is 2 and num2 is 0",
			num1: 2,
			num2 : 0,
			expectedSum: 2,
		},
		{
			description: "Should return 2 when num1 is 4 and num2 is -2",
			num1: 4,
			num2 : -2,
			expectedSum: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			sum := Sum(tC.num1,tC.num2)
			assert.Equal(t,tC.expectedSum,sum);
		})
	}
}