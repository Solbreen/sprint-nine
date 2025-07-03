package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	require.Equal(t, len(generateRandomElements(0)), 0)
	require.Equal(t, len(generateRandomElements(-2)), 0)
	require.Equal(t, len(generateRandomElements(1498)), 1498)
}

func TestMaximum(t *testing.T) {

	require.Equal(t, maximum([]int{}), 0)
	require.Equal(t, maximum([]int{10}), 10)
	require.Equal(t, maximum([]int{10, 2, 3, 5, -4, 123, 3}), 123)
	require.Equal(t, maximum([]int{-2, -6, -4}), -2)
	require.Equal(t, maximum([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}), 3)
	require.Equal(t, maximum([]int{math.MinInt}), math.MinInt)
	require.Equal(t, maximum([]int{10, 2, 3, 5, -4, 123, math.MaxInt}), math.MaxInt)
}
