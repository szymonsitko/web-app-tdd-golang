package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
	"log"
)

type pageData struct {
	Title string
	Content string
}

func IndexView(w http.ResponseWriter, r *http.Request) {
	rendered_data := pageData{
		Title: "Welcome to The Ticket Booker!",
		Content: "Log-in to create a new booking!",
	}
	dom := "<html><title>{{.Title}}</title><body><h1>Ticket booker</h1><br><h2>{{.Content}}</h2></body>"

	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("main").Parse(dom)
	log.Printf(r.URL.Path)
	if err == nil {
		tmpl.Execute(w, rendered_data)
	}
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexView)
	http.ListenAndServe(":8000", router)
}
