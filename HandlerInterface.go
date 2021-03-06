/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// HandlerInterface is ...
type HandlerInterface interface {
	execute(string)
	setNext(HandlerInterface) HandlerInterface
}
