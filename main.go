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
	router.HandleFunc("/api/stores/delete", controllers.DeleteStore).Methods("GET")
	router.HandleFunc("/api/categories/new", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/api/categories/update", controllers.UpdateCategory).Methods("POST")
	router.HandleFunc("/api/categories/delete", controllers.DeleteCategory).Methods("GET")
	router.HandleFunc("/api/items/new", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/api/items/update", controllers.UpdateItem).Methods("POST")

	router.HandleFunc("/api/items/delete", controllers.DeleteItem).Methods("GET")
	router.HandleFunc("/api/orders/new", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/update", controllers.UpdateOrder).Methods("POST")
	router.HandleFunc("/api/orders/delete", controllers.DeleteOrder).Methods("GET")
	router.HandleFunc("/api/orderitems/new", controllers.CreateOrderItem).Methods("POST")
	router.HandleFunc("/api/orderitems/update", controllers.UpdateOrderItem).Methods("POST")
	router.HandleFunc("/api/orderitems/delete", controllers.DeleteOrderItem).Methods("GET")

	// Search
	router.HandleFunc("/api/stores/search", controllers.SearchStoreByName).Methods("POST")

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
