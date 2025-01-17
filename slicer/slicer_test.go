package slicer_test

import (
	"testing"

	"github.com/WantBeASleep/med_ml_lib/slicer"
	"github.com/stretchr/testify/assert"
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

func TestPackSlice_Success_DataSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int
		expected []any
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []any{},
		},
		{
			name:     "single element",
			input:    []int{42},
			expected: []any{42},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3},
			expected: []any{1, 2, 3},
		},
		{
			name:     "negative integers",
			input:    []int{-1, -2, -3},
			expected: []any{-1, -2, -3},
		},
		{
			name:     "mixed integers",
			input:    []int{0, 10, 20},
			expected: []any{0, 10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := slicer.PackSlice(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
