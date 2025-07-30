package main

import (
	"fmt"
	"testing"
)

func TestExpandToCombinations(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name:     "empty input",
			input:    [][]string{},
			expected: [][]string{{}},
		},
		{
			name:     "single component",
			input:    [][]string{{"A", "B"}},
			expected: [][]string{{"A"}, {"B"}},
		},
		{
			name:     "two components",
			input:    [][]string{{"imgA", "imgB"}, {"txtX", "txtY"}},
			expected: [][]string{{"imgA", "txtX"}, {"imgB", "txtX"}, {"imgA", "txtY"}, {"imgB", "txtY"}},
		},
		{
			name:  "three components",
			input: [][]string{{"1", "2"}, {"A", "B"}, {"X", "Y"}},
			// 用纸笔来算，可以算出来：约定层的间隙越多
			expected: [][]string{{"1", "A", "X"}, {"2", "A", "X"}, {"1", "B", "X"}, {"2", "B", "X"}, {"1", "A", "Y"}, {"2", "A", "Y"}, {"1", "B", "Y"}, {"2", "B", "Y"}},
		},
		{
			name:  "four components",
			input: [][]string{{"1", "2"}, {"A", "B"}, {"X", "Y"}, {"i", "j"}},
			// 用纸笔来算，可以算出来：越顶层的间隙越多，比如1和2是最开始的数组，那么在输出结果里是1,2,1,2 的输出，而A,B 是 A, A, B, B 的输出，每一层的连续次数是这一层层数减一的2幂数
			expected: [][]string{{"1", "A", "X", "i"}, {"2", "A", "X", "i"}, {"1", "B", "X", "i"}, {"2", "B", "X", "i"}, {"1", "A", "Y", "i"}, {"2", "A", "Y", "i"}, {"1", "B", "Y", "i"}, {"2", "B", "Y", "i"}, {"1", "A", "X", "j"}, {"2", "A", "X", "j"}, {"1", "B", "X", "j"}, {"2", "B", "X", "j"}, {"1", "A", "Y", "j"}, {"2", "A", "Y", "j"}, {"1", "B", "Y", "j"}, {"2", "B", "Y", "j"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := expandToCombinations(tt.input)
			fmt.Println(result)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d combinations, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if !equal(result[i], tt.expected[i]) {
					t.Errorf("combination %d mismatch: expected %v, got %v", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestExpandToCombinations2(t *testing.T) {
	componentLists := [][]string{
		{"imgA", "imgB"},
		{"txtX", "txtY"},
	}

	results := expandToCombinations(componentLists)
	fmt.Println(results)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
