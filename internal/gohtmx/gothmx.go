package gothmx

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Site struct {
	Heading      string
	Intro        string
	HomeTemplate *template.Template
}

func NewSite(webRoot string) *Site {

	// Create initial instance of site
	site := Site{
		Heading:      "Welcome, from Go HTMX!",
		Intro:        "The home page for my Go with HTMX play site",
		HomeTemplate: &template.Template{},
	}

	// Configure Home Template
	homeTmplFile := webRoot+"/www/home.gotmpl"
	log.Printf("Load home template using: %v", homeTmplFile)
	homeTmpl, err := os.ReadFile(homeTmplFile)
	doOrDie(err)
	
	site.HomeTemplate, _ = template.New("HomePage").Parse(string(homeTmpl))

	return &site
}

func (s *Site) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	err := s.HomeTemplate.Execute(w, s)
	doOrDie(err)

}

func HiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<script src="https://unpkg.com/htmx.org@1.9.4/dist/htmx.min.js"></script>
		</head>

		<body>
			<h1>Hi</h1>
			<button hx-get="/greeting" hx-swap="#data-container">Show Greeting</button>
			<div id="data-container"></div>
		</body>
		</html>
	`)

	_, _ = w.Write(homePage)

}

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	homePage := []byte("Hello from the gohtmx project!")

	_, _ = w.Write(homePage)

}

func doOrDie(err error) {
	if err != nil {
		log.Panicf("Oops err: %v ", err)
	}
}
