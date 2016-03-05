package tkz

import (
	"fmt"
	"testing"
)

func BenchmarkUncdMatchTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzUnicode(ThoreauOne, NewUnicodeMatchDigest())
	}
}

func BenchmarkUncdMatchTnkzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzUnicode(BadStr, NewUnicodeMatchDigest())
	}
}

func BenchmarkUncdMatchTknzBytesGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzUnicodeBytes(thoneByte, NewUnicodeMatchDigest())
	}
}

func BenchmarkUncdMatchTnkzBytesBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TknzUnicodeBytes(badstrByte, NewUnicodeMatchDigest())
	}
}

func TestTknzUnicodeBytes(t *testing.T) {
	b := []byte(ThoreauThree)
	digest := TknzUnicodeBytes(b, NewUnicodeMatchDigest())
	//fmt.Printf("UNI bytes %v\n", digest.Bytes)
	//fmt.Printf("UNI bytes stringified %v\n", string(digest.Bytes))

	/*
		if len(digest.Bytes) != 95 {
			t.Log("Expected digest.Bytes to be length=95, got=", len(digest.TokenBytes))
			t.Fail()
		}
	*/

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpt_tkz.Digest" {
		t.Log("Expected digest to be *nlpt_tkz.StateFnDigest", typ)
		t.Fail()
	}
}

func TestUncdMatchBadStr(t *testing.T) {
	_, digest := TknzUnicode(BadStr, NewUnicodeMatchDigest())
	first_symbol := digest.Symbol[0]
	second_symbol := digest.Symbol[1]
	test_first_symbol := "<"
	test_second_symbol := "="

	if len(digest.Letter) != 128 {
		t.Log("Expected letter length to be 128")
		t.Fail()
	}

	if len(digest.Number) != 3 {
		t.Log("Expected number length to be 3")
		t.Fail()
	}

	if len(digest.Punct) != 5 {
		t.Log("Expected punctuation count to be 5")
		t.Fail()
	}

	if first_symbol != test_first_symbol {
		t.Log("Expected to see", test_first_symbol, "instead got", first_symbol)
		t.Fail()
	}

	if second_symbol != test_second_symbol {
		t.Log("Expected to see", test_first_symbol, "instead got", first_symbol)
		t.Fail()
	}
}

func TestBuktUncdMatchBadString(t *testing.T) {
	tok, _ := TknzUnicode(BadStr, NewUnicodeMatchDigest())

	if len(tok) != 27 {
		t.Log("Expected BadStr string to be length=25, got=", len(tok))
		t.Fail()
	}
}

func TestUncdMatchTknzOne(t *testing.T) {
	tok1, _ := TknzUnicode(ThoreauOne, NewUnicodeMatchDigest())

	if len(tok1) != 54 {
		t.Log("Expected thoreauOne string to be length=54, got=", len(tok1))
		t.Fail()
	}
}

func TestUncdMatchTknzTwo(t *testing.T) {
	tok2, _ := TknzUnicode(ThoreauTwo, NewUnicodeMatchDigest())

	if len(tok2) != 30 {
		t.Log("Expected thoreauTwo string to be length=30", len(tok2))
		t.Fail()
	}
}

func TestUncdMatchTknzThree(t *testing.T) {
	tok3, _ := TknzUnicode(ThoreauThree, NewUnicodeMatchDigest())

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string to be lenght=19, got=", len(tok3))
		t.Fail()
	}
}
