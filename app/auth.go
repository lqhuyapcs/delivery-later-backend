package app

import (
	"context"
	"fmt"
	"golang-api/models"
	u "golang-api/utils"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//JwtAuthentication - JWT Token
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{
			"/api/orderitems/new",
			"/api/orders/new",
			"/api/items/new",
			"/api/categories/new",
			"/api/stores/new", //Bỏ cái tên API vào đây để đánh dấu là nó không cần dùng token (không cần đăng nhập)
			"/api/accounts/new",
			"/api/stores/update",
			"/api/stores/delete",
			"/api/accounts/login",
			"/api/stores/search",
			"/api/stores/nearestStore",
			"/api/items/delete",
			"/api/categories/delete",
			"/api/reviews/new",
			"/api/stores/highestRateStore",
			"/api/accounts/update",
			"/api/stores/getAllStoreLocation",
			"/api/stores/deleteStoreLocation",
			"/api/stores/newestStore",
			"/api/reviews/search",
			"/api/orders/completedOrder",
			"/api/orders/incompletedOrder",
			"/api/items/update",
			"/api/orders/distance",
			"/api/orders/dateOrder",
			"/api/orders/completedDate",
		} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = u.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.UserID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
