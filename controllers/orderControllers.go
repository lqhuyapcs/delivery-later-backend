package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)


//CreateOrder - controller
var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	order := &m.Order{}                       
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := order.CreateOrder() //Gọi hàm tạo mới trong model của Store
	u.Respond(w, resp)                //Trả về response
}

//UpdateOrder - controller
var UpdateOrder = func(w http.ResponseWriter, r *http.Request) {
	order := &m.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := order.UpdateOrder()
	u.Respond(w, resp)
}

//DeleteOrder - controller
var DeleteOrder = func(w http.ResponseWriter, r *http.Request) {
	order := &m.Order{}
	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := order.DeleteOrder()
	u.Respond(w, resp)
}