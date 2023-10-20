package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) ordersHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		orderId := req.FormValue("orderId")
		order, err := s.repos.OrderRepository.CollectOrderById(orderId)
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
