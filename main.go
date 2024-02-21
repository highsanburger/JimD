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

func serveHTML(mux *http.ServeMux, path, templateFile string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles(templateFile))
		tmp.Execute(w, nil)
		// fmt.Println(r)
	})
}
func main() {

	PORT := ":6969"
	m := http.NewServeMux()
	s := &http.Server{
		Addr:           PORT,
		Handler:        m,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	serveHTML(m, "/", "templates/index.html")
	serveHTML(m, "/login", "templates/login.html")
	serveHTML(m, "/signup", "templates/signup.html")

	m.HandleFunc("/signup-submit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			fmt.Println(r.Form)
			fmt.Println(r.Form.Get("email"))
		default:
			fmt.Println("Default")
		}
	})
	m.HandleFunc("/login-submit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			fmt.Println(r.Form)
			fmt.Println(r.Form.Get("username"))
		default:
			fmt.Println("Default")
		}
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
