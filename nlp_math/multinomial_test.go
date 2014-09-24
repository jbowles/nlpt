package nlp_math

import "testing"

func TestMultiCountCreation(t *testing.T) {
	mc := MultiCount{}
	mc.Count()

	if mc.counts[0] != 0 {
		t.Log("Expected key[0] to have value 0")
		t.Fail()
	}
}

func TestMultiCountIncrement(t *testing.T) {
	mc := MultiCount{}
	mc.Count()

	for i := 0; i < 200; i++ {
		mc.Increment(i)
	}

	z := mc.counts[1]
	if z != 1 {
		t.Log("Expected key[0] to be 1 but got: ", z)
		t.Fail()
	}

	tf := mc.counts[25]
	if tf != 1 {
		t.Log("Expected key[25] to be 201, but got: ", tf)
		t.Fail()
	}

	numKeys := len(mc.counts)
	if numKeys != 200 {
		t.Log("Expected Number of Kyes to be 201, but got: ", numKeys)
		t.Fail()
	}
}
