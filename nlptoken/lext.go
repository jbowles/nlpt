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
	wordCount     int
	punctCount    int
	spaceCount    int
	lineCount     int
	charCount     int
	emptyLine     bool
	tokens        []string
	punct         []string
	lastTokenType lexer.TokenType
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
	bytesNonWord = []byte{' ', '\t', '\f', '\v', '\n', '\r', '.', '?', '!', ':', '\\', '"'}
	bytesPunct   = []byte{'.', '?', '!', ':', '\\', '"'}
	bytesSpace   = []byte{' ', '\t', '\f', '\v'}
)

func LexToken(text string) *lexDigest {
	reader := bytes.NewBuffer([]byte(text))
	ldigest := lexDigest{
		wordCount:     0,
		punctCount:    0,
		spaceCount:    0,
		lineCount:     0,
		charCount:     0,
		emptyLine:     true,
		tokens:        make([]string, 0, 0),
		punct:         make([]string, 0, 0),
		lastTokenType: T_NIL,
	}

	lex := lexer.NewSize(lexFunc, reader, 100, 1)

	// Processing the lexer-emitted tokens
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {
		ldigest.charCount += len(t.Bytes())
		switch t.Type() {
		case T_WORD:
			if ldigest.lastTokenType != T_WORD {
				ldigest.wordCount++
				ldigest.tokens = append(ldigest.tokens, string(t.Bytes()))
			}
			ldigest.emptyLine = false
		case T_PUNCT:
			ldigest.punctCount++
			ldigest.punct = append(ldigest.punct, string(t.Bytes()))
			ldigest.emptyLine = false
		case T_NEWLINE:
			ldigest.lineCount++
			ldigest.spaceCount++
			ldigest.emptyLine = true
		case T_SPACE:
			ldigest.spaceCount += len(t.Bytes())
			ldigest.emptyLine = false
		default:
			panic("unreachable")
		}
		ldigest.lastTokenType = t.Type()
	}
	// If last line not empty, up line count
	if !ldigest.emptyLine {
		ldigest.lineCount++
	}
	return &ldigest
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
