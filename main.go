package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const appTitle = "How's My HTTP?"
const appDescription = "Find out if your HTTP is good"
const routes = "/ /about"

var (
	assetsToPush = [2]string{"/static/css/shared.css", "/static/css/fonts.css"}
	httpsAddr    = flag.String("httpsAddr", "localhost:4043", "HTTPS address to listen on")
	httpAddr     = flag.String("httpAddr", "", "Plain HTTP address to listen on")
)

type page struct {
	Title       string
	Description string
}

func makeStaticHandler() http.Handler {
	//  TODO gzip handler
	fs := http.FileServer(http.Dir("static"))
	return http.StripPrefix("/static/", fs)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page) {
	t, err := template.ParseGlob("./src/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, tmpl, p)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(routes, r.URL.Path) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404")) // TODO 404 template
		return
	}

	if pusher, ok := w.(http.Pusher); ok {
		for _, path := range assetsToPush {
			if err := pusher.Push(path, nil); err != nil {
				log.Printf("Failed to push %v", path)
			}
		}
	}

	path := strings.Trim(r.URL.Path, "/")
	title := strings.Title(path) + " " + appTitle
	if len(path) == 0 {
		path = "index"
		title = appTitle
	}
	p := page{title, appDescription}
	renderTemplate(w, path, &p)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ci := getClientInfo()
	json, _ := json.Marshal(ci)
	w.Write(json)
	return
}

func main() {
	flag.Parse()
	app := http.NewServeMux()
	app.Handle("/static/", makeStaticHandler())
	app.HandleFunc("/api/v1/", apiHandler)
	app.HandleFunc("/", pageHandler)

	if *httpAddr != "" {
		log.Printf("Listening on http://%v", *httpAddr)
		log.Fatal(http.ListenAndServe(*httpAddr, app))
	}

	log.Printf("Listening on https://%v", *httpsAddr)
	log.Fatal(http.ListenAndServeTLS(*httpsAddr, "development.crt", "development.key", app))
}
