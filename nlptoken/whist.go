/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

package nlptoken

import "strings"

type WhiteSpaceDigest struct {
	Tokens     []string
	SpaceCount int
	CharCount  int
}

func NewWhiteSpaceDigest() *WhiteSpaceDigest {
	return &WhiteSpaceDigest{
		Tokens:     make([]string, 0, 0),
		CharCount:  0,
		SpaceCount: 0,
	}
}

//WhiteSpace uses strings package Split() with a read separator for more friendly humna reading of slices
func (wdigest *WhiteSpaceDigest) Tknz(text string) ([]string, *WhiteSpaceDigest) {
	wdigest.Tokens = strings.Split(text, " ")
	wdigest.SpaceCount = strings.Count(text, " ")
	wdigest.CharCount = len(text)

	return wdigest.Tokens, wdigest
}
