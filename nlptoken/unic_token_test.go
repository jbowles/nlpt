package nlptoken

import (
	"testing"
)

func TestUnicBucket(t *testing.T) {
	bucket := UnicBucket(BadStr)
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
