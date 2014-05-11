/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* A Bukt language tokenizer uses unicode code point ranges to segment text.
* It filters through strings leveraging standard library functions to parse and
* collect unicode segments into a 'bucket digest' of letter, number, punctuation, space, and symbol.
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

//Tknz implement Tokenizer interface. Here it uses Unicode package to match runes for tokenization. It can be useful for really noisy data sets. For example, a sequence like 'expect0.7rant7!' will be tokenized into 3 buckets of LETTER: 'expectrant', NUMBER: '0 7 7', and PUNCT: '. !'.
//Caution should be used, however, as there is a great amount of information loss too. Date sequences, monetary sequences, urls, or any other complex combination of unicode sequences will be bucketized.
//One use of this tokenizer is to clean up naoisy data or for post-processing of already tokenized data for specific data-mining tasks. This is not a typical tokenizer. If you want basic tokenization see the Whist (whitespace), Lext (lexical scanner), Punkt (sentence segmenter) tokenizers.

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
