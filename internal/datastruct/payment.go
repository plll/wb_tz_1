package datastruct

import (
	"errors"
)

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

func (p Payment) Validate() error {
	if p.Transaction == "" {
		return errors.New("No Transaction")
	}
	if p.RequestId == "" {
		return errors.New("No RequestId")
	}
	if p.Currency == "" {
		return errors.New("No Currency")
	}
	if p.Provider == "" {
		return errors.New("No Provider")
	}
	if p.Bank == "" {
		return errors.New("No Bank")
	}
	return nil
}
