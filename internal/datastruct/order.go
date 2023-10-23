package datastruct

import (
	"errors"
	"time"
)

type Order struct {
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
	Delivery          Delivery
	Payment           Payment
	Items             []Item
}

func (o Order) Validate() error {
	if o.OrderUid == "" {
		return errors.New("No order_uid")
	}
	if o.CustomerId == "" {
		return errors.New("No customer_id")
	}
	if o.Entry == "" {
		return errors.New("No entry")
	}
	if o.InternalSignature == "" {
		return errors.New("No internal_signature")
	}
	if o.OofShard == "" {
		return errors.New("No oof_shard")
	}
	if o.Locale == "" {
		return errors.New("No locale")
	}
	if o.DeliveryService == "" {
		return errors.New("No delivery_service")
	}
	if o.TrackNumber == "" {
		return errors.New("No track_number")
	}
	if o.Shardkey == "" {
		return errors.New("No shardkey")
	}
	return nil
}
