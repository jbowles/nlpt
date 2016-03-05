package tkz

import (
	"testing"
)

func BenchmarkWhiteSpaceTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzWhiteSpace(ThoreauThree, NewWhiteSpaceDigest())
	}
}

func BenchmarkWhiteSpaceTknzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzWhiteSpace(BadStr, NewWhiteSpaceDigest())
	}
}

func TestWhiteSpaceTknz(t *testing.T) {
	tok3, digest := TknzWhiteSpace(ThoreauThree, NewWhiteSpaceDigest())

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string length=19, got=", len(tok3))
		t.Fail()
	}

	if len(ThoreauThree) != digest.CharCount {
		t.Log("Expected string and digest character counts to be equal, got string length=", len(ThoreauThree), "CharCount=", digest.CharCount)
		t.Fail()
	}
}
