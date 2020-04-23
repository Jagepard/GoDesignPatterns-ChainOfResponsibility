/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	notice := &NoticeHandler{name: "NoticeHandler"}

	notice.setNext(&WarningHandler{name: "WarningHandler"}).setNext(&ErrorHandler{name: "ErrorHandler"})
	notice.execute("NoticeHandler")
	notice.execute("WarningHandler")
	notice.execute("ErrorHandler")
}
