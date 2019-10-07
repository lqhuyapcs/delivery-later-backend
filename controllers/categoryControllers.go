package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CreateStore - controller
var CreateCategory = func(w http.ResponseWriter, r *http.Request) {
	category := &m.Category{}                       //Gán biến store kế thừa model Store để map giữa controller và model
	err := json.NewDecoder(r.Body).Decode(category) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := category.CreateCategory() //Gọi hàm tạo mới trong model của Store
	u.Respond(w, resp)                //Trả về response
}

//UpdateStore - controller
var UpdateCategory = func(w http.ResponseWriter, r *http.Request) {
	category := &m.Category{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := category.UpdateCategory()
	u.Respond(w, resp)
}

//DeleteStore - controller
var DeleteCategory = func(w http.ResponseWriter, r *http.Request) {
	category := &m.Category{}
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := category.DeleteCategory()
	u.Respond(w, resp)
}
