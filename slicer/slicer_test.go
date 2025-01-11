package slicer_test

import (
	"github.com/WantBeASleep/goooool/slicer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten2DArray_Success_DataSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "empty 2D array",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "one row",
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "multiple rows",
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "mixed lengths",
			input:    [][]int{{1}, {2, 3}, {}, {4, 5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "single empty row",
			input:    [][]int{{}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := slicer.Flatten2DArray(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
