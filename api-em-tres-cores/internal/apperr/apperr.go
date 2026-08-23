package apperr

import (
	"fmt"
)

type ErrDuplicatePlayer struct {
	Name   string
	Number uint
}

func (e *ErrDuplicatePlayer) Error() string {
	return fmt.Sprintf("O Jogador numero %d já existe, e é o %s", e.Number, e.Name)
}
