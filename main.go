package main

import (
	"fmt"
	"math"
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
	if size <= 0 {
		return []int{}
	}

	numbers := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		numbers[i] = r.Intn(10_000_000)
	}

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := math.MinInt
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) <= CHUNKS {
		return maximum(data)
	}

	max := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(chunkNum int) {
			defer wg.Done()
			start := chunkNum * chunkSize
			end := start + chunkSize
			if chunkNum == CHUNKS-1 {
				end = len(data)
			}
			chunk := data[start:end]
			max[chunkNum] = maximum(chunk)
		}(i)
	}

	wg.Wait()

	return maximum(max)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	d := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	// ваш код здесь
	startTime := time.Now()
	max := maximum(d)
	elapsed := time.Since(startTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	startTime = time.Now()
	max = maxChunks(d)
	elapsed = time.Since(startTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
