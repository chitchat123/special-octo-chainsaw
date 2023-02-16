package main

import (
	"math/rand"
	"testing"
)

func Test_cachedIsPrime(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"test-1", 1, false},
		{"test-2", 2, true},
		{"test-3", 3, true},
		{"test-4", 4, false},
		{"test-5", 5, true},
		{"test-6", 6, false},
		{"test-7", 7, true},
		{"test-8", 8, false},
		{"test-9", 9, false},
		{"test-10", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cachedIsPrime(tt.args); got != tt.want {
				t.Errorf("cachedIsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"test-1", 1, false},
		{"test-2", 2, true},
		{"test-3", 3, true},
		{"test-4", 4, false},
		{"test-5", 5, true},
		{"test-6", 6, false},
		{"test-7", 7, true},
		{"test-8", 8, false},
		{"test-9", 9, false},
		{"test-10", 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(rand.Intn(100)) //nolint:gosec
	}
}

func Benchmark_cachedIsPrime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cachedIsPrime(rand.Intn(100)) //nolint:gosec
	}
}
