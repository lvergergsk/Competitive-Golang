package isomorphic_string_0205

import "testing"

func TestIsIsomorphic01(t *testing.T) {
	rep := isIsomorphic("egg", "add")
	if rep != true {
		t.Error("expect true get ", rep)
	}
}

func TestIsIsomorphic02(t *testing.T) {
	rep := isIsomorphic("foo", "bar")
	if rep != false {
		t.Error("expect false get ", rep)
	}
}

func TestIsIsomorphic03(t *testing.T) {
	rep := isIsomorphic("paper", "title")
	if rep != true {
		t.Error("expect true get ", rep)
	}
}

func TestIsIsomorphic04(t *testing.T) {
	rep := isIsomorphic("ab", "aa")
	if rep != false {
		t.Error("expect false get ", rep)
	}
}
