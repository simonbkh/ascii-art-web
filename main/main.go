package main

import (
	"net/http"
	hand "youmed/handlers"
)

func main() {
	http.HandleFunc("/", hand.HandlerHome)
	http.HandleFunc("/ascii-art", hand.Handlerascii)
	http.ListenAndServe(":8080", nil)
}
