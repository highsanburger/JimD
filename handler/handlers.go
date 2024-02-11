package handler

import (

	// t "JimD/timer"

	"html/template"
	"net/http"
	// "time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("./templates/index.html"))
	tmp.Execute(w, nil)
}
