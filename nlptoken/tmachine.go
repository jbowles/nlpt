package nlptoken

// Tmachine is a Tokenizer State Machine based on Denis Papathanasiou's 
// article http://denis.papathanasiou.org/2013/02/10/state-machines-in-go-golang/

type TokenHandler func(interface{}) (string, interface{})

type Tmachine struct {
	TokenHandlers	map[string]TokenHandler
	StartState	string
	EndStates	map[string]bool
}

func (machine *Tmachine) AddState(handlerName string, handlerFn TokenHandler) {
	machine.TokenHandlers[handlerName] = handlerFn
}

func (machine *Tmachine) AddEndState(endState string) {
	machine.EndStates[endState] = true
}

func (machine *Tmachine) Execute(cargo interface{}) {
	if handler, present := machine.TokenHandlers[machine.StartState]; present {
		for {
			nextState, nextCargo := handler(cargo)
			_, finished := machine.EndStates[nextState]
			if finished {
				break
			} else {
				handler, present = machine.TokenHandlers[nextState]
				cargo = nextCargo
			}
		}
	}
}
