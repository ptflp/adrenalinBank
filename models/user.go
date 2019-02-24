package models

type Color string

const white Color = "w"
const black Color = "b"

type User struct {
	Token   string
	ProductId string
}