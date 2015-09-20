package main

import (
	"fmt"
	"github.com/lambrospetrou/gomicroblog/auth"
	"github.com/lambrospetrou/gomicroblog/gen"
	"github.com/lambrospetrou/gomicroblog/view"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"time"
)

var validPath = regexp.MustCompile("^/blog/(view|edit|save|del)/([a-zA-Z0-9_-]+)$")

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

// show all posts
func rootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love you %s\n", r.URL.Path)
	/*
		posts, err := LoadAllBlogPosts()
		if err != nil {
			http.Error(w, "Could not load blog posts", http.StatusInternalServerError)
			return
		}
	*/
	bundle := &view.TemplateBundleIndex{
		Footer: &view.FooterStruct{Year: time.Now().Year()},
		Header: &view.HeaderStruct{Title: "All posts"},
		//Posts:  posts,
		Posts: nil,
	}
	view.Render(w, "index", bundle)
}

func main() {
	fmt.Println("Go Microblog service started!")

	// use all the available cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Blog endpoints
	http.HandleFunc("/gen-site", auth.BasicHandler(gen.GenerateHandler))
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/s/", http.StripPrefix("/s/", fs))

	log.Fatal(http.ListenAndServe(":40080", nil))

}
