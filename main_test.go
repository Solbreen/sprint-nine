package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле

var testsTGRE = []struct {
	in  int
	out int
}{
	{len(generateRandomElements(0)), 0},
	{len(generateRandomElements(-2)), 0},
	{len(generateRandomElements(1498)), 1498},
}

var testsTM = []struct {
	in  int
	out int
}{
	{maximum([]int{}), 0},
	{maximum([]int{10}), 10},
	{maximum([]int{10, 2, 3, 5, 4, 123, 3}), 123},
	{maximum([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3}), 3},
	{maximum([]int{10, 2, 3, 5, 4, 123, math.MaxInt}), math.MaxInt},
}

func TestGenerateRandomElements(t *testing.T) {
	for _, tt := range testsTGRE {
		require.Equal(t, tt.in, tt.out)
	}
}

func TestMaximum(t *testing.T) {
	for _, tt := range testsTM {
		require.Equal(t, tt.in, tt.out)
	}
}
