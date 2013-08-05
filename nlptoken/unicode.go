// Copyright ©2013 The nlpt Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package nlptoken

// Define the basic CodePoint orders for a crap ton of Unicode!!
// See http://www.utf8-chartable.de/ for a table of code points.
var (
	BasicLatin							= CodePoint{order: []rune{0, 1024}}		  //
	Cyrllic								= CodePoint{order: []rune{1024, 2047}}	  //
	Samaritan							= CodePoint{order: []rune{2048, 3071}}	  //
	Telugu								= CodePoint{order: []rune{3072, 4095}}	  //
	Myanmar								= CodePoint{order: []rune{4096, 5119}}	  //
	UnifiedCanadianAboriginalSyllabics	= CodePoint{order: []rune{5120, 6143}}	  //
	Mongolian							= CodePoint{order: []rune{6144, 7167}}	  //
	Lepcha								= CodePoint{order: []rune{7168, 8191}}	  //
	GeneralPunctuation					= CodePoint{order: []rune{8192, 9125}}	  //
	ControlPictures						= CodePoint{order: []rune{9216, 10239}}	  //
	BraillePatterns						= CodePoint{order: []rune{10240, 11263}}  //
	Glagolitic							= CodePoint{order: []rune{11264, 12287}}  //
	CjkSymbolsPunctuation				= CodePoint{order: []rune{12288, 13311}}  // Chinese, Japanese, Korean
	CjkUnifiedIdeographsExtA			= CodePoint{order: []rune{13312, 20479}}  // Chinese, Japanese, Korean
	CjkUnifiedIdeographs				= CodePoint{order: []rune{20480, 40959}}  // Chinese, Japanese, Korean
	YiSyllables							= CodePoint{order: []rune{0, 0}}  //
)
