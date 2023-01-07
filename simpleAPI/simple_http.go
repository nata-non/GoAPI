package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
	Age       int    `json:"age"`
}

var users = []User{
	{ID: 1, FirstName: "Natanon", LastName: "Hanchana", Age: 20},
}

func handle(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func main() {
	http.HandleFunc("/users", handle)
	log.Println("Server started at :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("Server close, Bye Bye")
}
