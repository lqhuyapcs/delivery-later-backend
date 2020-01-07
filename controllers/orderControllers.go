package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

type RequestDate struct {
	ID uint
}

//CreateOrder - controller
var CreateOrder = func(w http.ResponseWriter, r *http.Request) {
	order := []m.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	var temp []map[string]interface{}
	for i := range order {
		resp := m.CreateOrder(order[i])
		temp = append(temp, resp)
	}
	u.RespondOrder(w, temp)
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

//SearchOrderByDate - controller
var SearchOrderByDate = func(w http.ResponseWriter, r *http.Request) {
	dateorder := &m.DateOrder{}
	err := json.NewDecoder(r.Body).Decode(dateorder)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchOrderByDate(dateorder.ID, dateorder.Date)
	u.Respond(w, resp)
}

//SearchDate - controller
var SearchDate = func(w http.ResponseWriter, r *http.Request) {
	request := &RequestDate{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchDate(request.ID)
	u.Respond(w, resp)
}

//GetDistanceAfterUpdateAddress - controller
var GetDistanceAfterUpdateAddress = func(w http.ResponseWriter, r *http.Request) {
	getdistance := &m.GetDistance{}
	err := json.NewDecoder(r.Body).Decode(getdistance)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	var resp map[string]interface{}
	resp = u.Message(true, "Get distance success")
	Lat1 := getdistance.Lat1
	Lng1 := getdistance.Lng1
	Lat2 := getdistance.Lat2
	Lng2 := getdistance.Lng2
	if resp["distance"] == nil {
		resp["distance"] = u.Distance(Lat1, Lng1, Lat2, Lng2)
	}
	u.Respond(w, resp)
}
