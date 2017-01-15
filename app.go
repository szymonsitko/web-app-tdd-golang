package main

import (
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/template"
)

type pageData struct {
	Title   string
	Content string
}

type userDetails struct {
	Username string
	Password string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

var tpl *template.Template

func IndexView(w http.ResponseWriter, r *http.Request) {
	rendered_data := pageData{
		Title:   "Welcome to The Ticket Booker!",
		Content: "Log-in to create a new booking!",
	} // Think about removing this data structure, move it into template!
	w.Header().Add("Content Type", "text/html")
	err := tpl.ExecuteTemplate(w, "main.gohtml", rendered_data)
	if err != nil {
		log.Printf("Error encountered: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Println(r.URL.Path)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexView)
	http.ListenAndServe(":8000", router)
	log.Print("Http response: ", router)
}
