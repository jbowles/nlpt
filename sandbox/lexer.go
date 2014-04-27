package main

import "os"
import "fmt"
import "github.com/jbowles/go_lexer"

// Usage : wordcount <filename>
func usage() {
	fmt.Printf("usage: %s <filename>\n", os.Args[0])
}

// We define our lexer tokens starting from the pre-defined EOF token
const (
	T_EOF lexer.TokenType = lexer.TokenTypeEOF
	T_NIL                 = lexer.TokenTypeEOF + iota
	T_SPACE
	T_NEWLINE
	T_WORD
	T_PUNCT
)

// List gleaned from isspace(3) manpage
var bytesNonWord = []byte{' ', '\t', '\f', '\v', '\n', '\r', '.', '?', '!', ':', '\\', '"'}

var bytesPunct = []byte{'.', '?', '!', ':', '\\', '"'}

var bytesSpace = []byte{' ', '\t', '\f', '\v'}

const charNewLine = '\n'

const charReturn = '\r'

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	var file *os.File
	var error error

	file, error = os.Open(os.Args[1])

	if error != nil {
		panic(error)
	}

	var chars int = 0

	var words int = 0
	tok := []string{}

	var spaces int = 0

	var lines int = 0

	var punctuation int = 0
	punkt := []string{}

	// To help us track last line
	var emptyLine bool = true

	// Create our lexer
	// NewSize(startState, reader, readerBufLen, channelCap)
	lex := lexer.NewSize(lexFunc, file, 100, 1)

	var lastTokenType lexer.TokenType = T_NIL

	// Process lexer-emitted tokens
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {

		chars += len(t.Bytes())

		switch t.Type() {
		case T_WORD:
			if lastTokenType != T_WORD {
				words++
				tok = append(tok, string(t.Bytes()))
			}
			emptyLine = false

		case T_PUNCT:
			punctuation++
			punkt = append(punkt, string(t.Bytes()))
			emptyLine = false

		case T_NEWLINE:
			lines++
			spaces++
			emptyLine = true

		case T_SPACE:
			spaces += len(t.Bytes())
			emptyLine = false

		default:
			panic("unreachable")
		}

		lastTokenType = t.Type()
	}

	// If last line not empty, up line count
	if !emptyLine {
		lines++
	}

	fmt.Printf("tokens %v\n", tok)
	fmt.Printf("punctuation %v\n", punkt)
	fmt.Printf("tokenLength %v\n", len(tok))
	fmt.Printf("%d words, %d punctuation, %d spaces, %d lines, %d chars\n", words, punctuation, spaces, lines, chars)
}

func lexFunc(l lexer.Lexer) lexer.StateFn {
	// EOF
	if l.MatchEOF() {
		l.EmitEOF()
		return nil // We're done here
	}

	// Non-Space run
	if l.NonMatchOneOrMoreBytes(bytesNonWord) {
		l.EmitTokenWithBytes(T_WORD)

		// Punctuation
	} else if l.MatchOneOrMoreBytes(bytesPunct) {
		l.EmitTokenWithBytes(T_PUNCT)

		// Space run
	} else if l.MatchOneOrMoreBytes(bytesSpace) {
		l.EmitTokenWithBytes(T_SPACE)

		// Line Feed
	} else if charNewLine == l.PeekRune(0) {
		l.NextRune()
		l.EmitTokenWithBytes(T_NEWLINE)
		l.NewLine()

		// Carriage-Return with optional line-feed immediately following
	} else if charReturn == l.PeekRune(0) {
		l.NextRune()
		if charNewLine == l.PeekRune(0) {
			l.NextRune()
		}
		l.EmitTokenWithBytes(T_NEWLINE)
		l.NewLine()
	} else {
		panic("unreachable")
	}

	return lexFunc
}
