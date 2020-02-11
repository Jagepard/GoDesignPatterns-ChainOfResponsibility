package main

import "fmt"

func main() {
	chain := Chain{chain: make(map[int]Handler)}

	noticeHandler := Handler{priority: 1, name: "NoticeHandler"}
	warningHandler := Handler{priority: 2, name: "WarningHandler"}
	errorHandler := Handler{priority: 3, name: "ErrorHandler"}

	chain.addToChain(noticeHandler)
	chain.addToChain(warningHandler)
	chain.addToChain(errorHandler)

	fmt.Println(chain.chain)
}
