package routes

import (
	controllers "github.com/gamorvi/restapi/controllers"
	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {
	s := r.PathPrefix(prefix).Subrouter()
	//s.HandleFunc("/", helloWorld).Methods("GET")
	s.HandleFunc("/users", controllers.AllUsers).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.OneUsers).Methods("GET")
	s.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
}
