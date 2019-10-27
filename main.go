package main

import (
	"log"
	"net/http"
	"text/template"
)

const appTitle = "How's My HTTP"
const appDescription = "Find out if your HTTP is good"

type pageData struct {
	Title       string
	Description string
}

func makeStaticHandler() http.Handler {
	//  TODO gzip handler
	fs := http.FileServer(http.Dir("static"))
	return http.StripPrefix("/static/", fs)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.html").ParseGlob("./src/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, pageData{appTitle, appDescription})
}

func main() {
	app := http.NewServeMux()
	app.Handle("/static/", makeStaticHandler())
	app.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":4000", app)
	if err != nil {
		log.Fatal(err)
	}
}
