package view

import (
	"bufio"
	"github.com/lambrospetrou/gomicroblog/post"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type FooterStruct struct {
	Year int
}

type HeaderStruct struct {
	Title string
}

type TemplateBundle struct {
	Post   *post.BPost
	Footer *FooterStruct
	Header *HeaderStruct
}

type TemplateBundleIndex struct {
	Posts  []*post.BPost
	Footer *FooterStruct
	Header *HeaderStruct
}

// Builder is the main object that will compile our views using the template layouts
// and the bundle objects based on each content.
type Builder struct {
	templates *template.Template
}

// NewBuilder returns a builder that will create the views based on the layouts defined
// inside the given directory name.
func NewBuilder(layouts_dir string, templatePaths []string) *Builder {
	builder := &Builder{}
	// compiles and holds all the templates in memory for fast creation
	for i, _ := range templatePaths {
		templatePaths[i] = filepath.Join(layouts_dir, templatePaths[i])
	}
	builder.templates = template.Must(template.ParseFiles(templatePaths...))
	/*
		builder.templates = template.Must(template.ParseFiles(
			filepath.Join(layouts_dir, "partials/header.html"),
			filepath.Join(layouts_dir, "partials/footer.html"),
			filepath.Join(layouts_dir, "post.html"),
			filepath.Join(layouts_dir, "index.html")),
		)
	*/
	return builder
}

const (
	LAYOUT_POST  string = "post"
	LAYOUT_INDEX string = "index"
)

// Render the given view name @vname using the given bundle object @o.
// It writes the output to the given ResponseWriter.
func (b *Builder) Render(w io.Writer, vname string, o interface{}) error {
	// now we can call the correct template by the basename filename
	return b.templates.ExecuteTemplate(w, vname+".html", o)
}

// Render the given view name @vname using the given bundle object @o.
// It writes the output to the given ResponseWriter.
func (b *Builder) RenderToPath(dst string, vname string, o interface{}) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	// now we can call the correct template by the basename filename
	err = b.templates.ExecuteTemplate(w, vname+".html", o)
	w.Flush()
	return err
}

/*
// Render the given view name @vname using the given bundle object @o.
// It writes the output to the given ResponseWriter.
func Render(w http.ResponseWriter, vname string, o interface{}) {
	// now we can call the correct template by the basename filename
	err := templates.ExecuteTemplate(w, vname+".html", o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/
