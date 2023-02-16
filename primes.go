package main

var cache = make(map[int]bool)

func cachedIsPrime(n int) (response bool) {
	if result, ok := cache[n]; ok {
		response = result
	} else {
		response = isPrime(n)
		cache[n] = response
	}

	return response
}

func isPrime(n int) bool {
	switch {
	case n <= 1:
		return false
	case n <= 3:
		return true
	case n%2 == 0 || n%3 == 0:
		return false
	}

	// Check if n is divisible by any number of the form 6k ± 1 up to sqrt(n)
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
