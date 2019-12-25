package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CreateStore - controller
var CreateStore = func(w http.ResponseWriter, r *http.Request) {
	store := &m.Store{}                          //Gán biến store kế thừa model Store để map giữa controller và model
	err := json.NewDecoder(r.Body).Decode(store) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := store.Create() //Gọi hàm tạo mới trong model của Store
	u.Respond(w, resp)     //Trả về response
}

//QueryStoreByName - controller
var SearchStoreByName = func(w http.ResponseWriter, r *http.Request) {
	query := &m.Query{}
	err := json.NewDecoder(r.Body).Decode(query)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchStoreByName(query.Name)
	u.Respond(w, resp)
}

//SearchNearestStore - controller
var SearchNearestStore = func(w http.ResponseWriter, r *http.Request) {
	accountLocation := &m.AccountLocation{}
	err := json.NewDecoder(r.Body).Decode(accountLocation)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchNearestStore(accountLocation.Address, accountLocation.Lat, accountLocation.Lng)
	u.Respond(w, resp)
}

//UpdateStore - controller
var UpdateStore = func(w http.ResponseWriter, r *http.Request) {
	store := &m.Store{}
	err := json.NewDecoder(r.Body).Decode(store)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := store.UpdateStore()
	u.Respond(w, resp)
}

//DeleteStore - controller
var DeleteStore = func(w http.ResponseWriter, r *http.Request) {
	store := &m.Store{}
	err := json.NewDecoder(r.Body).Decode(store)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := store.DeleteStore()
	u.Respond(w, resp)
}
