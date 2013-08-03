package nlptoken

import (
	"fmt"
)

// Tmachine is a Tokenizer State Machine based on Denis Papathanasiou's 
// article http://denis.papathanasiou.org/2013/02/10/state-machines-in-go-golang/

type TokHandler func(interface{}) (string, interface{})

type Tmachine struct {
	TokHandlers	map[string]TokHandler
	StartState	string
	EndStates	map[string]bool
}

func (machine *Tmachine) AddState(handlerName string, handlerFn TokHandler) {
	machine.TokHandlers[handlerName] = handlerFn
}

func (machine *Tmachine) AddEndState(endState string) {
	machine.EndStates[endState] = true
}

func (machine *Tmachine) Execute(cargo interface{}) {
	if handler, present := machine.TokHandlers[machine.StartState]; present {
		for {
			nextState, nextCargo := handler(cargo)
			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.TokHandlers[nextState]
				cargo = nextCargo
			}
		}
	}
}

/*
*	See nlptoken/siw.go for Alphabets of
*	LatinSet, PunctNumSet, SpaceCharSet,
*/

func doRune(rns []rune) []rune {
	return rns
}

func OnesTok() TokHandler {
	return func(token interface{}) (nextState string, nextSeq interface{}) {
		nextState = ""
		nextSeq = token.(rune)
		
		fmt.Printf("Whitespace State: %v \t",nextSeq)
		for {
			switch {
			case nextSeq.(rune) <= 0 || nextSeq.(rune) >= 30:
				nextState = "outofrange"
			case nextSeq.(rune) >= 20 && nextSeq.(rune) < 30:
				nextState = "twenties"
			case nextSeq.(rune) >= 10 && nextSeq.(rune) < 20:
				nextState = "tens"
			default:
				fmt.Printf("%v \n", nextSeq.(rune))
			}

			if len(nextState) > 0 {
				break
			}
			//nextSeq = doRune([]rune("This is a sentence. It has punctuation."))


			fmt.Printf(" >> %v\n")
		}
		return
	}
}

func TensTok() TokHandler {
	return func(token interface{}) (nextState string, nextSeq interface{}) {
		nextState = ""
		nextSeq = token.(rune)
		
		fmt.Printf("Whitespace State: %v \t",nextSeq)
		for {
			switch {
			case nextSeq.(rune) <= 0 || nextSeq.(rune) >= 30:
				nextState = "outofrange"
			case nextSeq.(rune) >= 20 && nextSeq.(rune) < 30:
				nextState = "twenties"
			case nextSeq.(rune) >= 1 && nextSeq.(rune) < 10:
				nextState = "ones"
			default:
				fmt.Printf("%v \n", nextSeq.(rune))
			}

			if len(nextState) > 0 {
				break
			}
			//nextSeq = doRune([]rune("This is a sentence. It has punctuation."))


			fmt.Printf(" >> %v\n")
		}
		return
	}
}

func TwentiesTok() TokHandler {
	return func(token interface{}) (nextState string, nextSeq interface{}) {
		nextState = ""
		nextSeq = token.(rune)
		
		fmt.Printf("Whitespace State: %v \t",nextSeq)
		for {
			switch {
			case nextSeq.(rune) <= 0 || nextSeq.(rune) >= 30:
				nextState = "outofrange"
			case nextSeq.(rune) >= 10 && nextSeq.(rune) < 0:
				nextState = "tens"
			case nextSeq.(rune) >= 1 && nextSeq.(rune) < 10:
				nextState = "ones"
			default:
				fmt.Printf("%v \n", nextSeq.(rune))
			}

			if len(nextState) > 0 {
				break
			}
			//nextSeq = doRune([]rune("This is a sentence. It has punctuation."))


			fmt.Printf(" >> %v\n")
		}
		return
	}
}
