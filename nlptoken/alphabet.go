/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* An Alphabet is a range over the Unicode set. Since this
* package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*
* SimpleAlphabet and TokenBase:	    are deprecated but kept around for some historical reasons
* UniAlph and TokRange:	are the preferred way to build an alphabet.
 */
package nlptoken

//////////////////////////////////////////////////
/////////  LEXER (go-style lexer) MODEL //////////
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

/* UniAlph builds a range of Unicode values with specific ordering range, unicode iota type, and a human readable readtyp
*	  It is to be used in the context of tokenizing text via checking set membership of the token to the token range
*	  Example:
*		s := UniAlph(BasicLatin,Cyrillic,GeneralPunctuation)
*		fmt.Println(s.cp)
*		for _,t := range s.uniset {
*			  fmt.Prtinf("Character: %c, Rune: %v", t, t)
*		}
 */
func UniAlph(sets ...CodePoint) TokRange {
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

/////////////////////////////////////////////////////////////////////////////////
////////////////////       DEPRECATED            ////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////
// TokenBase is the Interface to build() with different sets of rune
// structs for tokenizing text.
// It's been locked down to 2 rune slices per struct.
type TokenBase interface {
	build() ([]rune, []rune)
}

// Parametric function to group all rune Structs into one rune slice
func SimpleAlphabet(tks ...TokenBase) []rune {
	var result []rune
	//range over types, collect multiple return values, append to new slice, append final slice
	for _, tok := range tks {
		//return both rune slices
		s, t := tok.build()
		// append both rune slices and assign to new slice
		u := append(s, t...)
		//final append to result slice
		result = append(result, u...)
	}
	return result
}

// The Latin Struct for the tokenizer
// Languge Ex: English
type Latin struct {
	Lower, Upper []rune
}

// The LatinLigature Struct defines a common set extra unicode characters for the Latin alphabet
// A Ligature is the combination of two or more graphemes/characters
// NOTE: For Slavic languages using the Cyrillic Struct
// Languge Ex: French
type LatinLigature struct {
	LigatureUpper, LigatureLower []rune
}

// The LatinDiacritc Struct defines a common set extra unicode characters for the Latin alphabet
// A Diacritic is mark above, below, through, or between alphabetic characters generally signfying pronunciation
// NOTE: For Slavic languages use the Cyrillic Struct
// Languge Ex: Portuguese (Brazilian and Portugal), Spanish, French, German
type LatinDiacritic struct {
	DiacriticUpper, DiacriticLower []rune
}

// The Punctiation and Number Struct for the tokenizer
type PunctNum struct {
	Punctuation, Number []rune
}

// The Space and various Characters Struct for the tokenizer
type SpaceChar struct {
	Space, Character []rune
}

// The Tab and New Line Struct for the tokenizer
type TabLine struct {
	Tab, NewLine []rune
}

// The Greek letter Struct for the tokenizer
// Languge Ex: Greek
type Greek struct {
	Upper, Lower []rune
}

// The CyrillicSimple letter Struct for the tokenizer
// Languge Ex: Russian, Belarusian, Ukrainian, Rusyn, Serbian, Bulgarin, Macedonian, Chechen, and other Slavic langauges.
type CyrillicSimple struct {
	Upper, Lower []rune
}

// build() methods for TokenBase interface using pointers.

func (l *Latin) build() ([]rune, []rune) {
	l.Lower = []rune("abcdefghijklmnopqrstuvwxyz")
	l.Upper = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return l.Lower, l.Upper
}
func (llg *LatinLigature) build() ([]rune, []rune) {
	llg.LigatureLower = []rune("œæ")
	llg.LigatureUpper = []rune("ŒÆ")
	return llg.LigatureLower, llg.LigatureUpper
}
func (ld *LatinDiacritic) build() ([]rune, []rune) {
	ld.DiacriticLower = []rune("áâãàäçèéêíóôõöúüñÿ")
	ld.DiacriticUpper = []rune("ÁÄÂÃÀßÇÉÊÍÓÔÕÖÚÜÑŸ")
	return ld.DiacriticLower, ld.DiacriticUpper
}
func (pn *PunctNum) build() ([]rune, []rune) {
	pn.Punctuation = []rune(".,?!")
	pn.Number = []rune("0123456789")
	return pn.Punctuation, pn.Number
}
func (sc *SpaceChar) build() ([]rune, []rune) {
	sc.Space = []rune(" ")
	sc.Character = []rune("\"'`@#$-+=/*&%[]{}()")
	return sc.Character, sc.Space
}
func (tl *TabLine) build() ([]rune, []rune) {
	tl.NewLine = []rune("\n")
	tl.Tab = []rune("\t")
	return tl.Tab, tl.NewLine
}
func (g *Greek) build() ([]rune, []rune) {
	g.Lower = []rune("αβγδεζηθϑικλμνξοπϖρςστυφχψω")
	g.Upper = []rune("ΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥϒΦΧΨΩ")
	return g.Lower, g.Upper
}
func (c *CyrillicSimple) build() ([]rune, []rune) {
	c.Lower = []rune("аӑӓәӛӕбвгґѓғӷҕдђеѐёӗҽҿєжӂҗӝзҙӟӡѕиѝӥӣіїӀйҋјкқҟҡӄҝлӆљмӎнӊңӈҥњоӧөӫҩпҧрҏсҫтҭћќуўӳӱӯүұфхҳһцҵчӵҷӌҹџшщъыӹьҍэӭюя")
	c.Upper = []rune("АӐӒӘӚӔБВГҐЃҒӶҔДЂЕЀЁӖҼҾЄЖӁҖӜЗҘӞӠЅИЍӤӢІЇӀЙҊЈКҚҞҠӃҜЛӅЉМӍНӉҢӇҤЊОӦӨӪҨПҦРҎСҪТҬЋЌУЎӲӰӮҮҰФХҲҺЦҴЧӴҶӋҸЏШЩЪЫӸЬҌЭӬЮЯ")
	return c.Lower, c.Upper
}
