/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

//nlptoken package implements various ways to tokenize natural language text.
package nlptoken

// SimpleWord tokenizer returns words spearated by comma.
// Uses WhiteSpaceTokenizer, which is a splitter on whitespace only.
// Tokenizing on whitespace is naive and will produce poor results.
// Though it can be useful in some cases, this function is not recommended
// for robust processing needs. See wstoken.go
func SimpleWord(text string) []string {
	var latin = UnicodeAlphabet(BasicLatin)
	return WhiteSpaceTokenizer(latin.uniset, text, ", ")
}
