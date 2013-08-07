/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.

* An Alphabet is a range over the Unicode set. Since this 
* package uses runes the set is limited to characters of 4 bytes
* or less (the limit of the Rune type).
*/
package nlptoken

// Encodings is the preferred way to build the alphabet.
// It defines a scope function for ordered unicode code points.
// See codepoints.go of list.
type Encodings interface {
	scope() (rune, rune)
}

// CodePoint is the struct for unicode.go,
// it contains the order of the range code points
// and implement the Encodings interface.
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
// fmt.Println("BasicLatin:", UnicodeAlphabet(BasicLatin,Cyrllic))
func UnicodeAlphabet(sets ...Encodings) []rune {
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
