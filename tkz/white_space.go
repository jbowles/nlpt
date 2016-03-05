/*
* Copyright ©2015 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

package tkz

import "strings"

//NewWhiteSpaceDigest intitializes a digest for white space tokenization.
func NewWhiteSpaceDigest() *Digest {
	return &Digest{
		Tokens:     make([]string, 0, 0),
		CharCount:  0,
		SpaceCount: 0,
	}
}

//Tknz implements the Tokenizer interface. This uses the strings package Split() with a white space separator as well as collecting some other metadata for the digest.
func TknzWhiteSpace(text string, digest *Digest) ([]string, *Digest) {
	digest.Tokens = strings.Split(text, " ")
	digest.SpaceCount = strings.Count(text, " ")
	digest.CharCount = len(text)

	return digest.Tokens, digest
}
