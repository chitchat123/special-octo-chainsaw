//go:build cache

package main

import "math"

// if project built with cache tag we initialize cache with first 10000 numbers
func init() {
	sieveOfEratosthenes(10000)

}

// sieveOfEratosthenes save to cache precomputed prime numbers up to limit
func sieveOfEratosthenes(limit int) {
	// Create a boolean map indicating whether each number is prime
	nums := make(map[int]bool, limit+1)
	for i := 2; i <= limit; i++ {
		nums[i] = true
	}

	// Iterate over all numbers up to the square root of the limit
	sqrtLimit := int(math.Sqrt(float64(limit)))
	for i := 2; i <= sqrtLimit; i++ {
		if nums[i] {
			// Mark all multiples of i as not prime
			for j := i * i; j <= limit; j += i {
				nums[j] = false
			}
		}
	}
	cache = nums
}
