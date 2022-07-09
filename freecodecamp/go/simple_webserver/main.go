package main

import (
	"fmt"
	"net/http"
)

func main() {
	file_server := http.FileServer(http.Dir("./serve"))
	http.Handle("/", file_server)
	http.HandleFunc("/home", index_handler)
	http.HandleFunc("/form", form_handler)
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8080", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
func form_handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	// fmt.Println(name, email)
	fmt.Fprintf(w, "name: %s \n email: %s", name, email)
	// if err := r
}
