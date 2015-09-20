package gen

import (
	"fmt"
	"net/http"
)

// show all posts
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I will generate the static site\n\tpath: %s\n", r.URL.Path)
}
