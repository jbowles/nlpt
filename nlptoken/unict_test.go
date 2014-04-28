package nlptoken

import (
	"testing"
)

func BenchmarkUTokenizeGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UTokenize(ThoreauOne)
	}
}

func BenchmarkUTokenizeBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UTokenize(BadStr)
	}
}

func TestUToken(t *testing.T) {
	_, bucket := UTokenize(BadStr)
	first_symbol := bucket.Symbol[0]
	second_symbol := bucket.Symbol[1]
	test_first_symbol := "<"
	test_second_symbol := "="

	if len(bucket.Letter) != 128 {
		t.Log("Expected letter length to be 128")
		t.Fail()
	}

	if len(bucket.Number) != 3 {
		t.Log("Expected number length to be 3")
		t.Fail()
	}

	if len(bucket.Punct) != 5 {
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

func TestTokenizeBadString(t *testing.T) {
	tok, _ := UTokenize(BadStr)

	if len(tok) != 27 {
		t.Log("Expected thoreauThree string to be 19 words")
		t.Fail()
	}
}

func TestUTokenize(t *testing.T) {
	tok1, _ := UTokenize(ThoreauOne)
	tok2, _ := UTokenize(ThoreauTwo)
	tok3, _ := UTokenize(ThoreauThree)

	if len(tok1) != 44 {
		t.Log("Expected thoreauOne string to be 44 words")
		t.Fail()
	}

	if len(tok2) != 30 {
		t.Log("Expected thoreauTwo string to be 30 words")
		t.Fail()
	}

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string to be 19 words")
		t.Fail()
	}
}
