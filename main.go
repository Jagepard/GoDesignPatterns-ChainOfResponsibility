package main

func main() {
	chain := Chain{chain: make(map[int]Handler)}

	notice := NoticeHandler{Handler{priority: 1, name: "NoticeHandler"}}
	warning := WarningHandler{Handler{priority: 2, name: "WarningHandler"}}
	error := ErrorHandler{Handler{priority: 3, name: "ErrorHandler"}}

	addToChain(chain, notice.Handler)
	addToChain(chain, warning.Handler)
	addToChain(chain, error.Handler)

	execute(chain, 1)
	execute(chain, 2)
	execute(chain, 3)
}
