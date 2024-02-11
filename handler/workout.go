package handler

import (
	g "JimD/global"
	md "JimD/markdown"
	d "JimD/models"

	// t "JimD/timer"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "time"
)

func AddWorkout(w http.ResponseWriter, r *http.Request) {

	fmt.Println("two")
	exss, _ := md.ReadLinesFromFile(g.Phile)
	var f []*d.Exercise
	for _, exs := range exss {
		if md.IsExerciseFormat(exs) {
			ex, _ := md.ParseExerciseString(exs)
			f = append(f, ex)
		}
	}
	tmp := template.Must(template.ParseFiles("./templates/addwo.html"))
	fmt.Println(tmp.Execute(w, f))
}

func AddExercise(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("exName")
	reps, _ := strconv.Atoi(r.FormValue("exRep"))
	sets, _ := strconv.Atoi(r.FormValue("exSet"))
	weight, _ := strconv.ParseFloat(r.FormValue("exWeight"), 32)
	ex := d.Exercise{Name: name, Reps: reps, Sets: sets, Weight: weight}

	hasDate, _ := md.HasDateInFirstLine(g.Phile)
	if !hasDate {
		md.EnterDate(g.Phile)
	}
	md.EnterEx(g.Phile, ex)

	tmp := template.Must(template.ParseFiles("./templates/addex.html"))
	tmp.Execute(w, ex)
}
