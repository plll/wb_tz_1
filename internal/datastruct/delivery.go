package datastruct

import (
	"errors"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

func (d Delivery) Validate() error {
	if d.Name == "" {
		return errors.New("No Name")
	}
	if d.Phone == "" {
		return errors.New("No Phone")
	}
	if d.Zip == "" {
		return errors.New("No Zip")
	}
	if d.City == "" {
		return errors.New("No City")
	}
	if d.Address == "" {
		return errors.New("No Address")
	}
	if d.Region == "" {
		return errors.New("No Region")
	}
	if d.Email == "" {
		return errors.New("No Email")
	}
	return nil
}
