package api

import (
	"errors"
	"github.com/ptflp/adrenalinBank/db"
	"github.com/ptflp/adrenalinBank/models"
	"regexp"
	"strings"
)

func Move(p uint8, m models.ClientMoveObject) (string, error) {

	m.Move = strings.Replace(strings.ToLower(m.Move), " ", "", -1)

	if m.Move == "" {
		return "", errors.New("Пустые данные")
	}

	m.Move = strings.Replace(m.Move, "а", "a", -1)
	m.Move = strings.Replace(m.Move, "y", "a", -1)
	m.Move = strings.Replace(m.Move, "б", "b", -1)
	m.Move = strings.Replace(m.Move, "п", "b", -1)
	m.Move = strings.Replace(m.Move, "с", "c", -1)
	m.Move = strings.Replace(m.Move, "з", "c", -1)
	m.Move = strings.Replace(m.Move, "ц", "c", -1)
	m.Move = strings.Replace(m.Move, "s", "c", -1)
	m.Move = strings.Replace(m.Move, "д", "d", -1)
	m.Move = strings.Replace(m.Move, "т", "d", -1)
	m.Move = strings.Replace(m.Move, "е", "e", -1)
	m.Move = strings.Replace(m.Move, "э", "e", -1)
	m.Move = strings.Replace(m.Move, "и", "e", -1)
	m.Move = strings.Replace(m.Move, "i", "e", -1)
	m.Move = strings.Replace(m.Move, "в", "f", -1)
	m.Move = strings.Replace(m.Move, "ф", "f", -1)
	m.Move = strings.Replace(m.Move, "v", "f", -1)
	m.Move = strings.Replace(m.Move, "j", "g", -1)

	if b, err := regexp.MatchString("[a-h][1-8][a-h][1-8]", m.Move); !b || err != nil {
		return "", errors.New("Некорректные данные")
	}

	status, err := validateMove(p, m)

	if err != nil {
		return "", err
	}

	wm := models.WebMoveObject{
		Player: p,
		Move:   m.Move,
		Status: status,
	}

	db.SaveMove("test", wm)

	return status, nil
}
