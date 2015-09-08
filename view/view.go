package view

import (
	"html/template"
	"net/http"
)

// compiles and holds all the templates in memory for fast response
var templates = template.Must(template.ParseFiles(
	"templates/partials/header.html",
	"templates/partials/footer.html",
	"templates/view.html",
	"templates/edit.html",
	"templates/add.html",
	"templates/index.html"))

type FooterStruct struct {
	Year int
}

type HeaderStruct struct {
	Title string
}

type TemplateBundle struct {
	//Post   *BPost
	Footer *FooterStruct
	Header *HeaderStruct
}

type TemplateBundleIndex struct {
	//Posts  []*BPost
	Footer *FooterStruct
	Header *HeaderStruct
}

// Render the given view name @vname using the given bundle object @o.
// It writes the output to the given ResponseWriter.
func Render(w http.ResponseWriter, vname string, o interface{}) {
	// now we can call the correct template by the basename filename
	err := templates.ExecuteTemplate(w, vname+".html", o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
