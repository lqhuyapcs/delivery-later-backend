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
