package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/thezeeshann/http-server/data"
)

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	html.Execute(w, data.GetAll()[1])
}

func main() {

	server := http.NewServeMux()
	server.HandleFunc("/template", HandleTemplate)

	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from go lang"))
	})

	// fs := http.FileServer(http.Dir("./index.tmpl"))
	// server.Handle("/", fs)

	err := http.ListenAndServe(":8001", server)
	if err != nil {
		fmt.Println("Error while opening server:", err)
	}

}
