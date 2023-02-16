package main

// cache hold all precomputed and used prime numbers
var cache = make(map[int]bool)

// cachedIsPrime check for number in cache, if number not exist call isPrime
func cachedIsPrime(number int) (response bool) {
	if result, ok := cache[number]; ok {
		response = result
	} else {
		response = isPrime(number)
		cache[number] = response
	}

	return response
}

// isPrime check if number is prime
func isPrime(number int) bool {
	switch {
	case number <= 1:
		return false
	case number <= 3:
		return true
	case number%2 == 0 || number%3 == 0:
		return false
	}

	// Check if number is divisible by any number of the form 6k ± 1 up to sqrt(number)
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}
