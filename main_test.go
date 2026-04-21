package main

import "testing"

type Test struct {
	name     string
	a, b     int
	expected int
}

func TestAdd(t *testing.T) {
	var test1 = Test{"add positive numbers", 5, 7, 12}
	var test2 = Test{"adding negative", -3, -4, -7}
	var test3 = Test{"adding positive and negative", 5, -7, -2}
	var test4 = Test{"adding positive and zero", 5, 0, 5}
	var test5 = Test{"adding negative and zero", -3, 0, -3}
	var test6 = Test{"adding zero and zero", 0, 0, 0}

	var tests = []Test{test1, test2, test3, test4, test5, test6}

	for _, tableTest := range tests {
		t.Run(tableTest.name, func(t *testing.T) {
			result := add(tableTest.a, tableTest.b)
			if result != tableTest.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tableTest.a, tableTest.b, result, tableTest.expected)
			}
		})
	}
}
