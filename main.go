package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	src := rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(1000)))
	if size < 1 {
		return nil
	}
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = int(src.Int63())
	}
	//	fmt.Println(numbers)
	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	maxValue := data[0]
	for _, v := range data {
		if maxValue < v {
			maxValue = v
		}
	}
	return maxValue

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	if len(data) < CHUNKS {
		maxValue := maximum(data)
		return maxValue
	}
	chunkSize := len(data) / CHUNKS
	if len(data)%CHUNKS != 0 {
		chunkSize++
	}

	var wg sync.WaitGroup
	resultSlice := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		if start >= end {
			continue
		}

		wg.Add(1)
		go func(i int, chunk []int) {
			defer wg.Done()
			resultSlice[i] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()

	maxValue := maximum(resultSlice)
	return maxValue
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	numbers := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
