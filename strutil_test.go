package strutil_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/salawhaaat/strutil"
)

func TestReverse(t *testing.T) {
	// Test cases
	tests := map[string]struct {
		input    string
		expected string
	}{
		"empty":      {input: "", expected: ""},
		"single":     {input: "a", expected: "a"},
		"hello":      {input: "hello", expected: "olleh"},
		"unicode":    {input: "Hello, 世界", expected: "界世 ,olleH"},
		"palindrome": {input: "racecar", expected: "racecar"},
		"number":     {input: "1234567890", expected: "0987654321"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := strutil.Reverse(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestSymbolCount(t *testing.T) {
	// Test cases
	tests := map[string]struct {
		input    string
		expected int
	}{
		"empty":      {input: "", expected: 0},
		"single":     {input: "a", expected: 1},
		"hello":      {input: "hello", expected: 5},
		"unicode":    {input: "Hello, 世界", expected: 9},
		"palindrome": {input: "racecar", expected: 7},
		"number":     {input: "1234567890", expected: 10},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := strutil.SymbolCount(tc.input)
			if got != tc.expected {
				t.Errorf("SymbolCount(%s) = %d, expected %d", tc.input, got, tc.expected)
			}
		})
	}
}

func ExampleReverse() {
	fmt.Println(strutil.Reverse("hello"))
	// Output:
	// olleh
}

func ExampleSymbolCount() {
	fmt.Println(strutil.SymbolCount("Hello, 世界"))
	// Output:
	// 9
}
