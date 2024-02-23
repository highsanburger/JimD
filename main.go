package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
	"unicode"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func serveHTML(mux *http.ServeMux, path, templateFile string) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		tmp := template.Must(template.ParseFiles(templateFile))
		tmp.Execute(w, nil)
	})
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "167865"
	dbname   = "postgres"
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
			username := r.Form.Get("username")
			password := r.Form.Get("password")

			var flag string
			if !userExists(db, username) {
				flag = "You don't exist!"
			} else if !correctPwd(db, username, password) {
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

			username := r.Form.Get("username")
			email := r.Form.Get("email")
			password := r.Form.Get("password")
			confirm_password := r.Form.Get("confirm_password")
			cps := checkPasswordStrength(password)
			var flag string
			if userExists(db, username) {
				flag = "User already exists!"
			} else if emailExists(db, email) {
				flag = "Email already in use!"
			} else if password != confirm_password {
				flag = "Passwords don't match"
			} else if cps != "" {
				flag = cps
			} else {
				flag = "Signup successful!"
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				insertUser(db, User{username, email, string(hashedPassword)})
			}
			tmp := template.Must(template.ParseFiles("./templates/dot.html"))
			tmp.Execute(w, flag)
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

func checkPasswordStrength(password string) string {

	if len(password) < 8 {
		return "Password must be at least 8 characters long."
	}

	containsUppercase := false
	for _, char := range password {
		if unicode.IsUpper(char) {
			containsUppercase = true
			break
		}
	}
	if !containsUppercase {
		return "Password must contain at least one uppercase letter."
	}

	containsLowercase := false
	for _, char := range password {
		if unicode.IsLower(char) {
			containsLowercase = true
			break
		}
	}
	if !containsLowercase {
		return "Password must contain at least one lowercase letter."
	}

	containsDigit, _ := regexp.MatchString("[0-9]", password)
	if !containsDigit {
		return "Password must contain at least one digit."
	}

	containsSpecialChar, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, password)
	if !containsSpecialChar {
		return "Password must contain at least one special character."
	}

	return "" // Password is strong
}
func createUser(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY, 
	username VARCHAR(255) UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL
	)`

	_, err := db.Exec(q)

	if err != nil {
		log.Fatal(err)
	}

}

func emailExists(db *sql.DB, email string) bool {
	q := "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
	var exists bool
	db.QueryRow(q, email).Scan(&exists)
	return exists
}

func userExists(db *sql.DB, username string) bool {
	q := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	var exists bool
	db.QueryRow(q, username).Scan(&exists)
	return exists
}

func correctPwd(db *sql.DB, username string, password string) (exists bool) {
	q := "SELECT password FROM users WHERE username = $1"
	var hashed string
	db.QueryRow(q, username).Scan(&hashed)
	b := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if b == nil {
		exists = true
	}
	return exists
}

func insertUser(db *sql.DB, user User) {
	q := `INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3) RETURNING id`
	var pk int
	err := db.QueryRow(q, user.username, user.email, user.password).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted row id :- ", pk)
}
