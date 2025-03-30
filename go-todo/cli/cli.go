package cli

import (
	"Sranjan0208/go-todo/board"
)

var boardInstance *board.Board

// SetBoardInstance initializes the board instance
func SetBoardInstance(b *board.Board) {
	boardInstance = b
}
