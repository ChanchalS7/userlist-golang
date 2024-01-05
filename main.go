package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

type User struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			Name:   "chanchal",
			Age:    25,
			Gender: "Male",
		},
		{
			Name:   "Harsh",
			Age:    25,
			Gender: "Male",
		},
	}
	usersjson, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(usersjson)
}
func main() {
	http.HandleFunc("/users", UsersHandler)
	fmt.Printf("Server is listening on port number %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
