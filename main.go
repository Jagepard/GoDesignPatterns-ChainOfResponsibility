/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	notice := &NoticeHandler{name: "NoticeHandler"}
	warning := &WarningHandler{name: "WarningHandler"}
	error := &ErrorHandler{name: "ErrorHandler"}

	notice.setNext(warning)
	warning.setNext(error)
	notice.execute()
}
