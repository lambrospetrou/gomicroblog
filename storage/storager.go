package storage

// Storager defines the interface required by each storage method
// for the blog posts to allow either file-based or other db-based
// solution to be used without changing the code a lot.
type Storager interface {
	Store(p *BPost) (id_t, error)
	Load(id id_t) (*BPost, error)
	Delete(id id_t) error

	LoadAll() []BPost
}
