package helpers

import "fmt"

// Wrap оборачивает ошибки для прокидывания наверх по стеку вызова
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}
