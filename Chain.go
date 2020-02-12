package main

import "fmt"
import "errors"

// Chain is ...
type Chain struct {
	chain map[int]Handler
}

func (chain Chain) addToChain(handler Handler) {
	if val, ok := chain.chain[handler.priority]; ok {
		fmt.Println(errors.New(val.name + ":already exists"))
		panic("Wrong argument!")
	}

	chain.chain[handler.priority] = handler
}

func (chain Chain) execute(handlerPriority int) {
	if handler, ok := chain.chain[handlerPriority]; ok {
		for i := 1; i <= handler.priority; i++ {
			fmt.Println(chain.chain[i].name)
		}

		return
	}

	// fmt.Println(errors.New(handlerPriority + ": does not exist in the chain"))
}
