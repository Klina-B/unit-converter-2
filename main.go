package main

import (
	"fmt"
	"net/http"

	"github.com/Klina-B/unit-converter-2/handlers"
)


func main(){
	fmt.Println("Started at http://localhost:8080")
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/length", handlers.LengthHandler)
	http.HandleFunc("/weight", handlers.WeightHandler)
	http.ListenAndServe(":8080", nil)
}