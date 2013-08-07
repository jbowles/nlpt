/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*/

package nlptoken

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

