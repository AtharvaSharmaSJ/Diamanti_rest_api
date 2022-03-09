package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
    Name string `json:"Name"`
    Age string `json:"Age"`
}
var Users []User

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "You've hit home")
}

func users(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(Users)
}

func handleRequests() {
    http.HandleFunc("/", home)
	http.HandleFunc("/users",users)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	Users = []User{
		User{Name:"Atharva", Age:"26"},
		User{Name:"Sharma", Age:"26"},
	}
    handleRequests()
}