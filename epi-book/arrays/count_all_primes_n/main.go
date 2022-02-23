package epi

// countPrimes returns the number of prime numbers less than or equal to n.
// eliminate (1), (2), and number itself, multiples of these numbers, all even numbers.
func countPrimes(n int) int {
	// eliminate an odd prime 1 and only even prime 2.
	if n < 3 {
		return 0
	}

	nonPrimes := make([]bool, n) // all are false

	count := n / 2 // we will decrement it when we found odd non-prime

	// sieve of Eratosthenes
	for i := 3; i*i < n; i += 2 {
		if nonPrimes[i] {
			continue
		}

		// mark all multiples of i as non-prime
		for j := i * i; j < n; j += 2 * i {
			if !nonPrimes[j] {
				//nonPrimes[j] = false
				count--
			}
			nonPrimes[j] = true
		}
	}
	return count
}
