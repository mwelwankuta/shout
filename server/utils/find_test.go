package utils

import "testing"

type Test struct {
	array    []uint
	value    uint
	expected bool
}

var testCases = []Test{
	{[]uint{0, 1, 9, 3}, 2, false},
	{[]uint{0, 4, 1, 5}, 4, true},
	{[]uint{1, 2, 8, 3}, 2, true},
	{[]uint{20, 22, 18, 23}, 5, false},
}

func TestFind(t *testing.T) {
	for _, test := range testCases {
		if Find(test.array, test.value) != test.expected {
			t.Fatalf("Array: %v Value %v Expected: %v", test.array, test.value, test.expected)
		}
	}
}
