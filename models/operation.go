package models

type Operation struct {
	Id   string
	Date string
	Timestamp int64
	Amount float64
	Currency string
	Details string
	CounterpartName string
	CounterpartBankName string
}