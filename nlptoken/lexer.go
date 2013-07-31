package nlptoken

// Following Rob Pike's discussion of building a lexer that
// "Executes an action, returns the next state—as a state function."
// something like:

/*
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
*/
