package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CreateOrder - controller
var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	order := []m.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		println("1")
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	for i := range order {
		resp := m.CreateOrder(order[i])
		u.Respond(w, resp) //Trả về response
	}
	//Gọi hàm tạo mới trong model của Store
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

//SearchCompletedOrder - controller
var SearchCompletedOrder = func(w http.ResponseWriter, r *http.Request) {
	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchCompletedOrder(account.ID)
	u.Respond(w, resp)
}

//SearchIncompletedOrder - controller
var SearchIncompletedOrder = func(w http.ResponseWriter, r *http.Request) {
	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchIncompletedOrder(account.ID)
	u.Respond(w, resp)
}
