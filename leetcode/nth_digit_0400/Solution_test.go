package nth_digit_0400

import "testing"

func TestCountPrime01(t *testing.T) {
	rep := findNthDigit(3)
	if rep != 3 {
		t.Error("expect 3 get ", rep)
	}
}

func TestCountPrime02(t *testing.T) {
	rep := findNthDigit(11)
	if rep != 0 {
		t.Error("expect 0 get", rep)
	}
}

func TestCountPrime03(t *testing.T) {
	rep := findNthDigit(10)
	if rep != 1 {
		t.Error("expect 1 get", rep)
	}
}

func TestCountPrime04(t *testing.T) {
	rep := findNthDigit(14)
	if rep != 1 {
		t.Error("expect 1 get", rep)
	}
}
func TestCountPrime05(t *testing.T) {
	rep := findNthDigit(15)
	if rep != 2 {
		t.Error("expect 2 get", rep)
	}
}
