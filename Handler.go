/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// Handler is ...
type Handler struct {
	priority int
	name     string
}

func (handler Handler) exec() {
	fmt.Println(handler.name)
}
