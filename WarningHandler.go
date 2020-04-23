/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// WarningHandler is ...
type WarningHandler struct {
	name string
	next HandlerInterface
}

func (W *WarningHandler) execute(request string) {

	if request == W.name {
		fmt.Println(W.name)
		return
	}

	if W.next != nil {
		W.next.execute(request)
	}
}

func (W *WarningHandler) setNext(next HandlerInterface) HandlerInterface {
	W.next = next
	return next
}
