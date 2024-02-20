package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	// "github.com/a-h/templ"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {

	PORT := ":6969"
	mux := http.NewServeMux()
	s := &http.Server{
		Addr:           PORT,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("index.html"))
		tmp.Execute(w, nil)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("login.html"))
		tmp.Execute(w, nil)
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles("signup.html"))
		tmp.Execute(w, nil)
	})

	fmt.Println("Server is listening on " + PORT)
	log.Fatal(s.ListenAndServe())

	connStr := "postgresql://highsanburger:Jb6IpVlqig3z@ep-rapid-waterfall-a1qqcafq.ap-southeast-1.aws.neon.tech/jimd?sslmode=require"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	fmt.Println(err)
	fmt.Println(rows)

}
