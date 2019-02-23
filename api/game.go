package api

import (
	"errors"
	"github.com/ptflp/adrenalinBank/models"
	"github.com/andrewbackes/chess/game"
	"github.com/andrewbackes/chess/position/move"
)

var g *game.Game

func StartGame() {

	g = game.New()

}

func validateMove(p uint8, m models.ClientMoveObject) (string, error) {

	if p != uint8(g.ActiveColor()) {
		return "", errors.New("Сейчас не ваш ход")
	}

	s := move.Parse(m.Move)

	status, err := g.MakeMove(s)

	if err != nil {
		return "", errors.New("Невозможный ход")
	}

	var statusText string

	if p == 0 {
		statusText = "Белый "
	} else {
		statusText = "Черный "
	}

	switch status.String() {
	case "In progress":
		statusText += "сходил " + s.Source.String() + " " + s.Destination.String()
	case "Black checkmated Black":
		fallthrough
	case "White checkmated White":
		statusText += "получил мат"
	case "Threefold":
		fallthrough
	case "Fifty move rule":
		fallthrough
	case "Stalemate":
		fallthrough
	case "Insufficient material":
		statusText = "Ничья"
	}

	return statusText, nil
}
