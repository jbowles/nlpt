// Copyright ©2013 The rivet Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package nlptoken

// Encodings is the interface for unicode.go,
// it defines a scope function for ordered unicode
// code points.
// See codepoints.go of list.
type Encodings interface {
	scope() (rune, rune)
}

// CodePoint is the struct for unicode.go,
// it contains the order of the range code points
type CodePoint struct {
	order []rune
}

// scope implements the Encodings interface through
// CodePoint structs defined over ordered range of
// Unicode Code Points.
func (cp CodePoint) scope() (startidx, stopidx rune) {
	return cp.order[0], cp.order[1]
}

// UnicodeSet builds a slice of runes based on unicode ranges
// for any Encodings struct that has the Unicode Code Points range
// defined.
// fmt.Println("BasicLatin:", UnicodeSet(BasicLatin,Cyrllic))
func UnicodeSet(sets ...Encodings) []rune {
	var uniset []rune
	for _, s := range sets {
		startidx, stopidx := s.scope()
		tmp := make([]rune, stopidx, stopidx)
		for i := startidx; i < stopidx; i++ {
			//fmt.Println(tmp[i])
			tmp[i] = rune(i)
		}
		uniset = append(uniset, tmp[startidx:]...)
	}
	return uniset
}
