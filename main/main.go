package main

import (
	"net/http"
	hand "youmed/handlers"
)

func main() {
	http.HandleFunc("/", hand.HandlerHome)
	http.ListenAndServe(":8080", nil)

}
