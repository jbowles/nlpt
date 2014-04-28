/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Following Lexical Scanning in Go, uses a 'State Function' to acheive tokenization.
* For now, adapting go_lexer and the word_count example.
 */

package nlptoken

import (
	"bytes"
	"github.com/jbowles/go_lexer"
)

type lexDigest struct {
	WordCount     int
	PunctCount    int
	SpaceCount    int
	LineCount     int
	CharCount     int
	EmptyLine     bool
	Tokens        []string
	Punct         []string
	LastTokenType lexer.TokenType
}

//Lexer tokens starting from the pre-defined EOF token
const (
	T_EOF lexer.TokenType = lexer.TokenTypeEOF
	T_NIL                 = lexer.TokenTypeEOF + iota
	T_SPACE
	T_NEWLINE
	T_WORD
	T_PUNCT
	charNewLine = '\n'
	charReturn  = '\r'
)

// List gleaned from isspace(3) manpage
var (
	bytesNonWord = []byte{' ', '\t', '\f', '\v', '\n', '\r', '.', '?', '!', ':', '\\', '"', ','}
	bytesPunct   = []byte{'.', '?', '!', ':', '\\', '"', ','}
	bytesSpace   = []byte{' ', '\t', '\f', '\v'}
)

func LexTokenize(text string) *lexDigest {
	reader := bytes.NewBuffer([]byte(text))
	ldigest := lexDigest{
		WordCount:     0,
		PunctCount:    0,
		SpaceCount:    0,
		LineCount:     0,
		CharCount:     0,
		EmptyLine:     true,
		Tokens:        make([]string, 0, 0),
		Punct:         make([]string, 0, 0),
		LastTokenType: T_NIL,
	}

	lex := lexer.NewSize(lexFunc, reader, 100, 1)

	// Processing the lexer-emitted tokens
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {
		ldigest.CharCount += len(t.Bytes())
		switch t.Type() {
		case T_WORD:
			if ldigest.LastTokenType != T_WORD {
				ldigest.WordCount++
				ldigest.Tokens = append(ldigest.Tokens, string(t.Bytes()))
			}
			ldigest.EmptyLine = false
		case T_PUNCT:
			ldigest.PunctCount++
			ldigest.Punct = append(ldigest.Punct, string(t.Bytes()))
			ldigest.EmptyLine = false
		case T_NEWLINE:
			ldigest.LineCount++
			ldigest.SpaceCount++
			ldigest.EmptyLine = true
		case T_SPACE:
			ldigest.SpaceCount += len(t.Bytes())
			ldigest.EmptyLine = false
		default:
			panic("unreachable")
		}
		ldigest.LastTokenType = t.Type()
	}
	// If last line not empty, up line count
	if !ldigest.EmptyLine {
		ldigest.LineCount++
	}
	return &ldigest
}

//lexFunc is a State-Function that matches ranges of bytes, emits those bytes, and returns its own StatFn.
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
