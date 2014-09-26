package nlp_stats

import (
	"fmt"
	"testing"
)

// Randomness comes in each definition of the Uniform function.
// Once defined, the results can be predicted and so we can test it.
func TestUniformReturnsRandFloat64(t *testing.T) {
	unf := Uniform()
	val := unf()
	if val != 0.2318300419376769 {
		t.Log("Expected first call Uniform() to return '0.2318300419376769', but got", val)
		t.Fail()
	}
}

func TestUniformReturnsFloat64(t *testing.T) {
	unf := Uniform()
	typ := fmt.Sprintf("%T", unf())

	if typ != "float64" {
		t.Log("Expected Uniform() to return 'float64' type, but got", typ)
		t.Fail()
	}
}
