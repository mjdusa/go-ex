package ext

import (
	"fmt"
)

func WrapError(message string, err error) error {
	if len(message) > 0 {
		return fmt.Errorf("%s: %w", message, err)
	}
	return fmt.Errorf("%w", err)
}
