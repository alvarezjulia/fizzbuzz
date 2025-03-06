package domain

import "fmt"

type ErrInvalidRequest struct {
	Field  string
	Reason string
}

func (e ErrInvalidRequest) Error() string {
	return fmt.Sprintf("invalid request: %s %s", e.Field, e.Reason)
}
