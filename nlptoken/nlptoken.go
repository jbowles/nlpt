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
func WhiteSpace(text, readSep string) (words []string) {
	words = strings.Split(text, " ")
	if readSep != "" {
		return Readable(words, readSep)
	} else {
		return
	}
}

//DefaultTokenizer uses the WhiteSpace() function with a comma ', ' read separator.
//This is what many people have gotten used to seeing so I provide it here. However, this should not be used to anything but testing or display.
func DefaultTokenizer(text, readSep string) []string {
	return WhiteSpace(text, readSep)
}

//UnicToken uses the UnicBucket unicode tokenizer. It can be very useful for really noisy data sets. See documention for for UnicBucket for more details. It returns a slice of tokenized words and a the bucket.
func UnicToken(s string) ([]string, Bucket) {
	var bucket = UnicBucket(s)
	return strings.Split(strings.Join(bucket.Letter, ""), ", "), bucket
}

//UnicTokenReadable offers a human readable slice of words given a separator. It uses the UnicBucket unicode tokenizer. It returns a slice of tokenized words and a the bucket.
func UnicTokenReadable(s, readSep string) ([]string, Bucket) {
	var bucket = UnicBucket(s)
	tok := strings.Split(strings.Join(bucket.Letter, ""), ", ")
	return Readable(tok, readSep), bucket
}
