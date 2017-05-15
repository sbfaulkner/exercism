package queenattack

import "errors"

const testVersion = 2

// CanQueenAttack determines whether or not two queens are positioned so that they can attack each other
func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errors.New("Both queens cannot occupy the same position")
	}

	whiteFile, whiteRank, err := parsePosition(whitePosition)

	if err != nil {
		return false, err
	}

	blackFile, blackRank, err := parsePosition(blackPosition)

	if err != nil {
		return false, err
	}

	if blackFile == whiteFile || blackRank == whiteRank {
		return true, nil
	}

	if abs(blackFile-whiteFile) == abs(blackRank-whiteRank) {
		return true, nil
	}

	return false, nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func parsePosition(position string) (file int, rank int, err error) {
	if file = int(position[0] - 'a'); file < 0 || file > 7 {
		err = errors.New("Invalid position %q - rank should be from a to h")
	}

	if rank = int(position[1] - '1'); rank < 0 || rank > 7 {
		err = errors.New("Invalid position %q - file should be from 1 to 8")
	}

	return
}
