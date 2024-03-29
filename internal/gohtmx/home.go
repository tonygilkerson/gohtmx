package gohtmx

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Home struct {
	Heading string
	Intro   string
}

func (ctx *HandlerContext) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")

	pageValues := Home{
		Heading: "Welcome, from Go HTMX!",
		Intro:   "The home page for my Go with HTMX play site",
	}

	// Reuse template if possible
	tmpl, exists := ctx.PageTemplates["HomePage"]
	if !exists {
		tmplFile := ctx.webRoot + "/templates/home.gotmpl"
		log.Printf("Create template from: %v", tmplFile)

		tmplStr, err := os.ReadFile(tmplFile)
		doOrDie(err)

		tmpl, err = template.New("HomePage").Parse(string(tmplStr))
		doOrDie(err)

		ctx.PageTemplates["HomePage"] = tmpl
	}

	log.Printf("Execute template: %v", tmpl.Name())
	tmpl.Execute(w, pageValues)

}
