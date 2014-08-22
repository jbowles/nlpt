package nlptoken

import (
	//"fmt"
	"testing"
)

func BenchmarkWhiteSpace(*testing.B) {
	var wsdigestone = NewWhiteSpaceDigest()
	wsdigestone.Tknz(ThoreauThree)
}

func TestWhistTknz(t *testing.T) {
	var wsdigest = NewWhiteSpaceDigest()
	tok3, digest := wsdigest.Tknz(ThoreauThree)

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string length=19, got=", len(tok3))
		t.Fail()
	}

	if len(ThoreauThree) != digest.CharCount {
		t.Log("Expected string and digest character counts to be equal, got string length=", len(ThoreauThree), "CharCount=", digest.CharCount)
		t.Fail()
	}
}
