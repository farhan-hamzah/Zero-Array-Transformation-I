package main

import (
	"fmt"
)

// Fungsi untuk cek apakah semua elemen pada rentang setiap query adalah nol
func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	prefixNonZero := make([]int, n+1) // prefixNonZero[i] = jumlah elemen bukan nol sampai indeks i-1

	for i := 0; i < n; i++ {
		prefixNonZero[i+1] = prefixNonZero[i]
		if nums[i] != 0 {
			prefixNonZero[i+1]++
		}
	}

	for _, q := range queries {
		l, r := q[0], q[1]
		if l < 0 || r >= n || l > r {
			// Jika query invalid, langsung return false
			return false
		}
		if prefixNonZero[r+1]-prefixNonZero[l] > 0 {
			// Ada elemen bukan nol pada rentang query
			return false
		}
	}
	return true
}

func main() {
	nums := []int{0, 0, 0, 1, 0, 0}
	queries := [][]int{
		{0, 2}, // semua nol
		{4, 5}, // semua nol
		{0, 3}, // ada 1 di indeks 3, bukan nol
	}

	result := isZeroArray(nums, queries)
	fmt.Println("Apakah semua rentang query bernilai nol?", result)
}
