package nth_digit_0400

import "math"

func findNthDigit(n int) int {
	res := 0
	digitCount := 1
	digit := 1
	for n-9*digit*digitCount > 0 {
		n -= 9 * digit * digitCount
		res += 9 * digit
		digit *= 10
		digitCount++
	}
	//println(res, digit, digitCount, n)

	{
		step := int(math.Ceil(float64(n) / float64(digitCount)))
		res += step
		n -= (step-1) * digitCount+1
	}

	//println(res, digit, digitCount, n)
	//e.g. 12345, 2 -> (/10^(5-2)=/100) 12 -> (%10) 2
	if digitCount-(n+1)>0 {
		//println(int(math.Pow(float64(10), float64(digitCount-(n+1)))))
		res /= int(math.Pow(float64(10), float64(digitCount-(n+1))))
	}

	res %= 10
	return res
}
