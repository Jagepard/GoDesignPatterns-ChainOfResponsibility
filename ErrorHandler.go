/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// ErrorHandler is ...
type ErrorHandler struct {
	name string
	next HandlerInterface
}

func (E *ErrorHandler) execute(request string) {

	if request == E.name {
		fmt.Println(E.name)
		return
	}

	if E.next != nil {
		E.next.execute(request)
	}
}

func (E *ErrorHandler) setNext(next HandlerInterface) HandlerInterface {
	E.next = next
	return next
}
