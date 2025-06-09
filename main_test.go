package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"zero size", 0, 0},
		{"single element", 1, 1},
		{"small slice", 100, 100},
		{"large slice", SIZE, SIZE},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice := generateRandomElements(tt.size)
			assert.Equal(t, tt.wantLen, len(slice))
			if tt.size > 0 {
				assert.NotEqual(t, slice, generateRandomElements(tt.size))
			} else {
				assert.Nil(t, slice)
			}

		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single number", []int{42}, 42},
		{"all positive", []int{1, 5, 3, 9, 2}, 9},
		{"with negatives", []int{-1, -5, -3}, -1},
		{"mixed values", []int{-1, 0, 1, -100}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximum(tt.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single number", []int{42}, 42},
		{"numbers = chunks", []int{1, 2, 3, 4, 5, 6, 7, 0}, 7},
		{"more than chunks", []int{1, 5, 3, 9, 2, 8, 4, 6, 7, 0}, 9},
		{"all negatives", []int{-1, -5, -3, -9}, -1},
		{"mixed values", []int{-1, 0, 1, -100, 42}, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxChunks(tt.data))
		})
	}
}
