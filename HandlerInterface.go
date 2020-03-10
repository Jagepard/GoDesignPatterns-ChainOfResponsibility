/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// HandlerInterface is ...
type HandlerInterface interface {
	exec()
}

func exec(handler HandlerInterface) {
	handler.exec()
}
