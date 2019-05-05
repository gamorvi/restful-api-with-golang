package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	controllers "github.com/gamorvi/restapi/controllers"
	_ "github.com/gamorvi/restapi/models"
	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func startServer() {
	port := os.Getenv("web_port")
	prefix := os.Getenv("prefix")
	fmt.Println("Server started at " + port + "...")
	r := mux.NewRouter().StrictSlash(true)
	// Routes
	s := r.PathPrefix(prefix).Subrouter()
	r.HandleFunc("/", helloWorld).Methods("GET")
	s.HandleFunc("/users", controllers.AllUsers).Methods("GET")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.OneUsers).Methods("GET")
	s.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	//Start Server on the port set in your .env file
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	startServer()
}
