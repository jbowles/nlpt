package nlpt_math

import "testing"

func TestMinInt(t *testing.T) {
	max := int(483264732)
	min := int(4387430)
	val := MinInt32(max, min)
	if val != min {
		t.Log("Expected MinInt to return: ", min, "but got: ", val)
		t.Fail()
	}
}

func TestMaxInt32(t *testing.T) {
	max := int(483264732)
	min := int(-456)
	val := MaxInt32(max, min)
	if val != max {
		t.Log("Expected MinInt to return: ", max, "but got: ", val)
		t.Fail()
	}
}
