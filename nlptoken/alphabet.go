// Copyright ©2013 The rivet Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TokenBase interface defines the base token alphabet for functions of the rvt_tokenizer
package nlptoken

// Interface to Compose() with different sets of rune Structs for tokenizing text
// It's been locked down to 2 rune slices per struct; this may change.
type TokenBase interface {
	build() ([]rune, []rune)
}

// Parametric function to group all rune Structs into one rune slice
func Alphabet(tks ...TokenBase) []rune {
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

// The Cyrillic letter Struct for the tokenizer
// Languge Ex: Russian, Belarusian, Ukrainian, Rusyn, Serbian, Bulgarin, Macedonian, Chechen, and other Slavic langauges.
type Cyrillic struct {
	Upper, Lower []rune
}

// Small set of Mandarin characters and pinyin for the tokenizer
// Languge Ex: Chinese Mandarin basic ideograms and pinyin
type Mandarin struct {
	Ideogram, Pinyin []rune
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
func (c *Cyrillic) build() ([]rune, []rune) {
	c.Lower = []rune("аӑӓәӛӕбвгґѓғӷҕдђеѐёӗҽҿєжӂҗӝзҙӟӡѕиѝӥӣіїӀйҋјкқҟҡӄҝлӆљмӎнӊңӈҥњоӧөӫҩпҧрҏсҫтҭћќуўӳӱӯүұфхҳһцҵчӵҷӌҹџшщъыӹьҍэӭюя")
	c.Upper = []rune("АӐӒӘӚӔБВГҐЃҒӶҔДЂЕЀЁӖҼҾЄЖӁҖӜЗҘӞӠЅИЍӤӢІЇӀЙҊЈКҚҞҠӃҜЛӅЉМӍНӉҢӇҤЊОӦӨӪҨПҦРҎСҪТҬЋЌУЎӲӰӮҮҰФХҲҺЦҴЧӴҶӋҸЏШЩЪЫӸЬҌЭӬЮЯ")
	return c.Lower, c.Upper
}
