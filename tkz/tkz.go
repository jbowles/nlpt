/*
* Top level public API for using tokenizers.
 */
package tkz

import "github.com/jbowles/nlpt_tkz/Godeps/_workspace/src/github.com/jbowles/go_lexer"

type Digest struct {
	Tokens         []string
	DowncaseTokens []string
	TokenBytes     map[string][]byte
	Bytes          []byte
	SpaceCount     int
	CharCount      int
	Letter         []string
	Title          []string
	Number         []string
	Punct          []string
	Space          []string
	Symbol         []string
	TokenCount     int
	PunctCount     int
	LineCount      int
	EmptyLine      bool
	LastTokenType  lexer.TokenType
}

//TokenizeStr is a top-level function to choose tokenizer types: TknzStateFun, TknzUnicode, TknzWhiteSpace.
//It creates an abstraction around all the tokenizer implementations for a simple API, facilitating an easy call from the client and for binary installations.
func TokenizeStr(text, typ string) (tokens []string, digest *Digest) {

	switch typ {
	case "lex":
		tokens, digest = TknzStateFun(text, NewStateFnDigest())
	case "unicode":
		tokens, digest = TknzUnicode(text, NewUnicodeMatchDigest())
	case "whitespace":
		tokens, digest = TknzWhiteSpace(text, NewWhiteSpaceDigest())
	default:
		panic("Tokenizer type not supported")
	}
	return
}

//TokenizeBytes is a top-level function to choose tokenizer types: TknzStateFunBytes, TknzUnicodeBytes.
//It creates an abstraction around tokenizer implementations that accept byte slices for a simple API, facilitating an easy call from the client and for binary installations.
func TokenizeBytes(textBytes []byte, typ string) (digest *Digest) {

	switch typ {
	case "lex":
		digest = TknzStateFunBytes(textBytes, NewStateFnDigestBytes())
	case "unicode":
		digest = TknzUnicodeBytes(textBytes, NewUnicodeMatchDigest())
	default:
		panic("Tokenizer type not supported")
	}
	return
}
