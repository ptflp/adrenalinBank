package models

type ClientMoveObject struct {
	Move string `json:"text"`
}

type WebMoveObject struct {
	Player uint8  `json:"user"`
	Move   string `json:"move"`
	Status string `json:"status"`
}

type MoveObjects []WebMoveObject
