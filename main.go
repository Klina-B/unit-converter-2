package main

import (
	"fmt"
	"net/http"

	"github.com/Klina-B/unit-converter/handlers"
)


func main(){
	fmt.Println("Started")
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/length", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeightHandler)
	http.ListenAndServe(":8080", nil)
}