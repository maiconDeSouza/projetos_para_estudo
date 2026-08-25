package apperr

import (
	"errors"
	"fmt"
	"time"
)

type ErrDuplicatePlayer struct {
	Name   string
	Number uint
}

func (e *ErrDuplicatePlayer) Error() string {
	return fmt.Sprintf("O Jogador numero %d já existe, e é o %s", e.Number, e.Name)
}

type ErrDuplicateDate struct {
	Date time.Time
}

func (e *ErrDuplicateDate) Error() string {
	return fmt.Sprintf("Já existe uma partida nesta data: %s", e.Date.String())
}

var ErrNonExistentMatch = errors.New("jogo não encontrado")
var ErrNonExistentPlayer = errors.New("jogador não encontrado")
