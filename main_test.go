package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	slice := generateRandomElements(1)
	assert.Len(t, slice, 1)

	assert.Nil(t, generateRandomElements(0))

	slice = generateRandomElements(SIZE)
	assert.Len(t, slice, SIZE)

	slice = generateRandomElements(100)
	slice2 := generateRandomElements(100)
	assert.NotEqual(t, slice, slice2) //Провека действительно ли числа каждый раз разные
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
	data := []int{1, 5, 3, 9, 2, 8, 4, 6, 7, 0}

	t.Run("empty slice", func(t *testing.T) {
		assert.Equal(t, 0, maxChunks([]int{}))
	})

	t.Run("single number", func(t *testing.T) {
		assert.Equal(t, 42, maxChunks([]int{42}))
	})

	t.Run("numbers = chunks", func(t *testing.T) {
		assert.Equal(t, 7, maxChunks([]int{1, 2, 3, 4, 5, 6, 7, 0}))
	})

	t.Run("more numbers than chunks", func(t *testing.T) {
		assert.Equal(t, 9, maxChunks(data))
	})
}
