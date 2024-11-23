package browser

import "fmt"

func newDomError(msg string) error {
	return fmt.Errorf("DOMError: " + msg)
}
