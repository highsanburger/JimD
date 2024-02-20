package main

import (
	g "JimD/global" // assuming the correct import path for your global package
	"fmt"
	"github.com/a-h/templ"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("index.html"))
		tmp.Execute(w, nil)
	})

	c := hello("GET GOT")
	mux.Handle("/t", templ.Handler(c))

	fmt.Println("Server is listening on :" + g.PORT)

	log.Fatal(http.ListenAndServe(":"+g.PORT, mux))
}
