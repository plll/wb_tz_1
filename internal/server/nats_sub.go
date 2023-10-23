package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/datastruct"
	"log"
	"time"
)

func (s *Server) createSubscriber() {
	s.sc.Subscribe("orders",
		func(m *stan.Msg) {
			var o datastruct.Order
			err := json.Unmarshal(m.Data, &o)
			if err != nil {
				log.Println("Unmarshal error")
				return
			}
			err = o.Validate()
			if err != nil {
				log.Println("Error in json fields")
				return
			}
			err = o.Payment.Validate()
			if err != nil {
				log.Println("Error in json fields")
				return
			}
			err = o.Delivery.Validate()
			if err != nil {
				log.Println("Error in json fields")
				return
			}
			for _, i := range o.Items {
				err = i.Validate()
				if err != nil {
					log.Println("Error in json fields")
					return
				}
			}
			if err == nil {
				err = s.repos.Orders.AddNewOrder(context.Background(), s.repos.Items,
					s.repos.Payments, s.repos.Deliveries, s.repos.OrdersItems, o)
				if err != nil {
					fmt.Print(err)
				}
				s.cache.Add(o.OrderUid, o)
			} else {
				fmt.Print(err)
			}
		},
		stan.AckWait(20*time.Second))
}
