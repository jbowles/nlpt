/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

//nlptoken package implements various ways to tokenize natural language text.
package nlptoken

import "strings"

//Readable makes slices of strings a bit easier to read by joining the tokens by read separator
func Readable(tokens []string, separator string) (readableTok []string) {
	for _, t := range tokens {
		t = t + separator
		readableTok = append(readableTok, t)
	}
	return readableTok
}

//WhiteSpace uses strings package Split() with a read separator for more friendly humna reading of slices
func WhiteSpace(text ...string) (words []string) {
	words = strings.Split(text[0], " ")
	if len(text) == 2 {
		return Readable(words, text[1])
	} else {
		return
	}
}
