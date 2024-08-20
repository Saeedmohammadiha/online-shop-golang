package middlewares

import (
	"net/http"
	"strings"
)

func AuthProtect(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//check the header if there is a token
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader =="" {

		}
		authorizationHeaderArray := strings.Split(authorizationHeader, " ")
		token := authorizationHeaderArray[1]


	}
	//if there is check it in the user data that is stored in the context and check it if it is expired
	// send refresh token if it is necesary
	//if the token is valid pass the function
	//if not respose with 401 status code
	//return f(w, r)
}
