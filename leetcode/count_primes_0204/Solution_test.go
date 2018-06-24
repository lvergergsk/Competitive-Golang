package count_primes_0204

import "testing"

func TestCountPrime01(t *testing.T) {
	// non-prime: 0,1,4,6,8,9
	// prime: 2,3,5,7
	rep := countPrimes(10)
	if rep != 4 {
		t.Error("expect 4 get ", rep)
	}
}

func TestCountPrime02(t *testing.T) {
	// non-prime: 0,1
	// prime: -
	rep := countPrimes(2)
	if rep != 0 {
		t.Error("expect 0 get", rep)
	}
}
