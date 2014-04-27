package nlptoken

import "testing"

func TestBasicLatin(t *testing.T) {
	s := UnicodeAlphabet(BasicLatin)
	scprt := s.cp[0].readtyp
	scputyp := s.cp[0].utyp
	scporder := s.cp[0].order

	var read_type string = "Basic Latin"
	var u_type UnicodeType = 10
	var rune_order = []rune{0, 1023}

	if scprt != read_type {
		t.Log("Expected", read_type, " but got", scprt)
		t.Fail()
	}

	if scputyp != u_type {
		t.Log("Expected", u_type, " but got", scputyp)
		t.Fail()
	}

	if scporder[0] != rune_order[0] {
		t.Log("Expected", rune_order, " but got", scporder)
		t.Fail()
	}

	if scporder[1] != rune_order[1] {
		t.Log("Expected", rune_order, " but got", scporder)
		t.Fail()
	}
}

func TestCyrillicNotGeneralPunctuation(t *testing.T) {
	s := UnicodeAlphabet(Cyrillic)
	scprt := s.cp[0].readtyp
	scputyp := s.cp[0].utyp
	scporder := s.cp[0].order

	var read_type string = "General Punctuation"
	var u_type UnicodeType = 18
	var rune_order = []rune{8192, 9125}

	if scprt == read_type {
		t.Log("Unequal Expected", read_type, " but got", scprt)
		t.Fail()
	}

	if scputyp == 18 {
		t.Log("Unequal Expected", u_type, " but got", scputyp)
		t.Fail()
	}

	if scporder[0] == rune_order[0] {
		t.Log("Unequal Expected", rune_order, " but got", scporder)
		t.Fail()
	}

	if scporder[1] == rune_order[1] {
		t.Log("Unequal Expected", rune_order, " but got", scporder)
		t.Fail()
	}
}
