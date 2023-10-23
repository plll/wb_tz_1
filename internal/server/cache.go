package server

import "log"

func (s *Server) prepareCache() {
	tmp, err := s.repos.Orders.GetNLastOrders(s.ctx, s.repos.Items, s.repos.Payments, s.repos.Deliveries, 10)
	if err != nil {
		log.Fatal(err)
	}
	for _, x := range tmp {
		s.cache.Add(x.OrderUid, x)
	}
}
