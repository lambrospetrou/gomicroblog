package main

import (
	"fmt"
	//"github.com/lambrospetrou/gomicroblog/auth"
	//"html/template"
	"net/http"
	"regexp"
	"runtime"
)

type FooterStruct struct {
	Year int
}

type HeaderStruct struct {
	Title string
}

/*
type TemplateBundle struct {
	Post   *BPost
	Footer *FooterStruct
	Header *HeaderStruct
}

type TemplateBundleIndex struct {
	Posts  []*BPost
	Footer *FooterStruct
	Header *HeaderStruct
}
*/
var validPath = regexp.MustCompile("^/blog/(view|edit|save|del)/([a-zA-Z0-9_-]+)$")

/*
var templates = template.Must(template.ParseFiles(
	"templates/partials/header.html",
	"templates/partials/footer.html",
	"templates/view.html",
	"templates/edit.html",
	"templates/add.html",
	"templates/index.html"))

var BLOG_PREFIX = "/blog"


func renderTemplate(w http.ResponseWriter, tmpl string, o interface{}) {
	// now we can call the correct template by the basename filename
	err := templates.ExecuteTemplate(w, tmpl+".html", o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/

// BLOG HANDLERS
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

/*
func viewHandler(w http.ResponseWriter, r *http.Request, id string) {
	bp_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bp, err := LoadBlogPost(bp_id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	bp.BodyHtml = template.HTML(bp.ContentHtml)
	bundle := &TemplateBundle{
		Post:   bp,
		Footer: &FooterStruct{Year: time.Now().Year()},
		Header: &HeaderStruct{Title: bp.Title},
	}

	renderTemplate(w, "view", bundle)
}

func editHandler(w http.ResponseWriter, r *http.Request, id string) {
	bp_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bp, err := LoadBlogPost(bp_id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "edit", bp)
}

func saveHandler(w http.ResponseWriter, r *http.Request, id string) {
	// avoid changing the data with a GET request
	if strings.ToLower(r.Method) == "get" {
		http.Error(w, "Only through the /edit/:id url", http.StatusMethodNotAllowed)
		return
	}

	bp_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bp, err := LoadBlogPost(bp_id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	bp.ContentMarkdown = r.FormValue("markdown")
	bp.Title = r.FormValue("title")
	bp.DateCreated, _ = time.Parse("2006-01-02", r.FormValue("date-created"))
	err = bp.Save()
	if err != nil {
		http.Error(w, "Could not save post!", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, BLOG_PREFIX+"/view/"+bp.IdStr(), http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, id string) {
	bp_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bp, err := LoadBlogPost(bp_id)
	if err != nil {
		http.Error(w, "Could not find the post specified!", http.StatusBadRequest)
		return
	}
	if err = bp.Del(); err != nil {
		http.Error(w, "Could not delete the post specified!", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, BLOG_PREFIX+"/", http.StatusFound)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToLower(r.Method) == "get" {
		renderTemplate(w, "add", "")
		return
	} else if strings.ToLower(r.Method) == "post" {
		bp, err := NewBPost()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bp.ContentMarkdown = r.FormValue("markdown")
		bp.Title = r.FormValue("title")
		bp.Save()

		http.Redirect(w, r, BLOG_PREFIX+"/view/"+bp.IdStr(), http.StatusFound)
		return
	}
	http.Error(w, "Not supported method", http.StatusMethodNotAllowed)
}

// show all posts
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love you %s\n", r.URL.Path)
	posts, err := LoadAllBlogPosts()
	if err != nil {
		http.Error(w, "Could not load blog posts", http.StatusInternalServerError)
		return
	}
	bundle := &TemplateBundleIndex{
		Footer: &FooterStruct{Year: time.Now().Year()},
		Header: &HeaderStruct{Title: "All posts"},
		Posts:  posts,
	}
	renderTemplate(w, "index", bundle)
}
*/

func main() {
	fmt.Println("Go Microblog service started!")

	// use all the available cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Blog endpoints
	/*
		http.HandleFunc("/edit/", auth.BasicHandler(makeHandler(editHandler)))
		http.HandleFunc("/save/", auth.BasicHandler(makeHandler(saveHandler)))
		http.HandleFunc("/del/", auth.BasicHandler(makeHandler(deleteHandler)))
		http.HandleFunc("/add", auth.BasicHandler(addHandler))

		http.HandleFunc("/view/", makeHandler(viewHandler))
		http.HandleFunc("/all", rootHandler)
		http.HandleFunc("/", rootHandler)
	*/
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))

	http.ListenAndServe(":40080", nil)

}
