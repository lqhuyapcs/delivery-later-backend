package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CreateReview - controller
var CreateReview = func(w http.ResponseWriter, r *http.Request) {
	review := &m.Review{}
	err := json.NewDecoder(r.Body).Decode(review) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := review.CreateReview() //Gọi hàm tạo mới trong model của Review
	u.Respond(w, resp)            //Trả về response
}

//Search Review by Store ID - controller
var SearchReview = func(w http.ResponseWriter, r *http.Request) {
	store := &m.Store{}
	err := json.NewDecoder(r.Body).Decode(store)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := m.SearchReview(store.ID)
	u.Respond(w, resp)
}
