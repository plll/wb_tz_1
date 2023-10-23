package datastruct

import (
	"errors"
)

type Item struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (i Item) Validate() error {
	if i.TrackNumber == "" {
		return errors.New("No TrackNumber")
	}
	if i.Rid == "" {
		return errors.New("No Rid")
	}
	if i.Name == "" {
		return errors.New("No Name")
	}
	if i.Size == "" {
		return errors.New("No Size")
	}
	if i.Brand == "" {
		return errors.New("No Brand")
	}
	return nil
}
