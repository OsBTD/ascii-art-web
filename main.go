package main

import (
	"ascii-art-web/ascii"
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	temp, err := template.ParseFiles("templates\\index.html")
	if err != nil {
		http.Error(w, "Error : page not found", http.StatusNotFound)
	}
	temp.Execute(w, nil)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates\\index.html")
	if err != nil {
		http.Error(w, "Error : page not found", http.StatusNotFound)
	}
	input := r.FormValue("text")
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return

	}

	PrintArt := ascii.PrintArt(input, banner)

	temp.Execute(w, PrintArt)

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/ascii", Ascii)
	fmt.Println("local host running : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//skipping unprintable characters except \r and \n
