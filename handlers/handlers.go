package hand

import (
	"fmt"
	"html/template"
	"net/http"

	"youmed/ascii-art"
	check "youmed/utils"
)

type mok struct {
	Input  string
	Banner string
}

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("../templates/home.html")
	if err != nil {
		http.Error(w, "enternal server error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func Handlerascii(w http.ResponseWriter, r *http.Request) {
	v := ""

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("../templates/ascii-art.html")
	if err != nil {
		http.Error(w, "enternal server error", http.StatusInternalServerError)
		return
	}
	data := mok{
		Input:  r.FormValue("input"),
		Banner: r.FormValue("banner"),
	}
	 fmt.Println(data.Banner)
	if data.Input == "" || (data.Banner != "standard" && data.Banner != "shadow" && data.Banner != "thinkertoy") || !check.CheckIn(data.Input) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	v = ascii.Ascii(data.Banner, data.Input)
	tmpl.Execute(w, v)
}
