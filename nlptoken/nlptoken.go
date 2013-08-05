// Copyright ©2013 The rivet Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//nlptoken package implements various ways to tokenize natural language text.
package nlptoken

// Accessible Alphabets
var LatinSet = SimpleAlphabet(&Latin{})
var PunctNumSet = SimpleAlphabet(&PunctNum{})
var SpaceCharSet = SimpleAlphabet(&SpaceChar{})
var EnglishSet = SimpleAlphabet(&Latin{}, &PunctNum{}, &SpaceChar{})

// Whitespace tokenizer returns words spearated by comma.
// Uses SimpleWord, which is a splitter on whitespace only.
// Tokenizing on whitespace is naive and will produce poor results.
// Though it can be useful in some cases, this function is not recommeded
// for robust processing needs. See wstoken.go
func WhiteSpace(text string) []string {
	return SimpleWord(EnglishSet, text, ", ")
}
