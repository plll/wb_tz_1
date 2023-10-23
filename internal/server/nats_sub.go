package server

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/plll/wb_tz_1/internal/datastruct"
	"time"
)

func (s *Server) createSubscriber() {
	s.sc.Subscribe("orders",
		func(m *stan.Msg) {
			var o datastruct.Order
			err := json.Unmarshal(m.Data, &o)
			if err == nil {
				err = s.repos.OrdersRepository.AddNewOrder(s.ctx, &o)
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
