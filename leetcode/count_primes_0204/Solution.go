package count_primes_0204

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	//False: the number is a prime.
	//True: the number is not a prime.
	numbers := make([]bool, n)
	count := 0

	numbers[0] = true
	numbers[1] = true
	for i := 2; i < len(numbers); i++ {
		//println(i, numbers[i])
		if numbers[i] == false {
			count++
			for k := 2; k*i < len(numbers); k++ {
				numbers[k*i] = true
			}
		}
	}

	return count
}
