package nlptoken

import (
	"testing"
)

func BenchmarkWhiteSpace(*testing.B) {
	WhiteSpace(ThoreauThree, "")
}

func TestWhiteSpace(t *testing.T) {
	tok := WhiteSpace(ThoreauThree, "")
	if len(tok) != 19 {
		t.Log("Expected thoreauThree string to be 19 words")
		t.Fail()
	}
}

func TestWhiteSpaceReadSepComma(t *testing.T) {
	tokComma := WhiteSpace(ThoreauOne, ", ") //44
	for _, s := range tokComma {
		bt := []byte(s)
		if bt[len(bt)-2] != 44 {
			t.Log("Expected word to end with comma ', ' but got", s)
			t.Fail()
		}
	}
}

func TestWhiteSpaceReadSepDash(t *testing.T) {
	tokDash := WhiteSpace(ThoreauTwo, " - ") //45
	for _, s := range tokDash {
		bt := []byte(s)
		if bt[len(bt)-2] != 45 {
			t.Log("Expected word to end with dash ' - ' but got", s)
			t.Fail()
		}
	}
}
