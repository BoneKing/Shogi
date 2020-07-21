package server

import (
	"strconv"
	"strings"
)

func OwnsPiece(piece string, playerNum int) bool {
	strPlayerNum := strconv.Itoa(playerNum)
	if strings.Contains(piece, strPlayerNum) { //if the piece and the player number have the same number than the player owns that piece
		return true
	}
	return false
}

func IsValid(board [][]string, NewX int, NewY int, player int) bool { //makes sure that the piece isnt moving out of bounds or on top a piece the player already owns
	if NewX >= len(board[0]) || NewX < 0 {
		return false
	}
	if NewY >= len(board) || NewY < 0 {
		return false
	}
	StrPlayer := strconv.Itoa(player)
	if strings.Contains(board[NewY][NewX], StrPlayer) {
		return false
	}
	return true
}

func CheckPromotion(Newy int, piece string) bool { //checks if the piece can be promoted.
	if Newy > 5 {
		switch piece {
		case "P1":
			return true
		case "L1":
			return true
		case "N1":
			return true
		case "S1":
			return true
		case "B1":
			return true
		case "R1":
			return true
		}
	}
	if Newy < 3 {
		switch piece {
		case "P2":
			return true
		case "L2":
			return true
		case "N2":
			return true
		case "S2":
			return true
		case "B2":
			return true
		case "R2":
			return true
		}
	}
	return false
}
