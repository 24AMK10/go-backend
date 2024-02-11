package main

import (
	"fmt"
	"net/http"
	"backend.com/example/go-backend/api"
)

func main(){
	api.StartMongo()
	fmt.Println("Hello There !!")

	http.HandleFunc("/signup", api.HandleSignUp)

	http.ListenAndServe(":8081", nil)
}