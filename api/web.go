package api

import (
	"encoding/json"
	"github.com/ptflp/adrenalinBank/db"
)

func GetMovesResponse(game string) (string, error) {
	moves, err := db.GetMoves(game)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(moves)
	if err != nil {
		return "", err
	} else {
		return string(b), err
	}
}
