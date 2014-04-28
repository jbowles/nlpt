#nlptoken

## Benchmarking and Profiling
Without interface `Tokenizer` and tokenizer Digests:

```sh
BenchmarkUnicTokenGoodStr	        100000	      18441 ns/op	  12664 B/op	  187 allocs/op
BenchmarkUnicTokenBucketBadStr	  200000	      12400 ns/op	  7249 B/op	    133 allocs/op
BenchmarkWhiteSpace	              2000000000	  0.00 ns/op	  0 B/op	      0 allocs/op
```

With interface `Tokenizer` and Digests:

```sh
BenchmarkBuktTknzGoodStr	  100000	      18553 ns/op	    12808 B/op	    188 allocs/op
BenchmarkBuktTnkzBadStr	    200000	      12732 ns/op	    7374 B/op	      134 allocs/op
BenchmarkWhiteSpace	        2000000000	  0.00 ns/op	    0 B/op	        0 allocs/op

# new tokenizer
BenchmarkLexTknzGoodStr	    50000	        54472 ns/op	    6472 B/op	      178 allocs/op
BenchmarkLexTknzBadStr	    50000	        47198 ns/op	    6201 B/op	      167 allocs/op
```

## Whitespace

```go
package main

import (
	"fmt"
	"github.com/jbowles/nlpt/nlptoken"
)

const (
	BadStr = "expect0.7rant7! Then I want to show Snow White and the Seven Dwarves. <=AndThe start of a new sentence. And\n then\n\nagain for One and NASA?"
)

func main() {
	tok := nlptoken.WhiteSpace(BadStr)
	fmt.Printf("WhiteSpace on Bad String: %v", tok)
}
/*
WhiteSpace on Bad String: [expect0.7rant7! Then I want to show Snow White and the Seven Dwarves. <=AndThe start of a new sentence. And
 then

again for One and NASA?]
*/
```


##UnicToken

```go
package main

import (
	"fmt"
	"github.com/jbowles/nlpt/nlptoken"
)

const (
	BadStr = "expect0.7rant7! Then I want to show Snow White and the Seven Dwarves. <=AndThe start of a new sentence. And\n then\n\nagain for One and NASA?"
)

func main() {
	ut, _ := nlptoken.UnicToken(BadStr)
	fmt.Printf("Unicode Token on Bad String: %v", ut)
}

/*
Unicode Token on Bad String: [expectrant Then I want to show Snow White and the Seven Dwarves AndThe start of a new sentence And  then  again for One and NASA]
*/
```

