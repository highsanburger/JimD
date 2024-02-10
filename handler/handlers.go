package handler

import (
	g "JimD/global"
	md "JimD/markdown"
	d "JimD/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("./templates/index.html"))
	tmp.Execute(w, nil)
}

func AddWorkout(w http.ResponseWriter, r *http.Request) {
	exss, _ := md.ReadLinesFromFile(g.Phile)
	var f []*d.Exercise
	for _, exs := range exss {
		ex, _ := md.ParseExerciseString(exs)
		f = append(f, ex)
	}
	tmp := template.Must(template.ParseFiles("./templates/addwo.html"))
	fmt.Println(tmp.Execute(w, f))
}

func AddExercise(w http.ResponseWriter, r *http.Request) {
	hasDate, _ := md.HasDateInFirstLine(g.Phile)
	if !hasDate {
		md.EnterDate(g.Phile)
	}
	name := r.FormValue("exName")
	reps, _ := strconv.Atoi(r.FormValue("exRep"))
	sets, _ := strconv.Atoi(r.FormValue("exSet"))
	weight, _ := strconv.ParseFloat(r.FormValue("exWeight"), 32)
	ex := d.Exercise{Name: name, Reps: reps, Sets: sets, Weight: weight}
	md.EnterEx(g.Phile, ex)
	tmp := template.Must(template.ParseFiles("./templates/addex.html"))
	tmp.Execute(w, ex)
}
