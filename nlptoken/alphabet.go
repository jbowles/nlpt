/*
* THIS PACKAGE IS DEPRECATED... THE FUNCTIONALITY ALREADY EXISTS IN Go Standard Library Unicode, SEE utoken.go FOR APPROPRIATE FUNCTIONALITY.
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* An Alphabet is a range over the Unicode set. Since this
* package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*
 */
package nlptoken

//////////////////////////////////////////////////
/////  DEFINE ALPHABETS WITH A LITTLE ///////////
/// MORE METADATA THAN USING unicdoe.CaseRange ///
//////////////////////////////////////////////////

type TokRange struct {
	cp     []CodePoint
	uniset []rune
}

type CodePoint struct {
	order   []rune
	utyp    unicodeType
	readtyp string
}

type unicodeType int

const (
	itemError unicodeType = iota
	itemEOF
	itemWhiteSpace
	itemAllUpperCase
	itemMultiUppercase
	itemUpperCaseSpaceUpperCase
	itemNounCompound
	itemNounDashCompound
	itemVerbCompound
	itemVerbDashCompound
	itemBasicLatin
	itemCyrillic
	itemSamaritan
	itemTelugu
	itemMyanmar
	itemUnifiedCanadianAboriginalSyllabics
	itemMongolian
	itemLepcha
	itemGeneralPunctuation
	itemControlPictures
	itemBraillePatterns
	itemGlagolitic
	itemCjkSymbolsPunctuation
	itemCjkUnifiedIdeographsExtA
	itemCjkUnifiedIdeographs
)

var (
	BasicLatin                         = CodePoint{order: []rune{0, 1023}, utyp: itemBasicLatin, readtyp: "Basic Latin"}
	Cyrillic                           = CodePoint{order: []rune{1024, 2047}, utyp: itemCyrillic, readtyp: "Cyrillic"}
	Samaritan                          = CodePoint{order: []rune{2048, 3071}, utyp: itemSamaritan, readtyp: "Samaritan"}
	Telugu                             = CodePoint{order: []rune{3072, 4095}, utyp: itemTelugu, readtyp: "Telugu"}
	Myanmar                            = CodePoint{order: []rune{4096, 5119}, utyp: itemMyanmar, readtyp: "Myanmar"}
	UnifiedCanadianAboriginalSyllabics = CodePoint{order: []rune{5120, 6143}, utyp: itemUnifiedCanadianAboriginalSyllabics, readtyp: "Canadian Aboriginal"}
	Mongolian                          = CodePoint{order: []rune{6144, 7167}, utyp: itemMongolian, readtyp: "Mongolian"}
	Lepcha                             = CodePoint{order: []rune{7168, 8191}, utyp: itemLepcha, readtyp: "Lepcha"}
	GeneralPunctuation                 = CodePoint{order: []rune{8192, 9125}, utyp: itemGeneralPunctuation, readtyp: "General Punctuation"}
	ControlPictures                    = CodePoint{order: []rune{9216, 10239}, utyp: itemControlPictures, readtyp: "Control Pictures"}
	BraillePatterns                    = CodePoint{order: []rune{10240, 11263}, utyp: itemBraillePatterns, readtyp: "Braille"}
	Glagolitic                         = CodePoint{order: []rune{11264, 12287}, utyp: itemGlagolitic, readtyp: "Glagolitic"}
	CjkSymbolsPunctuation              = CodePoint{order: []rune{12288, 13311}, utyp: itemCjkSymbolsPunctuation, readtyp: "CjkSymbolsPunctuation"}
	CjkUnifiedIdeographsExtA           = CodePoint{order: []rune{13312, 20479}, utyp: itemCjkUnifiedIdeographsExtA, readtyp: "CjkUnifiedIdeographsExtA"}
	CjkUnifiedIdeographs               = CodePoint{order: []rune{20480, 40959}, utyp: itemCjkUnifiedIdeographs, readtyp: "CjkUnifiedIdeographs"}
)

/* UnicodeAlphabet builds a range of Unicode values with specific ordering range, unicode iota type, and a human readable readtyp
*	  It is to be used in the context of tokenizing text via checking set membership of the token to the token range
*	  Example:
*		s := UniAlph(BasicLatin,Cyrillic,GeneralPunctuation)
*		fmt.Println(s.cp)
*		for _,t := range s.uniset {
*			  fmt.Prtinf("Character: %c, Rune: %v", t, t)
*		}
 */
func UnicodeAlphabet(sets ...CodePoint) TokRange {
	t := TokRange{uniset: make([]rune, 0), cp: make([]CodePoint, 0)}

	for _, cop := range sets {
		startidx := cop.order[0]
		stopidx := cop.order[1]
		t.cp = append(t.cp, cop)
		// allocate +1 index size and max for temporary slice
		tmp := make([]rune, stopidx+1, stopidx+1)
		for i := startidx; i <= stopidx; i++ {
			//fmt.Println(tmp[i])
			tmp[i] = rune(i)
		}
		t.uniset = append(t.uniset, tmp[startidx:]...)
	}
	return t
}
