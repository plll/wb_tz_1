package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const getOrderIdTemplateHTML = `<html>
  <body>
    <form action="/orders" method="post">
      <div> <label> orderId </label> <input name="orderId" type="text" /></div>
      <div><input type="submit" value="view"></div>
    </form>
  </body>
</html>
    `

func (s *Server) ordersHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		orderId := req.FormValue("orderId")
		value, ok := s.cache.Get(orderId)
		if ok {
			j, _ := json.Marshal(&value)
			res.Write(j)
			return
		}
		order, err := s.repos.OrderRepository.CollectOrderById(s.ctx,
			s.repos.Items,
			s.repos.Payments,
			s.repos.Deliveries,
			orderId)
		if err != nil {
			fmt.Print(err)
			res.WriteHeader(http.StatusBadRequest)
		}
		j, _ := json.Marshal(&order)
		res.Write(j)
	} else if req.Method == "GET" {
		res.Write([]byte(getOrderIdTemplateHTML))
	} else {
		res.WriteHeader(http.StatusForbidden)
	}
}
