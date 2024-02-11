package gohtmx

import (
	"log"
	"os"
	"text/template"
)

type HandlerContext struct {
	webRoot string
	PageTemplates map[string] *template.Template
}

func NewHandlerContext() *HandlerContext {

	webRoot, exists := os.LookupEnv("WEB_ROOT")
	if exists {
		log.Printf("Using environment variable WEB_ROOT: %v", webRoot)
	} else {
		webRoot, _ = os.Executable()
		log.Printf("WEB_ROOT environment variable not set, using default value: %v", webRoot)
	}

	var hc HandlerContext
	hc.webRoot = webRoot
	hc.PageTemplates = make(map[string] *template.Template)

	return &hc
}

func doOrDie(err error) {
	if err != nil {
		log.Panicf("Oops err: %v ", err)
	}
}
