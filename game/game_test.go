package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_advanceEmptyBoard(t *testing.T) {
	b := make(map[Pos]bool)

	board := Advance(b)

	assert.Equal(t, map[Pos]bool{}, board)
}

func Test_advanceBoardWith1Cell(t *testing.T) {
	b := make(map[Pos]bool)
	b[Pos{1, 1}] = true
	board := Advance(b)
	assert.Equal(t, map[Pos]bool{}, board)
}

func Test_advanceBoardWith3Cells(t *testing.T) {
	b := make(map[Pos]bool)
	b[Pos{1, 1}] = true
	b[Pos{1, 2}] = true
	b[Pos{1, 3}] = true
	board := Advance(b)
	assert.Equal(t, map[Pos]bool{{0, 2}: true, {1, 2}: true, {2, 2}: true}, board)

	board = Advance(board)
	assert.Equal(t, map[Pos]bool{{1, 1}: true, {1, 2}: true, {1, 3}: true}, board)
}

func Test_advanceToCube(t *testing.T) {
	b := make(map[Pos]bool)
	b[Pos{1, 1}] = true
	b[Pos{1, 2}] = true
	b[Pos{2, 1}] = true
	b[Pos{2, 2}] = false
	board := Advance(b)
	assert.Equal(t, map[Pos]bool{{1, 1}: true, {1, 2}: true, {2, 1}: true, {2, 2}: true}, board)

	board = Advance(board)
	assert.Equal(t, map[Pos]bool{{1, 1}: true, {1, 2}: true, {2, 1}: true, {2, 2}: true}, board)
}
