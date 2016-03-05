package tkz

import "unicode"

// Punct16 returns Lo and Hi ranges for the Punctuation type as defined in go's unicode Categories RangeTable for Range16.
// Range16 represents of a range of 16-bit Unicode code points. The range runs from Lo to Hi inclusive and has the specified stride.
func Punct16() (punct []uint16) {
	for _, range16 := range unicode.Categories["P"].R16 {
		punct = append(punct, range16.Lo)
		punct = append(punct, range16.Hi)
	}
	return
}

// Punct32 returns Lo and Hi ranges for the Punctuation type as defined in go's unicode Categories RangeTable for Range32.
// Range32 represents of a range of Unicode code points and is used when one or more of the values will not fit in 16 bits. The range runs from Lo to Hi inclusive and has the specified stride. Lo and Hi must always be >= 1<<16.
func Punct32() (punct []uint32) {
	for _, range32 := range unicode.Categories["P"].R32 {
		punct = append(punct, range32.Lo)
		punct = append(punct, range32.Hi)
	}
	return
}

//Range16 represents of a range of 16-bit Unicode code points. The range runs from Lo to Hi inclusive and has the specified stride.
func PunctRunes16() (punct []rune) {
	for _, range16 := range unicode.Categories["P"].R16 {
		punct = append(punct, rune(range16.Lo))
		punct = append(punct, rune(range16.Hi))
	}
	return
}

// naive way to define some byte ranges for punction.
func PunctBytes16Lim() (punct []byte) {
	for _, rn := range PunctRunes16() {
		if rn <= 191 {
			punct = append(punct, byte(rn))
		}
	}
	return
}
