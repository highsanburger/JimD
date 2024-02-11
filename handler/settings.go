package handler

import (
	g "JimD/global"
	// t "JimD/timer"

	"html/template"
	"net/http"
	// "time"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("./templates/settings.html"))
	tmp.Execute(w, g.Phile)
}

func FileLocn(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("file_locn")
	g.Phile = f
	tmp := template.Must(template.ParseFiles("./templates/dot.html"))
	tmp.Execute(w, g.Phile)
}
