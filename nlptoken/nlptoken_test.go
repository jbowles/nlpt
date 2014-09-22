package nlptoken

import (
	"fmt"
	"testing"
)

func TestTokenizeLextOption(t *testing.T) {
	tokens, digest := Tokenize("lext", ThoreauThree)
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlptoken.LexerDigest" {
		t.Log("Expected digest to be *nlptoken.LexerDigest", typ)
		t.Fail()
	}
}

func TestTokenizeDefaultOption(t *testing.T) {
	tokens, digest := Tokenize("lexer", ThoreauThree)
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlptoken.WhiteSpaceDigest" {
		t.Log("Expected digest to be *nlptoken.WhiteSpaceDigest", typ)
		t.Fail()
	}
}
