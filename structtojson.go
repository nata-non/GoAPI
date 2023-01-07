package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func main() {
	u := User{
		ID: 1, FirstName: "Natanon", LastName: "Hanchana", Age: 20}
	b, err := json.Marshal(u)
	fmt.Printf("byte : %s", b)
	fmt.Println((err))
}
