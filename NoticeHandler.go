/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// NoticeHandler is ...
type NoticeHandler struct {
	name string
	next HandlerInterface
}

func (N *NoticeHandler) execute(request string) {

	if request == N.name {
		fmt.Println(N.name)
		return
	}

	if N.next != nil {
		N.next.execute(request)
	}
}

func (N *NoticeHandler) setNext(next HandlerInterface) HandlerInterface {
	N.next = next
	return next
}
