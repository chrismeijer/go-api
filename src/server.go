package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	models "../models"
	userType "../types"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type User = userType.User
type N1qlUser = userType.N1qlUser

// MAIN FUNCTION
func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	// CONNECT COUCHBASE
	models.InitDB(os.Getenv("COUCHBASE_HOST"), os.Getenv("COUCHBASE_USERNAME"), os.Getenv("COUCHBASE_PASSWORD"))

	// SETUP ROUTER
	router := mux.NewRouter()

	// ROUTES
	// ROOT
	router.HandleFunc("/", HomeHandler).Methods("GET")

	// USERS
	router.HandleFunc("/users", models.AllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", models.GetUser).Methods("GET")
	router.HandleFunc("/users", models.AddUser).Methods("POST")

	// SETUP SERVER
	log.Fatal(http.ListenAndServe(":8081", router))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
