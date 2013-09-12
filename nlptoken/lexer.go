/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

package nlptoken

import (
	"fmt"
	//"strings"
	"unicode/utf8"
)

// create a new type struct that has CodePoint in it and utype

type TokenRange struct {
	cp     []CodePoint
	uniset []rune
}

type CodePoint struct {
	order   []rune
	utyp    unicodeType
	readtyp string
}

type unicodeType int

type item struct {
	typ unicodeType // Type, such as itemBasicLatin
	val string      // Value, such as "C"
}

const (
	itemEOF unicodeType = iota
	itemEOS
	itemSpaceMeta
	itemError
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

func UnicodeSet(sets ...CodePoint) TokenRange {
	t := TokenRange{uniset: make([]rune, 0), cp: make([]CodePoint, 0)}

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

type Lexer struct {
	name  string      // used only for error reports.
	input string      // the string being scanned.
	start int         // start position of this item.
	pos   int         // current position in the input.
	width int         // width of last rune read from input.
	state unicodeType // unicodeType
	items chan item   // channel of scanned items.

}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFunc func(*Lexer) stateFunc

// lex initializes itself to lex a string and launches the state machine as a goroutine, returning the lexer and a channel of items
func lexify(name, input string) (*Lexer, chan item) {
	l := &Lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run() //Concurrently run state machine
	fmt.Println("hello lexify")
	return l, l.items
}

/*
* run lexes the input by executing state functions until
*  the state is nil.
* lexer begins by looking for plain text:
*  initial state is the function lexToken()
*  It absorbs plain text until "character" is encountered
 */
func (l *Lexer) run() {
	//notice that since lexText is already in run() scope it will have access to the lexer pointer
	for state := lexToken; state != nil; {
		state = state(l)
		fmt.Println("hello run")
	}
	close(l.items)
}

// emit passes items back via item channel
func (l *Lexer) emit(t unicodeType) {
	//fmt.Println(l.input[l.start:l.pos])
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
	fmt.Println("hello emit")
}

func lexToken(l *Lexer) (s stateFunc) {
	for {
		if l.pos > l.start {
			l.emit(itemBasicLatin)
			fmt.Println("hello lextoken")
		}
		return lexNext
	}
	// Correctly reached EOF.
	if l.pos > l.start {
		l.emit(itemBasicLatin)
	}
	l.emit(itemEOF)
	return nil
}

func lexNext(l *Lexer) stateFunc {
	l.pos += 1
	l.emit(itemBasicLatin)
	return lexLexer
}

func lexLexer(l *Lexer) stateFunc {
	for {
		switch r := l.next(); {
		case r <= BasicLatin.order[0] || r >= BasicLatin.order[1023]:
			l.emit(itemBasicLatin)
		}
	}
}

// next returns the next rune in the input.
func (l *Lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return r
	}
	r, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}
