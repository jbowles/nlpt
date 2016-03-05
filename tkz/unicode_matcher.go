/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* A Unicode Matcher tokenizer uses unicode code point ranges to segment text.
* It filters through strings leveraging standard library functions to parse and
* collect unicode segments into a 'bucket digest' of letter, number, punctuation, space, and symbol.
* Since this package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*
 */

package tkz

import (
	"bytes"
	"strings"
	"unicode"
)

//NewStateFunDigest initializes a new BucketDigest and explicitly allocates a length and cap on all stirng slices.
func NewUnicodeMatchDigest() *Digest {
	return &Digest{
		Letter:     make([]string, 0, 0),
		Title:      make([]string, 0, 0),
		Number:     make([]string, 0, 0),
		Punct:      make([]string, 0, 0),
		Space:      make([]string, 0, 0),
		Symbol:     make([]string, 0, 0),
		Bytes:      make([]byte, 0, 0),
		TokenBytes: make(map[string][]byte),
	}
}

func NewUnicodeMatchDigestBytes() *Digest {
	return &Digest{
		Bytes: make([]byte, 0, 0),
	}
}

//Tknz implements Tokenizer interface. Here it uses Unicode package to match runes for tokenization. It can be useful for really noisy data sets. For example, a sequence like 'expect0.7rant7!' will be tokenized into 3 buckets of LETTER: 'expectrant', NUMBER: '0 7 7', and PUNCT: '. !'.
//Caution should be used, however, as there is a great amount of information loss too. Date sequences, monetary sequences, urls, or any other complex combination of unicode sequences will be bucketized.
//One use of this tokenizer is to clean up naoisy data or for post-processing of already tokenized data for specific data-mining tasks. This is not a typical tokenizer. If you want basic tokenization see the Whist (whitespace), Lext (lexical scanner), Punkt (sentence segmenter) tokenizers.

func TknzUnicode(text string, digest *Digest) ([]string, *Digest) {
	for _, v := range text {
		switch true {
		case unicode.IsTitle(v):
			digest.Title = append(digest.Title, string(v))
			digest.Letter = append(digest.Letter, string(v))
		case unicode.IsLetter(v):
			digest.Letter = append(digest.Letter, string(v))
		case unicode.IsSpace(v):
			digest.Letter = append(digest.Letter, ", ")
		case unicode.IsNumber(v):
			digest.Number = append(digest.Number, string(v))
		case unicode.IsPunct(v):
			digest.Punct = append(digest.Punct, string(v))
		case unicode.IsSymbol(v):
			digest.Symbol = append(digest.Symbol, string(v))
		}
	}
	digest.Tokens = strings.Split(strings.Join(digest.Letter, ""), ", ")
	for _, word := range digest.Tokens {
		digest.DowncaseTokens = append(digest.DowncaseTokens, strings.ToLower(word))
	}
	return digest.Tokens, digest
}

func TknzUnicodeBytes(byteSeq []byte, digest *Digest) *Digest {
	bufferCache := new(bytes.Buffer)
	bytePadding := []byte{32}
	for _, b := range byteSeq {
		runeBytes := rune(b)
		switch true {
		//case unicode.IsTitle(runeBytes):
		//	bufferCache.Write([]byte{b})
		case unicode.IsLetter(runeBytes):
			bufferCache.Write([]byte{b})
		case unicode.IsSpace(runeBytes):
			bufferCache.Write([]byte{b})
		case unicode.IsNumber(runeBytes):
			bufferCache.Write([]byte{b})
		case unicode.IsSymbol(runeBytes):
			bufferCache.Write(bytePadding)
		//	bufferCache.Write([]byte{b})
		//bufferCache.Write(bytePadding)
		case unicode.IsPunct(runeBytes):
			bufferCache.Write(bytePadding)
			//	bufferCache.Write([]byte{b})
			//bufferCache.Write(bytePadding)
		}
	}
	digest.Bytes = bufferCache.Bytes()
	return digest
}
