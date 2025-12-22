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
	// ваш код здесь
	if size < 1 {
		return []int{}
	}
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(SIZE)
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	s := data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > s {
			s = data[i]
		}
	}
	return s
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	ch := make(chan int, CHUNKS)
	var wg sync.WaitGroup
	chunkSize := len(data) / CHUNKS

	for i := 1; i < CHUNKS; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}
		if start > len(data) {
			wg.Done()
			continue
		}
		go func(start int, end int) {
			defer wg.Done()
			chunkMax := maximum(data[start:end])
			ch <- chunkMax
		}(start, end)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	maxVal, ok := <-ch
	if !ok {
		return 0
	}

	for val := range ch {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	linearStart := time.Now()
	max := maximum(data)
	elapsed := time.Since(linearStart).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	paralellStart := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(paralellStart).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
