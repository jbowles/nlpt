package nlpttoken

import (
	"fmt"
	"testing"
)

func TestTokenizeLexOption(t *testing.T) {
	tokens, digest := Tokenize("lex", ThoreauThree)
	//fmt.Printf("Tokens = %v\n DigestType = %T\n", tokens, digest)
	//fmt.Printf("DIGEST %v", digest)

	if len(tokens) != 19 {
		t.Log("Expected thoreauThree to be length=19, got=", len(tokens))
		t.Fail()
	}

	typ := fmt.Sprintf("%T", digest)
	if typ != "*nlpttoken.LexerDigest" {
		t.Log("Expected digest to be *nlpttoken.LexerDigest", typ)
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
	if typ != "*nlpttoken.WhiteSpaceDigest" {
		t.Log("Expected digest to be *nlpttoken.WhiteSpaceDigest", typ)
		t.Fail()
	}
}
