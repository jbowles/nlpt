/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* A utoken language tokenizer uses unicode code point ranges to segment text.
* It filters through strings leveraging standard library functions to parse and
* collect unicode segments into 'buckets' of letter, number, punctuation, spaces, and symbols.
* Since this package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*
 */

package nlptoken

import (
	"strings"
	"unicode"
)

type uToken struct {
	Letter []string
	Number []string
	Punct  []string
	Space  []string
	Symbol []string
}

func (u *uToken) Initialize() *uToken {
	u = &uToken{
		Letter: []string{},
		Number: []string{},
		Punct:  []string{},
		Space:  []string{},
		Symbol: []string{},
	}
	return u
}

//UTokenize uses unicode package to match runes for tokenization. It can be very useful for really noisy data sets. See documention for for UnicBucket for more details. It returns a slice of tokenized words and a the bucket.
func UTokenize(s string) (tokens []string, u uToken) {
	u.Initialize()
	for _, v := range s {
		switch true {
		case unicode.IsLetter(v):
			u.Letter = append(u.Letter, string(v))
		case unicode.IsSpace(v):
			u.Letter = append(u.Letter, ", ")
		case unicode.IsNumber(v):
			u.Number = append(u.Number, string(v))
			//u.Letter = append(u.Letter, string(v))
		case unicode.IsPunct(v):
			u.Punct = append(u.Punct, string(v))
		case unicode.IsSymbol(v):
			u.Symbol = append(u.Symbol, string(v))
		}
	}
	tokens = strings.Split(strings.Join(u.Letter, ""), ", ")
	return tokens, u
}
