/*
* Copyright ©2015 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Following Lexical Scanning in Go, uses a 'State Function' to acheive tokenization.
* For now, adapting go_lexer and the word_count example from a fork of https://github.com/iNamik/go_lexer.
 */

//TODO change lexer to use runes instead of bytes, this will give greater range and make it easier to assign code point blocks from go unicode package. For example, punctuation is pretty small but only up to rune 191 is there equality with byte casting.
package tkz

import (
	"bytes"
	"github.com/jbowles/nlpt_tkz/Godeps/_workspace/src/github.com/jbowles/go_lexer"
	"strings"
)

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
	bytesNonWord = []byte{' ', '\t', '\f', '\v', '\n', '\r', '.', '?', '!', ':', '\\', '"', ','} //, '+', '>', '|', '_', '*'} // '@'} //, '+', '>', '|', '_', '*'}
	bytesPunct   = []byte{'.', '?', '!', ':', '\\', '"', ',', '(', ')'}
	bytesSpace   = []byte{' ', '\t', '\f', '\v'}
	runesNonWord = []rune{' ', '\t', '\f', '\v', '\n', '\r', '.', '?', '!', ':', '\\', '"', ','}
	runesPunct   = []rune{'.', '?', '!', ':', '\\', '"', ',', '(', ')'}
	runesSpace   = []rune{' ', '\t', '\f', '\v'}
)

func NewStateFnDigest() *Digest {
	return &Digest{
		TokenCount:     0,
		PunctCount:     0,
		SpaceCount:     0,
		LineCount:      0,
		CharCount:      0,
		EmptyLine:      true,
		Tokens:         make([]string, 0, 0),
		DowncaseTokens: make([]string, 0, 0),
		Bytes:          make([]byte, 0, 0),
		TokenBytes:     make(map[string][]byte),
		Punct:          make([]string, 0, 0),
		LastTokenType:  T_NIL,
	}
}

func NewStateFnDigestBytes() *Digest {
	return &Digest{
		Bytes:         make([]byte, 0, 0),
		LastTokenType: T_NIL,
	}
}

func TknzStateFunBytes(byteSeq []byte, digest *Digest) *Digest {
	reader := bytes.NewBuffer(byteSeq)
	lex := lexer.NewSize(lexFunc, reader, 100, 1)

	bufferCache := new(bytes.Buffer)
	//bytePadding := []byte{32}
	// Processing the lexer-emitted tokens
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {
		switch t.Type() {
		case T_WORD:
			if digest.LastTokenType != T_WORD {
				bufferCache.Write(t.Bytes())
			}
		case T_PUNCT:
			//bufferCache.Write(bytePadding)
			bufferCache.Write(t.Bytes())
			//bufferCache.Write(bytePadding)
		case T_NEWLINE:
			bufferCache.Write(t.Bytes())
		case T_SPACE:
			bufferCache.Write(t.Bytes())
		default:
			panic("unreachable")
		}
		digest.LastTokenType = t.Type()
	}
	digest.Bytes = bufferCache.Bytes()
	return digest
}

func TknzStateFun(text string, digest *Digest) ([]string, *Digest) {
	reader := bytes.NewBuffer([]byte(text))
	lex := lexer.NewSize(lexFunc, reader, 100, 1)

	// Processing the lexer-emitted tokens
	for t := lex.NextToken(); lexer.TokenTypeEOF != t.Type(); t = lex.NextToken() {
		// save some cycles.. allocate variables from Bytes function
		lexBytes := t.Bytes()
		digest.CharCount += len(lexBytes)
		stringedBytes := string(lexBytes)
		switch t.Type() {
		case T_WORD:
			if digest.LastTokenType != T_WORD {
				digest.TokenCount++
				digest.Tokens = append(digest.Tokens, stringedBytes)
				digest.TokenBytes[stringedBytes] = lexBytes
				digest.DowncaseTokens = append(digest.DowncaseTokens, strings.ToLower(stringedBytes))
				digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytes)
			}
			digest.EmptyLine = false
		case T_PUNCT:
			digest.PunctCount++
			digest.Punct = append(digest.Punct, stringedBytes)
			digest.TokenBytes[string(lexBytes)] = lexBytes
			digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytes)
			digest.EmptyLine = false
		case T_NEWLINE:
			digest.LineCount++
			digest.SpaceCount++
			digest.EmptyLine = true
			digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytes)
		case T_SPACE:
			digest.SpaceCount += len(t.Bytes())
			digest.EmptyLine = false
			digest.Bytes = ConcatByteSlice(digest.Bytes, lexBytes)
		default:
			panic("unreachable")
		}
		digest.LastTokenType = t.Type()
	}
	// If last line not empty, up line count
	if !digest.EmptyLine {
		digest.LineCount++
	}
	return digest.Tokens, digest
}

//TODO change lexer to use runes instead of bytes, this will give greater range and make it easier to assign code point blocks from go unicode package. For example, punctuation is pretty small but only up to rune 191 is there equality with byte casting.
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
