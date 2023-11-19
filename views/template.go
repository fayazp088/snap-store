package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filePath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filePath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)

	if err != nil {
		log.Printf("Error executing template")
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}