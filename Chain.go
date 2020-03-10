/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "errors"

// Chain is ...
type Chain struct {
	chain map[int]Handler
}

func (chain Chain) addToChain(handler Handler) {
	if val, ok := chain.chain[handler.priority]; ok {
		panic(errors.New(val.name + " :already exists"))
	}

	chain.chain[handler.priority] = handler
}

func (chain Chain) execute(handlerPriority int) {
	if handler, ok := chain.chain[handlerPriority]; !ok {
		panic(errors.New(handler.name + " :does not exist in the chain"))
	}

	for i := 1; i <= handlerPriority; i++ {
		exec(chain.chain[i])
	}
}
