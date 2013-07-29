// Copyright ©2013 The NLPT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


// Very simple natural language tokenizer provides default word tokenizer and custom EachCharacter tokenizer for different sets of unicode alphabets. Pre-defined alphabets for Western Indo-European unicode as well as Slavic and Greek. API allows easy extension of unicode sets for custom alphabets. 
package nlpttoken

import (
	"strings"
)

//runeSep joins okens by a comma for easier reading
func runeSep(tokens []string, separator string) (readable_tokens []string) {
	for _, t := range tokens {
		t = t + separator
		readable_tokens = append(readable_tokens, t)
	}
	return
}

// define a set of vars for the default tokenizer
var latin = Latin{}
var punctuation = PunctNum{}
var space_char = SpaceChar{}
var english = Alphabet(&latin, &punctuation, &space_char)


//////////////////////////////////////////////////
//////////  NAIVE (Whitespace) MODEL /////////////
//////////////////////////////////////////////////

// Whitespace tokenizer returns words spearated by comma.
// Uses SimpleWord, which is a splitter on whitespace only.
// Tokenizing on whitespace is naive and will produce poor results.
// Though it can be useful in some cases, this function is not recommeded for robust processing needs.
func WhiteSpace(text string) []string {
	return SimpleWord(english, text, ", ")
}

// EachCharacter converts string into an array of character tokens
// by looping thorugh input and checking if they are part of the 
// allowed alphabet; if not they are discarded.
func EachCharacter(okRunes []rune, line, separator string) (tokens []string) {
	//make map of allowable characters
	var okChars = make(map[rune]bool)
	for _, okrn := range okRunes {
		//rune true if acceptable
		okChars[okrn] = true
	}
	//iterate runes and append given acceptable or not
	for _, rn := range line {
		if okChars[rn] {
			//convert rune to string and append to token slice
			tokens = append(tokens, string(rn))
		} else {
			//unnacceptable tokens use negation-sqrt placeholder
			tokens = append(tokens, "¬")
		}
	}
	return runeSep(tokens, separator)
}

// SimpleWord converts strings into slice of word tokens
// by splitting on whitespace only. It calls runeSep and accepts
// any type of word separator (defualt is a comma).
func SimpleWord(okRunes []rune, line, separator string) (words []string) {
	words = strings.Split(line, " ")
	return runeSep(words, separator)
}
