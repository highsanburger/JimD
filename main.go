// main.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	PORT = "6969"
)

type Exercise struct {
	Name string
	Reps int
	Sets int
	// reps int
	// sets int
	// tags []string
}

func enterDate(path string) {
	time := time.Now().Format("02/01/2006 15:04")
	err := os.WriteFile(path, []byte(time), 0644)
	if err != nil {
		panic(err)
	}
}

func enterEx(path string, ex Exercise) {
	exs := ex.Name + string(ex.Sets) + string(ex.Reps)
	err := os.WriteFile(path, []byte(exs), 0644)
	if err != nil {
		panic(err)
	}

}

func main() {
	http.Handle("/static/",

		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("./templates/index.html"))
		tmp.Execute(w, nil)
	})

	http.HandleFunc("/addwo", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("./templates/addwo.html"))
		tmp.Execute(w, nil)
	})

	http.HandleFunc("/addwo/addex", func(w http.ResponseWriter, r *http.Request) {
		enterDate("test.md")
		name := r.FormValue("exName")
		reps, _ := strconv.Atoi(r.FormValue("exRep"))
		sets, _ := strconv.Atoi(r.FormValue("exSet"))
		ex := Exercise{name, reps, sets}
		enterEx("test.md", ex)
		tmp := template.Must(template.ParseFiles("./templates/addex.html"))
		tmp.Execute(w, ex)
	})

	fmt.Println("Server is listening on :" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
