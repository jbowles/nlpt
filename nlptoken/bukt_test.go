package nlptoken

import (
	"testing"
)

func BenchmarkBuktTknzGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bdigestone = NewBucketDigest()
		bdigestone.Tknz(ThoreauOne)
	}
}

func BenchmarkBuktTnkzBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bdigesttwo = NewBucketDigest()
		bdigesttwo.Tknz(BadStr)
	}
}

func TestBuktBadStr(t *testing.T) {
	var bdigest = NewBucketDigest()
	_, digest := bdigest.Tknz(BadStr)
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

func TestBuktTknzBadString(t *testing.T) {
	var bdigest = NewBucketDigest()
	tok, _ := bdigest.Tknz(BadStr)

	if len(tok) != 27 {
		t.Log("Expected BadStr string to be length=25, got=", len(tok))
		t.Fail()
	}
}

func TestBucketTknzOne(t *testing.T) {
	var bdigest = NewBucketDigest()
	tok1, _ := bdigest.Tknz(ThoreauOne)

	if len(tok1) != 44 {
		t.Log("Expected thoreauOne string to be length=44, got=", len(tok1))
		t.Fail()
	}
}

func TestBucketTknzTwo(t *testing.T) {
	var bdigest = NewBucketDigest()
	tok2, _ := bdigest.Tknz(ThoreauTwo)

	if len(tok2) != 30 {
		t.Log("Expected thoreauTwo string to be length=30", len(tok2))
		t.Fail()
	}
}

func TestBucketTknzThree(t *testing.T) {
	var bdigest = NewBucketDigest()
	tok3, _ := bdigest.Tknz(ThoreauThree)

	if len(tok3) != 19 {
		t.Log("Expected thoreauThree string to be lenght=19, got=", len(tok3))
		t.Fail()
	}
}
