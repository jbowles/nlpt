package tkz

import (
	"testing"
)

func BenchmarkStateFnTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzStateFun(ThoreauOne, NewStateFnDigest())
	}
}

func BenchmarkStateFnTknzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzStateFun(BadStr, NewStateFnDigest())
	}
}
func BenchmarkStateFnTknzBytesGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzStateFunBytes(thoneByte, NewStateFnDigestBytes())
	}
}

func BenchmarkStateFnTknzBytesBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzStateFunBytes(badstrByte, NewStateFnDigestBytes())
	}
}

func TestStateFnGoodStr(t *testing.T) {
	tokens, lxd := TknzStateFun(ThoreauOne, NewStateFnDigest())

	if lxd.TokenCount != 53 {
		t.Log("Expected word count to be 53, but got", lxd.TokenCount)
		t.Fail()
	}

	if len(tokens) != lxd.TokenCount {
		t.Log("Expected tokens == lxd.TokenCount, but got", len(tokens))
		t.Fail()
	}

	if lxd.PunctCount != 6 {
		t.Log("Expected punct count to be 6, but got", lxd.PunctCount)
		t.Fail()
	}

	if lxd.SpaceCount != 53 {
		t.Log("Expected space count to be 53, but got", lxd.SpaceCount)
		t.Fail()
	}

	if lxd.LineCount != 1 {
		t.Log("Expected line count to be 1, but got", lxd.LineCount)
		t.Fail()
	}

	if lxd.CharCount != 231 {
		t.Log("Expected char count to be 231, but got", lxd.CharCount)
		t.Fail()
	}

	if lxd.LastTokenType != 4 {
		t.Log("Expected last token type to be 4, but got", lxd.LastTokenType)
		t.Fail()
	}

	if len(lxd.Tokens) != lxd.TokenCount {
		t.Log("Expected token and word count to be equal. Tokens=", len(lxd.Tokens), "Words=", lxd.TokenCount)
		t.Fail()
	}

}
