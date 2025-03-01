package utils

import "testing"

type TestCase struct {
	name  string
	input any
	want  bool
}

func TestIsNil(tester *testing.T) {
	tests := []TestCase{
		{"nil value", nil, true},
		{"nil slice", ([]int)(nil), true},
		{"non-nil slice", []int{}, false},
		{"nil map", (map[string]int)(nil), true},
		{"non-nil map", map[string]int{}, false},
		{"nil pointer", (*int)(nil), true},
		{"non-nil pointer", new(int), false},
		{"non-nil int", 123, false},
	}

	for _, testCase := range tests {
		tester.Run(testCase.name, func(t *testing.T) {
			if got := IsNil(testCase.input); got != testCase.want {
				t.Errorf("IsNil() = %v, want %v", got, testCase.want)
			}
		})
	}
}
