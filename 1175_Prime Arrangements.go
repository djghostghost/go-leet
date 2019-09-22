package main

import "math"

const MOD = 1000000007

func numPrimeArrangements(n int) int {

	if n == 1 {
		return 1
	}
	primeCount := 0

	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primeCount++
		}
	}

	return (fact(primeCount) * fact(n-primeCount)) % MOD

}

func isPrime(n int) bool {

	if n == 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

func fact(n int) int {

	result := 1
	for i := 1; i <= n; i++ {
		result = (result * i) % MOD
	}
	return result % MOD
}
