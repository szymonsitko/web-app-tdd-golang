package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template
var db *sql.DB
var err error
var LOCALHOST string = "http://localhost:8000"

type pageData struct {
	Content string
}

type userDetails struct {
	Username string
	Password string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	rendered_data := pageData{
		Content: "You are on main page. Please choose from below.",
	}
	w.Header().Add("Content Type", "text/html")
	err := tpl.ExecuteTemplate(w, "main.gohtml", rendered_data)
	if err != nil {
		log.Printf("Error encountered: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println(r.URL.Path)
}

func RegisterView(w http.ResponseWriter, r *http.Request) {
	rendered_data := pageData{
		Content: "Log-in to create a new booking!",
	}
	w.Header().Add("Content Type", "text/html")
	if r.Method == http.MethodPost {
		user := r.FormValue("username")
		passwd, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		email := r.FormValue("email")
		_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?);", user, string(passwd), email)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, LOCALHOST+"/details", 302)
	}
	err := tpl.ExecuteTemplate(w, "register.gohtml", rendered_data)
	if err != nil {
		log.Printf("Error encountered: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println(r.URL.Path)
}

func DetailsView(w http.ResponseWriter, r *http.Request) {
	rendered_data := pageData{
		Content: "User bookings details for User: example",
	}
	w.Header().Add("Content Type", "text/html")
	err := tpl.ExecuteTemplate(w, "details.gohtml", rendered_data)
	if err != nil {
		log.Printf("Error encountered: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println(r.URL.Path)
}

func main() {
	// Init database connection & schema
	db, err = sql.Open("mysql", "simon:irekdudek@tcp/web_go")
	if err != nil {
		log.Fatalf("Error db connection: %s", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255) UNIQUE, password VARCHAR(500), email VARCHAR(355) UNIQUE);")
	if err != nil {
		log.Fatalf("Cannot initialize database. Forcing quit..")
	}
	defer db.Close()
	// Routes and views
	router := mux.NewRouter()
	router.HandleFunc("/register", RegisterView)
	router.HandleFunc("/details", DetailsView)
	router.HandleFunc("/", IndexView)
	http.ListenAndServe(":8000", router)
}