## Punkt
## Lexer
```go
// Copyright ©2013 The rivet Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package nlptoken

/*
// Following Rob Pike's discussion of building a lexer that
// "Executes an action, returns the next state—as a state function."
// http://cuddle.googlecode.com/hg/talk/lex.html#landing-slide
// something like:

SEE https://groups.google.com/forum/#!topic/golang-nuts/DcONCTJUDhY for discussion
NOTE: as of Go 1 we can use goroutines with init() functions

LEXER IN http://weekly.golang.org/src/pkg/text/template/parse/lex.go

	// lexer holds the state of the scanner.
	type lexer struct {
		name  string    // used only for error reports.
		input string    // the string being scanned.
		start int       // start position of this item.
		pos   int       // current position in the input.
		width int       // width of last rune read from input.
		items chan item // channel of scanned items.
	}
	// stateFn represents the state of the scanner
	// as a function that returns the next state.
	type stateFn func(*lexer) stateFn

	// run lexes the input by executing state functions
	// until the state is nil.
	func run() {
		for state := startState; state != nil; {
			state = state(lexer)
		}
	}

	// run lexes the input by executing state functions until
	// the state is nil.
	func (l *lexer) run() {
		for state := lexText; state != nil; {
			state = state(l)
		}
		close(l.items) // No more tokens will be delivered.
	}

	func lex(name, input string) (*lexer, chan item) {
		l := &lexer{
			name:  name,
			input: input,
			items: make(chan item),
		}
		go l.run()  // Concurrently run state machine.
		return l, l.items
	}

	// emit passes an item back to the client.
	func (l *lexer) emit(t itemType) {
		l.items <- item{t, l.input[l.start:l.pos]}
		l.start = l.pos
	}
	const leftMeta = "{{"

	func lexText(l *lexer) stateFn {
		for {
			if strings.HasPrefix(l.input[l.pos:], leftMeta) {
				if l.pos > l.start {
					l.emit(itemText)
				}
				return lexLeftMeta    // Next state.
			}
			if l.next() == eof { break }
		}
		// Correctly reached EOF.
		if l.pos > l.start {
			l.emit(itemText)
		}
		l.emit(itemEOF)  // Useful to make EOF a token.
		return nil       // Stop the run loop.
	}

	func lexLeftMeta(l *lexer) stateFn {
		l.pos += len(leftMeta)
		l.emit(itemLeftMeta)
		return lexInsideAction    // Now inside {{ }}.
	}

	func lexInsideAction(l *lexer) stateFn {
		// Either number, quoted string, or identifier.
		// Spaces separate and are ignored.
		// Pipe symbols separate and are emitted.
		for {
			if strings.HasPrefix(l.input[l.pos:], rightMeta) {
				return lexRightMeta
			}
			switch r := l.next(); {
			case r == eof || r == '\n':
				return l.errorf("unclosed action")
			case isSpace(r):
				l.ignore()
			case r == '|':
				l.emit(itemPipe)
			case r == '"':
				return lexQuote
			case r == '`':
				return lexRawQuote
			case r == '+' || r == '-' || '0' <= r && r <= '9':
				l.backup()
				return lexNumber
			case isAlphaNumeric(r):
				l.backup()
				return lexIdentifier
			}
		}
	}

	// next returns the next rune in the input.
	func (l *lexer) next() (rune int) {
		if l.pos >= len(l.input) {
			l.width = 0
			return eof
		}
		rune, l.width =
			utf8.DecodeRuneInString(l.input[l.pos:])
		l.pos += l.width
		return rune
	}

	// ignore skips over the pending input before this point.
	func (l *lexer) ignore() {
		l.start = l.pos
	}

	// backup steps back one rune.
	// Can be called only once per call of next.
	func (l *lexer) backup() {
		l.pos -= l.width
	}

	// peek returns but does not consume
	// the next rune in the input.
	func (l *lexer) peek() int {
		rune := l.next()
		l.backup()
		return rune
	}

	// accept consumes the next rune
	// if it's from the valid set.
	func (l *lexer) accept(valid string) bool {
		if strings.IndexRune(valid, l.next()) >= 0 {
			return true
		}
		l.backup()
		return false
	}

	// acceptRun consumes a run of runes from the valid set.
	func (l *lexer) acceptRun(valid string) {
		for strings.IndexRune(valid, l.next()) >= 0 {
		}
		l.backup()
	}

	func lexNumber(l *lexer) stateFn {
		// Optional leading sign.
		l.accept("+-")
		// Is it hex?
		digits := "0123456789"
		if l.accept("0") && l.accept("xX") {
			digits = "0123456789abcdefABCDEF"
		}
		l.acceptRun(digits)
		if l.accept(".") {
			l.acceptRun(digits)
		}
		if l.accept("eE") {
			l.accept("+-")
			l.acceptRun("0123456789")
		}
		// Is it imaginary?
		l.accept("i")
		// Next thing mustn't be alphanumeric.
		if isAlphaNumeric(l.peek()) {
			l.next()
			return l.errorf("bad number syntax: %q",
				l.input[l.start:l.pos])
		}
		l.emit(itemNumber)
		return lexInsideAction
	}

	// error returns an error token and terminates the scan
	// by passing back a nil pointer that will be the next
	// state, terminating l.run.
	func (l *lexer) errorf(format string, args ...interface{})
	  stateFn {
		l.items <- item{
			itemError,
			fmt.Sprintf(format, args...),
		}
		return nil
	}


!!!!!*******************************
	With Go 1, the lexer can go back to using a goroutine. I should
probably make that change.
-rob
!!!!!*******************************

	// lex creates a new scanner for the input string.
	func lex(name, input string) *lexer {
		l := &lexer{
			name:  name,
			input: input,
			state: lexText,
			items: make(chan item, 2), // Two items sufficient.
		}
		return l
	}

	// nextItem returns the next item from the input.
	func (l *lexer) nextItem() item {
		for {
			select {
			case item := <-l.items:
				return item
			default:
				l.state = l.state(l)
			}
		}
		panic("not reached")
	}
*/
```
