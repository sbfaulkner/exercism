package queenattack

import (
	"errors"
	"fmt"
)

const testVersion = 2

// Position represents a position on a chess board
type Position struct {
	file int
	rank int
}

// CanQueenAttack determines whether or not two queens are positioned so that they can attack each other
func CanQueenAttack(white, black string) (attack bool, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = errors.New(fmt.Sprint(r))
		}
	}()

	whitePosition := getPosition(white)
	blackPosition := getPosition(black)

	if whitePosition == blackPosition {
		e = errors.New("queens cannot occupy the same position")
	}

	attack = isMatchingFile(whitePosition, blackPosition) ||
		isMatchingRank(whitePosition, blackPosition) ||
		isMatchingDiagonal(whitePosition, blackPosition)

	return
}

func getPosition(square string) Position {
	position := Position{file: int(square[0] - 'a'), rank: int(square[1] - '1')}

	if !position.isValid() {
		panic(fmt.Sprintf("invalid position - %q", square))
	}

	return position
}

func (position Position) isValid() bool {
	return position.file >= 0 && position.file <= 7 && position.rank >= 0 && position.rank <= 7
}

func isMatchingFile(p1, p2 Position) bool {
	return p1.file == p2.file
}

func isMatchingRank(p1, p2 Position) bool {
	return p1.rank == p2.rank
}

func isMatchingDiagonal(p1, p2 Position) bool {
	return abs(p1.file-p2.file) == abs(p1.rank-p2.rank)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
