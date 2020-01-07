package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CreateItem - controller
var CreateItem = func(w http.ResponseWriter, r *http.Request) {
	item := &m.Item{}
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := item.CreateItem() //Gọi hàm tạo mới trong model của Store
	u.Respond(w, resp)        //Trả về response
}

//UpdateItem - controller
var UpdateItem = func(w http.ResponseWriter, r *http.Request) {
	item := &m.Item{}
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.UpdateItem()
	u.Respond(w, resp)
}

//DeleteItem - controller
var DeleteItem = func(w http.ResponseWriter, r *http.Request) {
	item := &m.Item{}
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := item.DeleteItem()
	u.Respond(w, resp)
}
