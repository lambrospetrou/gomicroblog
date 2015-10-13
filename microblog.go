package main

import (
	"flag"
	"fmt"
	"github.com/lambrospetrou/gomicroblog/gen"
	"github.com/lambrospetrou/gomicroblog/view"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"time"
)

var validPath = regexp.MustCompile("^/blog/(view|edit|save|del)/([a-zA-Z0-9_-]+)$")

var ViewBuilder *view.Builder = nil

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
	ViewBuilder.Render(w, "index", bundle)
}

// GenerateHandler is called by the website when we want to execute the generator
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I will generate the static site\n\tpath: %s\n", r.URL.Path)
}

func main() {
	fmt.Println("Go Microblog service started!")

	// use all the available cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	var dir_site = flag.String("site", "", "specify a directory that contains the site to generate")
	flag.Parse()

	log.Println("site:", *dir_site)
	if len(*dir_site) > 0 {
		ViewBuilder = view.NewBuilder(filepath.Join(*dir_site, "_layouts"))
		gen.GenerateSite(*dir_site, ViewBuilder)
		return
	} else {
		log.Fatalln("Site source directory not given")
		return
	}
}
