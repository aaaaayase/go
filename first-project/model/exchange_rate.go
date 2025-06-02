package model

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	FromCurrency string    `json:"from_currency"`
	ToCurrency   string    `json:"to_currency"`
	Rate         string    `json:"rate"`
	Date         time.Time `json:"date"`
}
