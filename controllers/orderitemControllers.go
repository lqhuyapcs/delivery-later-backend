package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)


//CreateOrderItem - controller
var CreateOrderItem = func(w http.ResponseWriter, r *http.Request) {
	orderItem := &m.OrderItem{}                       
	err := json.NewDecoder(r.Body).Decode(orderItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := orderItem.CreateOrderItem() //Gọi hàm tạo mới trong model của Store
	u.Respond(w, resp)                //Trả về response
}

//UpdateOrderItem - controller
var UpdateOrderItem = func(w http.ResponseWriter, r *http.Request) {
	orderItem := &m.OrderItem{}
	err := json.NewDecoder(r.Body).Decode(orderItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := orderItem.UpdateOrderItem()
	u.Respond(w, resp)
}

//DeleteOrderItem - controller
var DeleteOrderItem = func(w http.ResponseWriter, r *http.Request) {
	orderItem := &m.OrderItem{}
	err := json.NewDecoder(r.Body).Decode(orderItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := orderItem.DeleteOrderItem()
	u.Respond(w, resp)
}