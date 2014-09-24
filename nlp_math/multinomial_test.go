package nlp_math

import "testing"

func TestMaximumLikelihood(t *testing.T) {
	mn := Multinomial{}
	mc := MultiCount{}
	mc.Count()

	for i := 0; i < 200; i++ {
		mc.Increment(i)
	}
	mn.MaximumLikelihood(mc)

	sixProb := mn.Probability(6)
	if sixProb != 0.005 {
		t.Log("Expected probability to be 0.005 but got: ", sixProb)
		t.Fail()
	}

	for i := 0; i < 200; i++ {
		mn.distribution[12]++
	}
	twelveProb := mn.Probability(12)
	if twelveProb != 200.005 {
		t.Log("Expected probability to be 0.005 but got: ", twelveProb)
		t.Fail()
	}

	for i := 0; i < 800; i++ {
		mn.distribution[99]++
	}
	nineNineProb := mn.Probability(99)
	if nineNineProb != 800.005 {
		t.Log("Expected probability to be 0.005 but got: ", nineNineProb)
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

func TestMultiNomial(t *testing.T) {
	mn := Multinomial{
		distribution: make(map[int]float64),
	}

	for i := 0; i < 25; i++ {
		mn.distribution[i]++
	}

	threeProb := mn.Probability(3)
	if threeProb != 1 {
		t.Log("Expected probability to be 1.0 but got: ", threeProb)
		t.Fail()
	}

	for i := 0; i < 200; i++ {
		mn.distribution[6]++
	}
	sixProb := mn.Probability(6)
	if sixProb != 201 {
		t.Log("Expected probability to be 0.0 but got: ", sixProb)
		t.Fail()
	}
}
