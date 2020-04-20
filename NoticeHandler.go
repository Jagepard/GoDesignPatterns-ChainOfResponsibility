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

func (N *NoticeHandler) execute() {
	fmt.Println(N.name)
	if N.next != nil {
		N.next.execute()
	}
}

func (N *NoticeHandler) setNext(next HandlerInterface) {
	N.next = next
}
