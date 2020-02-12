package main

// ChainInterface is ...
type ChainInterface interface {
	addToChain(Handler)
	execute(int)
}

func addToChain(chain ChainInterface, handler Handler) {
	chain.addToChain(handler)
}

func execute(chain ChainInterface, handlerPriority int) {
	chain.execute(handlerPriority)
}
