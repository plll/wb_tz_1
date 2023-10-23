package server

func (s *Server) prepareCache() {
	s.cache = s.repos.Orders.GetNLastOrders(s.ctx, s.repos.Items, s.repos.Payments, s.repos.Deliveries, 10)
}
