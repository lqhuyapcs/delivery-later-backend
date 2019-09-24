package controllers

import (
	"encoding/json"
	m "golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//CustomerRegister - controller
var CustomerRegister = func(w http.ResponseWriter, r *http.Request) {

	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

//CustomerAuthenticate - controller
var CustomerAuthenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &m.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := m.Authenticate(account.Phone, account.Password)
	u.Respond(w, resp)
}
