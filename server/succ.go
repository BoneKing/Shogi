package server

import (
	"encoding/json"
	"fmt"
	"strings"
)

type PossibleMoves struct {
	piece string
	NewXs []int //Xs and Ys can be paired into coordinates by indexing them and the same spot
	NewYs []int //so NewXs[1] and NewYs[1] will give the second possible coordinate
}

func succ(pm PossibleMoves, piece string, CurrX int, CurrY int) []uint8 {
	var NewX int
	var NewY int
	switch piece {
	case "P1", "P2":
		if strings.Contains(piece, "1") {
			NewX = CurrX
			NewY = CurrY + 1
		}
		if strings.Contains(piece, "2") {
			NewX = CurrX
			NewY = CurrY - 1
		}
		pm.NewXs = append(pm.NewXs, NewX)
		pm.NewYs = append(pm.NewYs, NewY)
	case "L1":
		for j := CurrY; j < 9; j++ {
			if board[CurrX][j] != "O" {
				if strings.Contains(board[CurrX][j], "1") {
					break
				} else if strins.Contains(board[CurrX][j], "2") {
					NewY = j
					NewX = CurrX
					pm.NewXs = append(pm.NewXs, NewX)
					pm.NewYs = append(pm.NewYs, NewY)
					break
				}
				NewY = j
				NewX = CurrX
				pm.NewXs = append(pm.NewXs, NewX)
				pm.NewYs = append(pm.NewYs, NewY)
			}

		}
	case "L2":
		for j := CurrY; j >= 0; j-- {
			if board[CurrX][j] != "O" {
				if strings.Contains(board[CurrX][j], "2") {
					break
				} else if strings.Contains(board[CurrX][j], "1") {
					NewY = j
					NewX = CurrX
					pm.NewXs = append(pm.NewXs, NewX)
					pm.NewYs = append(pm.NewYs, NewY)
					break
				}
				NewY = j
				NewX = CurrX
				pm.NewXs = append(pm.NewXs, NewX)
				pm.NewYs = append(pm.NewYs, NewY)
			}

		}
	case "N1", "N2":
		if strings.Contains(piece, "1") {
			if CurrY+2 < 9 {
				NewY = CurrY + 2
			}
		} else if strings.Contains(piece, "2") {
			if CurrY-2 >= 0 {
				NewY = CurrY + 2
			}
		}
		NewX = CurrX + 1
		pm.NewXs = append(pm.NewXs, NewX)
		pm.NewYs = append(pm.NewYs, NewY)
		NewX = CurrX - 1
		pm.NewXs = append(pm.NewXs, NewX)
		pm.NewYs = append(pm.NewYs, NewY)
	case "S1", "S2":
		var moves []int
		if strings.Contains(piece, "1") {
			moves = []int{-1, -1, 0, -1, 1, -1, -1, 1, 1, 1}
		}
		if strings.Contains(piece, "2") {
			moves = []int{-1, 1, 0, 1, 1, 1, -1, -1, 1, -1}
		}
		for j := 0; j <= len(moves)-2; j += 2 {
			NewX := CurrX + moves[j]
			NewY := CurrY + moves[j+1]
			pm.NewXs = append(pm.NewXs, NewX)
			pm.NewYs = append(pm.NewYs, NewY)
		}
	}
	for i := 0; i < len(pm.NewXs); i++ {
		if !IsValid(board, pm.NewXs[i], pm.NewYs[i], playerNum) {
			pm.NewXs = append(pm.NewXs[:i], pm.NewXs[i+1:]...)
			pm.NewYs = append(pm.NewYs[:i], pm.NewYs[i+1:]...)
		}
	}
	result, err := json.Marshal(pm)
	if err != nil {
		fmt.Printf("ERROR: ", err)
	}
	return result
}
