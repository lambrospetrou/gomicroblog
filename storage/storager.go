package storage

import (
	"github.com/lambrospetrou/gomicroblog/post"
)

// Storager defines the interface required by each storage method
// for the blog posts to allow either file-based or other db-based
// solution to be used without changing the code a lot.
type Storager interface {
	// Returns the id as string or an error
	Store(p *post.BPost) (string, error)
	// Returns the loaded blog post or an error
	Load(id string) (*post.BPost, error)
	// Deletes the post with given id and returns an error or nil
	Delete(id string) error

	// Returns a slice with all the blog posts
	LoadAll() []post.BPost
}
