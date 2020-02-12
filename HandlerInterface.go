package main

// HandlerInterface is ...
type HandlerInterface interface {
	exec()
}

func exec(handler HandlerInterface) {
	handler.exec()
}
