package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   test one  ",
			expected: []string{"test", "one"},
		},
		{
			input:    "Test TwO",
			expected: []string{"test", "two"},
		},
		{
			input:    "Test    three",
			expected: []string{"test", "three"},
		},
		{
			input:    "test more four",
			expected: []string{"test", "more", "four"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("expected: %v, got: %v", c.expected, actual)
		}
	}

}
