package main

import (
	"fmt"
	"golang-api/app"
	"golang-api/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/accounts/new", controllers.CustomerRegister).Methods("POST")
	router.HandleFunc("/api/accounts/login", controllers.CustomerAuthenticate).Methods("POST")
	router.HandleFunc("/api/stores/new", controllers.CreateStore).Methods("POST") //Thêm API goị đến controller CreateStore
	router.HandleFunc("/api/stores/update", controllers.UpdateStore).Methods("POST")
	router.HandleFunc("/api/stores/delete", controllers.DeleteStore).Methods("POST")
	router.HandleFunc("/api/categories/new", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/categories/update", controllers.UpdateCategory).Methods("POST")
	router.HandleFunc("/api/categories/delete", controllers.DeleteCategory).Methods("POST")
	router.HandleFunc("/api/items/new", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/api/items/new", controllers.UpdateItem).Methods("POST")
	router.HandleFunc("/api/items/new", controllers.DeleteItem).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	// router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
