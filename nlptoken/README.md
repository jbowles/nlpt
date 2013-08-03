#nlptoken
## Whitespace

## Punkt

##Tmachine

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
