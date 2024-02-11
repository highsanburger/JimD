package main

import (
	g "JimD/global"
	h "JimD/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", h.Index)

	http.HandleFunc("/addwo", h.AddWorkout)
	http.HandleFunc("/addwo/addex", h.AddExercise)

	http.HandleFunc("/settings", h.Settings)
	http.HandleFunc("/settings/locn", h.FileLocn)

	fmt.Println("Server is listening on :" + g.PORT)

	log.Fatal(http.ListenAndServe(":"+g.PORT, nil))
}
