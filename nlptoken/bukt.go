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

type BucketDigest struct {
	Letter []string
	Title  []string
	Number []string
	Punct  []string
	Space  []string
	Symbol []string
}

//NewBucketDigest initializes a new BucketDigest and explicitly allocates a length and cap on all stirng slices.
func NewBucketDigest() *BucketDigest {
	return &BucketDigest{
		Letter: make([]string, 0, 0),
		Title:  make([]string, 0, 0),
		Number: make([]string, 0, 0),
		Punct:  make([]string, 0, 0),
		Space:  make([]string, 0, 0),
		Symbol: make([]string, 0, 0),
	}
}

//UTokenize uses unicode package to match runes for tokenization. It can be very useful for really noisy data sets. See documention for for UnicBucket for more details. It returns a slice of tokenized words and a the bucket.
func (bdigest *BucketDigest) Tknz(text string) ([]string, *BucketDigest) {
	for _, v := range text {
		switch true {
		case unicode.IsTitle(v):
			bdigest.Title = append(bdigest.Title, string(v))
		case unicode.IsLetter(v):
			bdigest.Letter = append(bdigest.Letter, string(v))
		case unicode.IsSpace(v):
			bdigest.Letter = append(bdigest.Letter, ", ")
		case unicode.IsNumber(v):
			bdigest.Number = append(bdigest.Number, string(v))
		case unicode.IsPunct(v):
			bdigest.Punct = append(bdigest.Punct, string(v))
		case unicode.IsSymbol(v):
			bdigest.Symbol = append(bdigest.Symbol, string(v))
		}
	}
	tokens := strings.Split(strings.Join(bdigest.Letter, ""), ", ")
	return tokens, bdigest
}
