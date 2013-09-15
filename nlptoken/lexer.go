/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	//"strings"
	"unicode/utf8"
)

// create a new type struct that has CodePoint in it and utype

type item struct {
	typ unicodeType // Type, such as itemBasicLatin
	val string      // Value, such as "C"
}

const whiteSpaceMeta = " "

func (i item) String() string {
	if len(i.val) > 100 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

type Lexer struct {
	name   string      // used only for error reports.
	input  string      // the string being scanned.
	output []rune      // output from tokenizing
	start  int         // start position of this item.
	pos    int         // current position in the input.
	width  int         // width of last rune read from input.
	state  unicodeType // unicodeType
	items  chan item   // channel of scanned items.

}

// stateFn represents the state of the scanner
// as a function that returns the next state.
type stateFunc func(*Lexer) stateFunc

// lex initializes itself to lex a string and launches the state machine as a goroutine, returning the lexer and a channel of items
func lexify(name, input string) (*Lexer, chan item) {
	l := &Lexer{
		name:   name,
		input:  input,
		output: make([]rune, 0),
		items:  make(chan item),
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
	defer close(l.items)
	//notice that since lexText is already in run() scope it will have access to the lexer pointer
	fmt.Println("hello run func", l.pos, l.start, l.input)
	for state := lexToken(l); state != nil; {
		fmt.Println("hello run state=", state, l.pos, l.start, l.input)
		state = state(l)
		fmt.Println("hello run state2=", state, l.pos, l.start, l.input)
	}
}

// emit passes items back via item channel
func (l *Lexer) emit(t unicodeType) {
	l.items <- item{t, l.input[l.start:l.pos]}
	//l.items <- item{t, l.input[l.start:l.pos], make([]string, len(l.input)+1)}
	l.start = l.pos
	fmt.Println("hello emit")
}

func lexToken(l *Lexer) stateFunc {
	for {
		fmt.Println("hello lextoken in for loop")
		if []rune(l.input[l.pos:])[0] == int32(32) {
			fmt.Println("hello lextoken has prefix spaceMeta state")
			if l.pos > l.start {
				l.emit(itemBasicLatin)
				fmt.Println("hello lextoken in spaceMeta state")
			}
			return lexNext // Next state.
		}
		return lexNext // Next state.
	}

	l.emit(itemEOF) // Useful to make EOF a token.
	return nil      // Stop the run loop.
}

func lexNext(l *Lexer) stateFunc {
	l.pos += 1
	l.emit(itemBasicLatin)
	return lexLexer
}

func lexLexer(l *Lexer) stateFunc {
	for {
		switch r := l.next(); {
		case utf8.ValidRune(r):
			fmt.Printf("hello rune %v\n", r)
			l.output = append(l.output, r)
			return lexNext
			/*
				case r >= 0 || r <= 1023: //<= BasicLatin.order[0] || r >= BasicLatin.order[1]:
					fmt.Println("hello stateFunc next:", r)
					l.emit(itemBasicLatin)
					return lexNext
			*/
		}
	}
}

// next returns the next rune in the input.
func (l *Lexer) next() (r rune) {
	fmt.Println("hello next():", l.pos, l.width, r)
	if l.pos >= len(l.input) {
		fmt.Println("hello next() position greater than input:", l.pos, l.width, r)
		l.width = 0
		return r
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	fmt.Printf("NEXT: rune = %c, width = %v\n", r, l.width)
	l.pos += l.width
	return r
}

func main() {
	//th := UnicodeSet(BasicLatin)
	//fmt.Println("Code Points", th.cp)
	l, lex_items := lexify("error?", "parse this string")
	fmt.Println(l.name, l.input, l.output)
	s := <-lex_items
	fmt.Println(l.name, l.input, l.output)
	//receiveItems(lex_items)
	fmt.Println("your string is parsed", s)
	//s.String()

	/*
		l, _ := lexify("error", "parse")
		fmt.Println(l.name, l.input)
		all := <-l.items
		fmt.Println(all)

		for _, char := range th.uniset {
			fmt.Println("value: ", string(char))
		}
	*/
}
