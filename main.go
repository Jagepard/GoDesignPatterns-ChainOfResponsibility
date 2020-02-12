package main

func main() {
	chain := Chain{chain: make(map[int]Handler)}

	noticeHandler := Handler{priority: 1, name: "NoticeHandler"}
	warningHandler := Handler{priority: 2, name: "WarningHandler"}
	errorHandler := Handler{priority: 3, name: "ErrorHandler"}

	addToChain(chain, noticeHandler)
	addToChain(chain, warningHandler)
	addToChain(chain, errorHandler)

	execute(chain, 1)
	execute(chain, 2)
	execute(chain, 3)
}
