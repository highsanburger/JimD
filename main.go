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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "167865"
	dbname   = "users"
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if e := db.Ping(); e != nil {
		log.Fatal(e)
	}

	createUser(db)

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

	m.HandleFunc("/login-submit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			u := User{
				username: r.Form.Get("username"),
				email:    "",
				password: r.Form.Get("password"),
			}
			var flag string
			if !userExists(db, u) {
				flag = "Incorrect Username or You don't exist!"
			} else if !correctPwd(db, u) {
				flag = "Incorrect Password"
			} else {
				flag = "Successfully logged in!"
			}
			tmp := template.Must(template.ParseFiles("./templates/dot.html"))
			tmp.Execute(w, flag)

		default:
			fmt.Println("Default")
		}
	})

	m.HandleFunc("/signup-submit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			r.ParseForm()
			insertUser(db, User{r.Form.Get("username"), r.Form.Get("email"), r.Form.Get("password")})
		default:
			fmt.Println("Default")
		}
	})

	fmt.Println("Server is listening on " + PORT)
	log.Fatal(s.ListenAndServe())
}

type User struct {
	username string
	email    string
	password string
}

func createUser(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS uzer(
	id SERIAL PRIMARY KEY, 
	username VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(30) NOT NULL
	)`

	_, err := db.Exec(q)

	if err != nil {
		log.Fatal(err)
	}

}

func userExists(db *sql.DB, user User) bool {
	q := "SELECT EXISTS(SELECT 1 FROM uzer WHERE username = $1)"
	var exists bool
	db.QueryRow(q, user.username).Scan(&exists)
	return exists
}

func correctPwd(db *sql.DB, user User) bool {
	q := "SELECT EXISTS(SELECT 1 FROM uzer WHERE username = $1 AND password = $2)"
	var exists bool
	db.QueryRow(q, user.username, user.password).Scan(&exists)
	return exists
}

func insertUser(db *sql.DB, user User) {
	q := `INSERT INTO uzer(username, email, password)
		VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(q, user.username, user.email, user.password).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted row id :- ", pk)
}
