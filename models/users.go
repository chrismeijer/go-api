package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/couchbase/gocb.v1"

	userType "../types"
)

type User = userType.User

type N1qlUser = userType.N1qlUser

// GET ALL USERS
func AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	query := gocb.NewN1qlQuery("SELECT * FROM `tires` AS `user`")
	rows, err := globalBucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		fmt.Fprintf(w, "Error")
		//return nil
	}

	var row N1qlUser
	for rows.Next(&row) {
		users = append(users, row.User)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GET N1 USER
func GetUser(w http.ResponseWriter, r *http.Request) {
	var N1qlParams []interface{}
	query := gocb.NewN1qlQuery("SELECT * FROM `tires` AS `user` WHERE META(`user`).id = $1")

	params := mux.Vars(r)

	N1qlParams = append(N1qlParams, params["id"])
	rows, err := globalBucket.ExecuteN1qlQuery(query, N1qlParams)
	if err != nil {
		fmt.Fprintf(w, "Error")
		//return nil
	}

	var row N1qlUser
	rows.One(&row)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(row.User)
}

// ADD USER
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var n1qlParams []interface{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	query := gocb.NewN1qlQuery("INSERT INTO `tires` (KEY, VALUE) VALUES ($1, {'firstname': $2, 'lastname': $3, 'email': $4})")
	u, _ := uuid.NewV4()
	n1qlParams = append(n1qlParams, u.String())
	n1qlParams = append(n1qlParams, user.Firstname)
	n1qlParams = append(n1qlParams, user.Lastname)
	n1qlParams = append(n1qlParams, user.Email)
	_, err := globalBucket.ExecuteN1qlQuery(query, n1qlParams)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(user)
}
