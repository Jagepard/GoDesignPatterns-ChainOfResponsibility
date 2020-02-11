package main

// Chain is ...
type Chain struct {
	chain map[int]Handler
}

func (chain Chain) addToChain(handler Handler) {
	chain.chain[handler.priority] = handler
}
