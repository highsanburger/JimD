// main.go
package main

import (
	"fmt"
	t "html/template"
	"log"
	h "net/http"
)

const (
	PORT = "6969"
)

func main() {

	h.HandleFunc("/", func(w h.ResponseWriter, r *h.Request) {
		t := t.Must(t.ParseFiles("./templates/index.html"))
		t.Execute(w, nil)
	})

	h.HandleFunc("/test", func(w h.ResponseWriter, r *h.Request) {
		t, _ := t.New("").Parse("<p> HII </p>")
		t.Execute(w, nil)
	})

	list := ""
	h.HandleFunc("/list", func(w h.ResponseWriter, r *h.Request) {
		list += "\n" + r.FormValue("listInput")
		t, _ := t.New("").Parse("<p> {{.}} </p>")
		t.Execute(w, list)
	})

	fmt.Println("Server is listening on :" + PORT)
	log.Fatal(h.ListenAndServe(":"+PORT, nil))
}
