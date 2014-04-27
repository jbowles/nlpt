package nlptoken

import (
	//"fmt"
	"testing"
)

const (
	BadStr = "expect0.7rant7! Then I want to show Snow White and the Seven Dwarves. <=AndThe start of a new sentence. And\n then\n\nagain for One and NASA?"
	//44 words
	ThoreauOne = "I went to the woods because I wished to live deliberately, to front only the essential facts of life, and see if I could not learn what it had to teach, and not, when I came to die, discover that I had not lived."
	//30 words
	ThoreauTwo = "If one advances confidently in the direction of his dreams, and endeavors to live the life which he has imagined, he will meet with a success unexpected in common hours."
	//19 words
	ThoreauThree = "What lies behind us and what lies ahead of us are tiny matters compared to what lives within us."
)

func BenchmarkWhiteSpace(*testing.B) {
	WhiteSpace(ThoreauThree, "")
}

func BenchmarkUnicTokenGoodStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnicToken(ThoreauOne)
	}
}

func BenchmarkUnicTokenBucketBadStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnicToken(BadStr)
	}
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

func TestTokenizeBadString(t *testing.T) {
	tok, _ := UnicToken(BadStr)

	if len(tok) != 27 {
		t.Log("Expected thoreauThree string to be 19 words")
		t.Fail()
	}
}

func TestUnicToken(t *testing.T) {
	tok1, _ := UnicToken(ThoreauOne)
	tok2, _ := UnicToken(ThoreauTwo)
	tok3, _ := UnicToken(ThoreauThree)

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
