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
	"unicode"
)

type Bucket struct {
	Letter []string
	Number []string
	Punct  []string
	Space  []string
	Symbol []string
}

func (b *Bucket) Initialize() *Bucket {
	b = &Bucket{
		Letter: []string{},
		Number: []string{},
		Punct:  []string{},
		Space:  []string{},
		Symbol: []string{},
	}
	return b
}

// Unic is a unicode tokenizer.
func UnicBucket(s string) (b Bucket) {
	b.Initialize()
	for _, v := range s {
		switch true {
		case unicode.IsLetter(v):
			b.Letter = append(b.Letter, string(v))
		case unicode.IsSpace(v):
			b.Letter = append(b.Letter, ", ")
		case unicode.IsNumber(v):
			b.Number = append(b.Number, string(v))
		case unicode.IsPunct(v):
			b.Punct = append(b.Punct, string(v))
		case unicode.IsSymbol(v):
			b.Symbol = append(b.Symbol, string(v))
		}
	}
	return
}
