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

	}
	result, err := json.Marshal(pm)
	if err != nil {
		fmt.Printf("ERROR: ", err)
	}
	return result
}
