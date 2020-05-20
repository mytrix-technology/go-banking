package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	helper "github.com/mytrix-technology/go-banking/helpers"
	"github.com/mytrix-technology/go-banking/vulnerableDB"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type Response struct {
	Data []vulnerableDB.User
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request)  {
	body, err := ioutil.ReadAll(r.Body)
	helper.HandleErr(err)

	var formatBody Login
	err = json.Unmarshal(body, &formatBody)
	helper.HandleErr(err)
	login := vulnerableDB.VulnerableLogin(formatBody.Username, formatBody.Password)

	if len(login) > 0 {
		resp := Response{Data: login}
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong Username and Password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartApi()  {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port : 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
